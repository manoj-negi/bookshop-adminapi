CREATE TABLE "permissions" (
  "id" integer PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL ,
  "description" varchar,
  "is_deleted" boolean DEFAULT false,
  "created_at" NOT NULL DEFAULT (now())
);