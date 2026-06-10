CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "slug" varchar UNIQUE NOT NULL
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "category_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "slug" varchar UNIQUE NOT NULL,
  "sku" varchar UNIQUE NOT NULL,
  "price" bigint NOT NULL,
  "stock" int NOT NULL DEFAULT 0,
  "attributes" jsonb,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_images" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "image_url" varchar NOT NULL,
  "is_primary" boolean NOT NULL DEFAULT false
);

CREATE INDEX ON "products" ("category_id");

CREATE INDEX ON "products" ("slug");

CREATE INDEX ON "product_images" ("product_id");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "product_images" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id") DEFERRABLE INITIALLY IMMEDIATE;
