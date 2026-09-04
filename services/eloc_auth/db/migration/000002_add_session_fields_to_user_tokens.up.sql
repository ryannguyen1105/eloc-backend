ALTER TABLE "user_tokens"
ADD "user_agent" varchar NOT NULL,
ADD "client_ip" varchar NOT NULL,
ADD "is_blocked" boolean NOT NULL DEFAULT false;