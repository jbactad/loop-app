CREATE TABLE "public"."surveys" (
  "id" uuid NOT NULL,
  "name" varchar(255) NOT NULL,
  "description" varchar(255) NOT NULL,
  "question" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  PRIMARY KEY ("id")
);
CREATE TABLE "public"."survey_responses" (
  "id" uuid NOT NULL,
  "survey_id" uuid NOT NULL,
  "answer" varchar(511) NOT NULL,
  "rating" int4 NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  PRIMARY KEY ("id"),
  FOREIGN KEY ("survey_id") REFERENCES "public"."surveys" ("id") ON DELETE CASCADE
);
