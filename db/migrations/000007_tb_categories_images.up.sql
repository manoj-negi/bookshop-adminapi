CREATE TABLE "categories_images" (
  "id" SERIAL PRIMARY KEY,
  "category_id" integer NOT NULL,
  "image" varchar,
  "is_deleted" boolean DEFAULT 'false',
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  CONSTRAINT fk_category_images_category FOREIGN KEY (category_id) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE RESTRICT
);