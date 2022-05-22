CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "login" varchar NOT NULL,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "date_of_creation" timestamp NOT NULL
);

CREATE TABLE "documents" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "specialty" varchar,
  "educational_level" varchar,
  "educational_program" varchar,
  "subject" varchar,
  "lectures" int,
  "practical_classes" int,
  "laboratory_classes" int,
  "date_of_creation" timestamp NOT NULL
);

ALTER TABLE
  "documents"
ADD
  CONSTRAINT "users_documents" FOREIGN KEY ("user_id") REFERENCES "users" ("id");

INSERT INTO
  users(id, login, username, password, date_of_creation)
VALUES
  (
    'bb3e9f4c-3cdc-4f3a-8626-22f318058146',
    'leeerraaa',
    'Кундик Валерія Олексіївна',
    '7375706572686173686b65797276ac07766e82cbe643360bea22303fb9aad401',
    CURRENT_TIMESTAMP
  );