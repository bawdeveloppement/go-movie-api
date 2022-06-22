CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "country_code" int
);

CREATE TABLE "countries" (
  "code" int PRIMARY KEY,
  "country_name" varchar NOT NULL,
  "continent_name" varchar NOT NULL
);

CREATE TABLE "movies" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar NOT NULL,
  "director_id" int,
  "admin_id" int,
  "release_year" date NOT NULL,
  "production_country_code" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "directors" (
  "id" SERIAL PRIMARY KEY,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL
);

ALTER TABLE "movies" ADD FOREIGN KEY ("director_id") REFERENCES "directors" ("id");

ALTER TABLE "movies" ADD FOREIGN KEY ("admin_id") REFERENCES "users" ("id");

ALTER TABLE "movies" ADD FOREIGN KEY ("production_country_code") REFERENCES "countries" ("code");

ALTER TABLE "users" ADD FOREIGN KEY ("country_code") REFERENCES "countries" ("code");
