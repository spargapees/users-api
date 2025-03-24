CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "name" varchar NOT NULL,
                         "surname" varchar NOT NULL,
                         "password_hash" varchar NOT NULL,
                         "email" varchar NOT NULL
);