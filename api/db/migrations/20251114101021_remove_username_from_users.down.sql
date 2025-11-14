-- Revert the removal of the username column from the users table, use a UUID as default
ALTER TABLE users ADD COLUMN IF NOT EXISTS username VARCHAR(50) UNIQUE NOT NULL DEFAULT uuid_generate_v4();