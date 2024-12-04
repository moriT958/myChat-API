-- Modify "posts" table
ALTER TABLE "posts" ALTER COLUMN "uuid" TYPE uuid USING uuid::uuid;
-- Modify "threads" table
ALTER TABLE "threads" ALTER COLUMN "uuid" TYPE uuid USING uuid::uuid;
