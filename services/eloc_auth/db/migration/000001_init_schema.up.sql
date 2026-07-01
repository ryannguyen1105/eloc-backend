CREATE TABLE "roles" (
  "id" varchar PRIMARY KEY,
  "description" varchar NOT NULL
);

INSERT INTO "roles" ("id", "description") VALUES 
  ('ADMIN', 'System Administrator with full access rights'),
  ('STAFF', 'Store Staff with limited operational access'),
  ('CUSTOMER', 'Default Customer account for general users')
ON CONFLICT (id) DO NOTHING;

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "role_id" varchar  NOT NULL,
  "is_active" boolean NOT NULL DEFAULT true,
  "is_verified" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_tokens" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "refresh_token" varchar UNIQUE NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("role_id");

CREATE INDEX ON "user_tokens" ("user_id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "user_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;;
