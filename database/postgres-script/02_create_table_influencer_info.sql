DROP TABLE IF EXISTS "influencer_info";

CREATE TABLE "influencer_info" (
	id INT PRIMARY KEY REFERENCES "influencer_credential"(id), 
	email VARCHAR(500), 
	phone_number VARCHAR(20),
	first_name VARCHAR(100),
	last_name VARCHAR(100),
	introduction TEXT, 
	avt_link TEXT,
	address VARCHAR(255),
	zipcode VARCHAR(10)
);
