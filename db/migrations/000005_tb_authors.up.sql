CREATE TABLE "authors" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL,
  "is_deleted" boolean DEFAULT 'false',
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);
