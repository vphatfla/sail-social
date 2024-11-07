DROP TABLE IF EXISTS "creator_info";

CREATE TABLE "creator_info" (
	id INT PRIMARY KEY REFERENCES "creator_credential"(id), 
	email VARCHAR(500), 
	phone_number VARCHAR(20),
	first_name VARCHAR(100),
	last_name VARCHAR(100),
	introduction TEXT, 
	avt_link TEXT,
	address VARCHAR(255),
	zipcode VARCHAR(10),
    city VARCHAR(255),
    state VARCHAR(255),
    country varchar(255)
);
