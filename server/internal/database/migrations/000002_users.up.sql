CREATE TABLE users (
	id SERIAL UNIQUE,
	login VARCHAR UNIQUE NOT NULL,
	password_hash VARCHAR NOT NULL
);