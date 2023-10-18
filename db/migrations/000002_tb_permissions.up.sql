CREATE TABLE "permissions" (
  "id" integer PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "permission" varchar,
  "is_deleted" boolean DEFAULT 'false',
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);