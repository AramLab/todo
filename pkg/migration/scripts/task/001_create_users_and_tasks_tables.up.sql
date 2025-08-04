CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS users;

CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       user_id UUID NOT NULL,
                       title TEXT NOT NULL,
                       description TEXT,
                       status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
                       created_at TIMESTAMP DEFAULT now()
);
