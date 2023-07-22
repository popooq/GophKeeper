CREATE TABLE keeper (
    id serial unique,
    username varchar(12),
    service varchar(32),
    entry varchar(255),
    metadata varchar(255)
);