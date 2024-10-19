DROP TABLE IF EXISTS "post";

CREATE TABLE "post" (
    id SERIAL PRIMARY KEY,                          
    business_id INT REFERENCES "business_info"(id), 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    content TEXT,                                   
    pay_amount FLOAT,                               
    is_active BOOLEAN DEFAULT TRUE,                
    work_time VARCHAR(255)                           
);
