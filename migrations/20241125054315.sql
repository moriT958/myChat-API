-- Create "threads" table
CREATE TABLE
    "threads" (
        "id" serial NOT NULL,
        "uuid" character varying(64) NOT NULL,
        "topic" text NOT NULL,
        "created_at" timestamp NOT NULL,
        PRIMARY KEY ("id"),
        CONSTRAINT "threads_uuid_key" UNIQUE ("uuid")
    );

-- Create "posts" table
CREATE TABLE
    "posts" (
        "id" serial NOT NULL,
        "uuid" character varying(64) NOT NULL,
        "body" text NOT NULL,
        "thread_id" integer NOT NULL,
        "created_at" timestamp NOT NULL,
        CONSTRAINT "posts_thread_id_fk" FOREIGN KEY ("thread_id") REFERENCES "threads" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
    );