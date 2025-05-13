
CREATE TYPE "user_role" AS ENUM (
  'regular_user',
  'admin_user'
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "role" user_role NOT NULL,
  "password" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp
);

CREATE INDEX ON "users" ("email");

CREATE UNIQUE INDEX users_email_unique ON "users" USING btree ("email", "deleted_at");
