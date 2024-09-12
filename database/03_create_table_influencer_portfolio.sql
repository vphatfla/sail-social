DROP TABLE IF EXISTS "influencer_portfolio";

CREATE TABLE "influencer_portfolio" (
    id INT PRIMARY KEY REFERENCES "influencer_info"(id),  -- Primary key with foreign key reference to influencer_info(id)
    instagram_link VARCHAR(255),                          -- Instagram link
    tiktok_link VARCHAR(255),                             -- TikTok link
    instagram_follower INT,                               -- Instagram follower count
    tiktok_follower INT                                   -- TikTok follower count
);
