CREATE TABLE IF NOT EXISTS album (
	id         SERIAL PRIMARY KEY,
	title      VARCHAR NOT NULL,
	artist     VARCHAR NOT NULL,
	price      DECIMAL(5,2) NOT NULL
);
