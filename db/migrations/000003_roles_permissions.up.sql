CREATE TABLE "role_permission" (
  "id" integer PRIMARY KEY,
  "role_id" integer,
  "permission_id" integer,
  "is_deleted" boolean,
  "created_at" timestamp
);


ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id");