CREATE TABLE "roles" (
  "id" integer PRIMARY KEY,
  "role" varchar UNIQUE NOT NULL,
  "desciption" varchar,
  "is_active" boolean DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);