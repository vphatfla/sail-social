DROP TABLE IF EXISTS "creator_post_applied";

CREATE TABLE "creator_post_applied" (
    creator_id INT REFERENCES "creator_info"(id),  -- Foreign key reference to creator_info(id)
    post_id INT REFERENCES "post"(id),                    -- Foreign key reference to post(id)
    PRIMARY KEY (creator_id, post_id)                  -- Composite unique primary key
);