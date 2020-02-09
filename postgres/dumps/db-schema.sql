-- CREATE DATABASE dbpostgres;
GRANT ALL PRIVILEGES ON DATABASE dbpostgres TO postgres;

\c dbpostgres;

CREATE TABLE "url" (
	"id" SERIAL NOT NULL UNIQUE,
	"long" VARCHAR(240) NOT NULL,
	"short" VARCHAR(50) UNIQUE,
	CONSTRAINT "url_pkey" PRIMARY KEY ("id")
);
CREATE INDEX "url_short" ON "url" ("short");

CREATE TABLE "urlstats" (
	"url_short" VARCHAR(50) NOT NULL,
	"today" DATE NOT NULL,
	"count" INTEGER NOT NULL DEFAULT 0,
	CONSTRAINT "urlstats_pkey" PRIMARY KEY ("url_short", "today")
);
ALTER TABLE "urlstats" ADD FOREIGN KEY ("url_short") REFERENCES "url" ("short");
