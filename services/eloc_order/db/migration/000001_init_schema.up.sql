CREATE TABLE "carts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "quantity" bigint NOT NULL DEFAULT 1,
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "total_amount" bigint NOT NULL,
  "status" varchar NOT NULL DEFAULT 'PENDING',
  "shipping_address" varchar NOT NULL,
  "customer_phone" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_items" (
  "id" bigserial PRIMARY KEY,
  "order_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "quantity" bigint NOT NULL,
  "price" bigint NOT NULL
);

CREATE INDEX ON "carts" ("user_id");

CREATE INDEX ON "orders" ("user_id");

CREATE INDEX ON "orders" ("status");

CREATE INDEX ON "order_items" ("order_id");

COMMENT ON COLUMN "orders"."status" IS 'PENDING | SHIPPING | DELIVERED';

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id") DEFERRABLE INITIALLY IMMEDIATE;
