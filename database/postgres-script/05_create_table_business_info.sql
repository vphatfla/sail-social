DROP TABLE IF EXISTS "business_info";

CREATE TABLE "business_info" (
    id INT PRIMARY KEY REFERENCES "business_credential"(id), 
    email VARCHAR(255),                             
    phone_number VARCHAR(20),                       
    first_name VARCHAR(100),                        
    last_name VARCHAR(100),                         
    business_name VARCHAR(255),                     
    business_type VARCHAR(100),                     
    introduction TEXT,                              
    avt_link TEXT,                                   
    address VARCHAR(255),                            
    zipcode VARCHAR(10)                              
);
