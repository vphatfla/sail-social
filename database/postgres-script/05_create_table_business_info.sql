DROP TABLE IF EXISTS "business_info";

CREATE TABLE "business_info" (
    id INT PRIMARY KEY REFERENCES "business_credential"(id),  -- Primary key with foreign key reference to business_credential(id)
    email VARCHAR(255),                                     -- Email
    phone_number VARCHAR(20),                               -- Phone number
    first_name VARCHAR(100),                                -- First name
    last_name VARCHAR(100),                                 -- Last name
    business_name VARCHAR(255),                             -- Business name
    business_type VARCHAR(100),                             -- Business type
    introduction TEXT,                                      -- Introduction (text)
    avt_link TEXT,                                          -- Avatar link
    street_address VARCHAR(255),                            -- Street address
    zipcode VARCHAR(10)                                    -- Zip code
);
