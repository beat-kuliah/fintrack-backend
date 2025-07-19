CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "email" varchar(256) UNIQUE NOT NULL,
    "name" varchar(256) NOT NULL,
    "hashed_password" varchar(256) NOT NULL,
    "verified" timestamptz ,
    "token_forgot" text,
    "onboard" bool DEFAULT (true),
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);