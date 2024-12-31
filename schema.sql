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

CREATE TABLE "rooms" (
    "id" SERIAL NOT NULL,
    "uuid" UUID NOT NULL,
    "topic" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "user_id" INTEGER NOT NULL,
    PRIMARY KEY ("id"),
    UNIQUE ("uuid"),
    CONSTRAINT "rooms_user_id_fk"
        FOREIGN KEY ("user_id")
        REFERENCES "users" ("id")
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE INDEX idx_users_uuid ON users(uuid);
CREATE INDEX idx_rooms_uuid ON rooms(uuid);
