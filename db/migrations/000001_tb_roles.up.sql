CREATE TABLE "roles" (
  "id" integer PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "description" varchar,
  "is_deleted" boolean DEFAULT 'false',
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);