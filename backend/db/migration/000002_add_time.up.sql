ALTER TABLE users 
ADD COLUMN created_at timestamp DEFAULT current_timestamp,
ADD COLUMN updated_at timestamp DEFAULT current_timestamp;

ALTER TABLE todos
ADD COLUMN created_at timestamp DEFAULT current_timestamp,
ADD COLUMN updated_at timestamp DEFAULT current_timestamp;