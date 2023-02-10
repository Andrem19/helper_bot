CREATE TABLE "commands" (
  "id" BIGSERIAL PRIMARY KEY,
  "cmd" varchar NOT NULL UNIQUE,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);