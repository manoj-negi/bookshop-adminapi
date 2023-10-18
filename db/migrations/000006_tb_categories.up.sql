
CREATE TABLE "categories" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL,
  "is_special" varchar,
  "is_deleted" boolean DEFAULT 'false',
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);