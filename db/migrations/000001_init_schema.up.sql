CREATE TABLE "cat" (
    "id" SERIAL PRIMARY KEY,
    "breed" VARCHAR (55) NOT NULL,
    "color" VARCHAR (55) NOT NULL,
    "gender" VARCHAR (55) NOT NULL,
    "owner" BIGINT REFERENCES owner (id)
);
CREATE TABLE "owner" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR (55) NOT NULL,
    "usermane" VARCHAR (55) NOT NULL,
    "cats" BIGINT [] REFERENCES cat (id)
)