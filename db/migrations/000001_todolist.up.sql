CREATE TABLE tasks (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    title VARCHAR,
    description VARCHAR,
    due_date TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)