DROP TABLE IF EXISTS "business_credential";

CREATE TABLE "business_credential" (
    id SERIAL PRIMARY KEY,              
    email VARCHAR(255) UNIQUE,          
    password BYTEA,                     
    phone_number VARCHAR(20),           
    is_verified BOOLEAN DEFAULT FALSE,
    is_onboarded BOOLEAN DEFAULT FALSE
);
