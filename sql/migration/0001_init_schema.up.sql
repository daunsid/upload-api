


CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "user_name" varchar NOT NULL,
    "password_hash" varchar(255) UNIQUE NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "files" (
    "id" bigserial PRIMARY KEY,
    "user_id" UUID NOT NULL,
    "file_name" varchar NOT NULL,
    "file_id" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("user_name");
CREATE INDEX ON "files" ("user_id");

ALTER TABLE "files" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");



