CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE table users(
    user_id serial PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255)  NOT NULL
);

CREATE table todos(
    todo_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id serial,
    title VARCHAR(255) NOT NULL,
    detail TEXT,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);