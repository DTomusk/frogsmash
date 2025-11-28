package repos

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/comparison/models"
	"frogsmash/internal/app/shared"
)

type SubmissionRepo struct{}

func NewSubmissionRepo() *SubmissionRepo {
	return &SubmissionRepo{}
}

func (r *SubmissionRepo) GetLatestSubmissionByUser(userID string, ctx context.Context, db shared.DBTX) (*models.ImageUpload, error) {
	query := "SELECT id, user_id, image_url, uploaded_at FROM image_uploads WHERE user_id = $1 ORDER BY uploaded_at DESC LIMIT 1"
	row := db.QueryRowContext(ctx, query, userID)
	var upload models.ImageUpload

	if err := row.Scan(&upload.ID, &upload.UserID, &upload.URL, &upload.UploadedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &upload, nil
}

func (r *SubmissionRepo) GetTotalDataUploaded(ctx context.Context, db shared.DBTX) (int64, error) {
	var totalSize int64
	query := "SELECT COALESCE(SUM(file_size), 0) FROM image_uploads"
	row := db.QueryRowContext(ctx, query)
	if err := row.Scan(&totalSize); err != nil {
		return 0, err
	}

	return totalSize, nil
}

func (r *SubmissionRepo) GetTimeOfLatestSubmission(userID string, ctx context.Context, db shared.DBTX) (string, error) {
	query := "SELECT uploaded_at FROM image_uploads WHERE user_id = $1 ORDER BY uploaded_at DESC LIMIT 1"
	row := db.QueryRowContext(ctx, query, userID)
	var uploadedAt string
	if err := row.Scan(&uploadedAt); err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return uploadedAt, nil
}

func (r *SubmissionRepo) InsertImageUploadRecord(userID string, fileSize int64, imageURL string, ctx context.Context, db shared.DBTX) error {
	query := "INSERT INTO image_uploads (user_id, file_size, image_url, uploaded_at) VALUES ($1, $2, $3, NOW())"
	_, err := db.ExecContext(ctx, query, userID, fileSize, imageURL)
	return err
}
