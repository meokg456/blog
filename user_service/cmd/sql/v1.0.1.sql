CREATE DATABASE user_service

CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(50) NOT NULL UNIQUE,
	password VARCHAR(200) NOT NULL,
	full_name TEXT NOT NULL
)
