CREATE TABLE "users" (
  "id" varchar(64) UNIQUE PRIMARY KEY NOT NULL,
  "username" varchar(36) UNIQUE NOT NULL,
  "password" bytea NOT NULL,
  "password_updated_at" varchar(64) NOT NULL DEFAULT (timezone('Asia/Singapore', now())),
  "last_login" varchar(64) NOT NULL DEFAULT (timezone('Asia/Singapore', now())),
  "status" varchar(10) NOT NULL DEFAULT 'offline',
  "created_at" timestamptz NOT NULL DEFAULT (timezone('Asia/Singapore', now()))
);

CREATE TABLE "profiles" (
  "id" varchar(64) UNIQUE PRIMARY KEY NOT NULL,
  "user_id" varchar(64) NOT NULL,
  "first_name" varchar(36) NOT NULL,
  "middle_name" varchar(36) NOT NULL,
  "last_name" varchar(36) NOT NULL,
  "full_name" varchar(100) NOT NULL,
  "email" varchar(36) UNIQUE NOT NULL,
  "contact_number" varchar(20) UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (timezone('Asia/Singapore', now())),
  "updated_at" varchar(64) NOT NULL DEFAULT (timezone('Asia/Singapore', now()))
);

CREATE TABLE "clients" (
  "id" varchar(64) UNIQUE PRIMARY KEY NOT NULL,
  "first_name" varchar(36) NOT NULL,
  "middle_name" varchar(36) NOT NULL,
  "last_name" varchar(36) NOT NULL,
  "full_name" varchar(100) NOT NULL,
  "email" varchar(36) UNIQUE NOT NULL,
  "message" varchar(360) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (timezone('Asia/Singapore', now())),
  "updated_at" varchar(64) NOT NULL DEFAULT (timezone('Asia/Singapore', now()))
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("status");

CREATE INDEX ON "profiles" ("full_name");

CREATE INDEX ON "profiles" ("user_id");

CREATE INDEX ON "profiles" ("email");

CREATE INDEX ON "profiles" ("contact_number");

CREATE INDEX ON "clients" ("first_name");

CREATE INDEX ON "clients" ("last_name");

CREATE INDEX ON "clients" ("full_name");

CREATE INDEX ON "clients" ("email");

ALTER TABLE "profiles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");