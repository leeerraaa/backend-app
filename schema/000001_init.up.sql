BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
    "id" uuid DEFAULT uuid_generate_v4(),
    "login" varchar NOT NULL,
    "password" varchar NOT NULL,
    "date_of_creation" datetime NOT NULL DEFAULT(GETDATE()),
    PRIMARY KEY ("id")
);

CREATE TABLE "documents" (
    "id" uuid DEFAULT uuid_generate_v4(),
    "user_id" uuid NOT NULL,
    "name_of_subject" varchar NOT NULL,
    "date_of_creation" datetime NOT NULL DEFAULT(GETDATE()),
    PRIMARY KEY ("id")
);

ALTER TABLE
    "documents"
ADD
    CONSTRAINT "users_documents" FOREIGN KEY ("user_id") REFERENCES "users" ("id");

COMMIT;

INSERT INTO
    users(login, password)
VALUES
    ('leeerraaa', '1234LK'),