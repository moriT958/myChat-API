CREATE TABLE "users" (
    "id" SERIAL,
    "uuid" UUID NOT NULL,
    "username" VARCHAR(30) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    PRIMARY KEY ("id"),
    UNIQUE ("uuid"),
    UNIQUE ("username")
);

CREATE TABLE "threads" (
    "id" SERIAL NOT NULL,
    "uuid" UUID NOT NULL,
    "topic" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "user_id" INTEGER NOT NULL,
    PRIMARY KEY ("id"),
    UNIQUE ("uuid"),
    CONSTRAINT "thread_user_id_fk"
        FOREIGN KEY ("user_id")
        REFERENCES "users" ("id")
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE "posts" (
    "id" SERIAL,
    "uuid" UUID NOT NULL,
    "body" TEXT NOT NULL,
    "thread_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "user_id" INTEGER NOT NULL,
    CONSTRAINT "posts_thread_id_fk" 
        FOREIGN KEY ("thread_id") 
        REFERENCES "threads" ("id")
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT "post_user_id_fk"
        FOREIGN KEY ("user_id")
        REFERENCES "users" ("id")
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);  

