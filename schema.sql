CREATE TABLE "threads" (
    "id" SERIAL NOT NULL,
    "uuid" VARCHAR(64) NOT NULL,
    "topic" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    PRIMARY KEY ("id"),
    UNIQUE ("uuid")
);

CREATE TABLE "posts" (
    "id" SERIAL,
    "uuid" VARCHAR(64) NOT NULL,
    "body" TEXT NOT NULL,
    "thread_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    CONSTRAINT "posts_thread_id_fk" 
        FOREIGN KEY ("thread_id") 
        REFERENCES "threads" ("id")
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);  