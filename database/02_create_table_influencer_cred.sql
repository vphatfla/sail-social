DROP TABLE IF EXISTS "influencer_credential";

CREATE TABLE "influencer_credential" (
	id SERIAL PRIMRARY KEY, 
	email VARCHAR(255) UNIQUE, 
	password BYTEA,
	phone_number VARCHAR(20), 
	is_verified BOOLEAN DEFAULT FALSE
);
