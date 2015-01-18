CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Table: "Users"

-- DROP TABLE "Users";

CREATE TABLE "Users" (
    "id" uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    "email" character varying(255) NOT NULL UNIQUE,
    "password" character(60) NOT NULL,
    "created_at" timestamp with time zone NOT NULL DEFAULT now()
);