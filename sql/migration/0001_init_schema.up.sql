CREATE EXTENSION IF NOT EXISTS "uuid-ossp";



CREATE TABLE "users" (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "user_name" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "files" (
    "id" bigserial PRIMARY KEY,
    "user_id" varchar NOT NULL,
    "file_name" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("user_name");
CREATE INDEX ON "files" ("user_id");

ALTER TABLE "files" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");



