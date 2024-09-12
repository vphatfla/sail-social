DROP TABLE IF EXISTS "influencer_post_saved";

CREATE TABLE "influencer_post_saved" (
    influencer_id INT REFERENCES "influencer_info"(id),  -- Foreign key reference to influencer_info(id)
    post_id INT REFERENCES "post"(id),                    -- Foreign key reference to post(id)
    PRIMARY KEY (influencer_id, post_id)                  -- Composite unique primary key
);
