DROP TABLE IF EXISTS "creator_portfolio";

CREATE TABLE "creator_portfolio" (
    id INT PRIMARY KEY REFERENCES "creator_info"(id),  
    instagram_link VARCHAR(255),                       
    tiktok_link VARCHAR(255),                          
    instagram_follower INT,                            
    tiktok_follower INT                                
);
