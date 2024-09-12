DROP TABLE IF EXISTS "post";

CREATE TABLE "post" (
    id SERIAL PRIMARY KEY,                          -- Auto-incrementing unique primary key
    business_id INT REFERENCES "business_info"(id), -- Foreign key reference to business_info(id)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of creation, defaults to current time
    content TEXT,                                   -- Content of the post
    pay_amount FLOAT,                               -- Payment amount (float for money)
    is_active BOOLEAN DEFAULT TRUE,                -- Boolean to indicate if the post is active, defaulting to TRUE
    work_time TIMESTAMP                            -- Timestamp for work time
);