CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "created_at" date NOT NULL DEFAULT now()
);

CREATE TABLE "scrape" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "scrapping_url" varchar NOT NULL,
  "scrapped_data" varchar NOT NULL,
  "created_at" date NOT NULL DEFAULT now()
);

ALTER TABLE "scrape" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "scrape" ("scrapping_url", "created_at");

DROP TABLE users;

DROP TABLE scrape;
