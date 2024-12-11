-- Create "users" table
CREATE TABLE "users" ("id" serial NOT NULL, "uuid" uuid NOT NULL, "username" character varying(30) NOT NULL, "password" character varying(255) NOT NULL, "created_at" timestamp NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "users_username_key" UNIQUE ("username"), CONSTRAINT "users_uuid_key" UNIQUE ("uuid"));
-- Modify "posts" table
ALTER TABLE "posts" ADD COLUMN "user_id" integer NOT NULL, ADD CONSTRAINT "post_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "threads" table
ALTER TABLE "threads" ADD COLUMN "user_id" integer NOT NULL, ADD CONSTRAINT "thread_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
