DROP TABLE IF EXISTS "business_credential";

CREATE TABLE "business_credential" (
    id SERIAL PRIMARY KEY,              -- Auto-incrementing primary key
    email VARCHAR(255) UNIQUE,          -- Unique email
    password BYTEA,                     -- Password stored as bytea for hashed values
    phone_number VARCHAR(20),           -- Phone number
    is_verified BOOLEAN DEFAULT FALSE   -- Boolean to check if verified, defaulting to false
);