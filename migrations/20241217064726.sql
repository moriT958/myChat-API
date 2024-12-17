-- Create index "idx_posts_uuid" to table: "posts"
CREATE INDEX "idx_posts_uuid" ON "posts" ("uuid");
-- Create index "idx_threads_uuid" to table: "threads"
CREATE INDEX "idx_threads_uuid" ON "threads" ("uuid");
-- Create index "idx_users_uuid" to table: "users"
CREATE INDEX "idx_users_uuid" ON "users" ("uuid");
