-- Create "users" table
CREATE TABLE "users" ("id" serial NOT NULL, "uuid" uuid NOT NULL, "username" character varying(30) NOT NULL, "password" character varying(255) NOT NULL, "created_at" timestamp NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "users_username_key" UNIQUE ("username"), CONSTRAINT "users_uuid_key" UNIQUE ("uuid"));
-- Create "threads" table
CREATE TABLE "threads" ("id" serial NOT NULL, "uuid" uuid NOT NULL, "topic" text NOT NULL, "created_at" timestamp NOT NULL, "user_id" integer NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "threads_uuid_key" UNIQUE ("uuid"), CONSTRAINT "thread_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "posts" table
CREATE TABLE "posts" ("id" serial NOT NULL, "uuid" uuid NOT NULL, "body" text NOT NULL, "thread_id" integer NOT NULL, "created_at" timestamp NOT NULL, "user_id" integer NOT NULL, CONSTRAINT "post_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "posts_thread_id_fk" FOREIGN KEY ("thread_id") REFERENCES "threads" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
