DROP TABLE IF EXISTS "creator_post_applied";

CREATE TABLE "creator_post_applied" (
    creator_id INT REFERENCES "creator_info"(id),  
    post_id INT REFERENCES "post"(id),
    message VARCHAR(255),
    PRIMARY KEY (creator_id, post_id)
);
