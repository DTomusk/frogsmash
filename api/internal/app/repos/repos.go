package repos

import "database/sql"

type EventsRepo struct {
	DB *sql.DB
}

type ItemsRepo struct {
	DB *sql.DB
}

func NewEventsRepo(db *sql.DB) *EventsRepo {
	return &EventsRepo{DB: db}
}

func (r *EventsRepo) LogEvent(winnerId, loserId string) error {
	_, err := r.DB.Exec(
		"INSERT INTO events (winner_id, loser_id) VALUES ($1, $2)",
		winnerId, loserId,
	)
	return err
}

func NewItemsRepo(db *sql.DB) *ItemsRepo {
	return &ItemsRepo{DB: db}
}
