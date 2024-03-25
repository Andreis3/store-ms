-- SELECT 'CREATE DATABASE store_db WITH ENCODING = ''UTF8'';' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'store_db')\gexec
------------------------------------------------------------------
--- Level 1 -> Product
------------------------------------------------------------------
CREATE TYPE status AS ENUM ('active', 'inactive');

CREATE TABLE IF NOT EXISTS "groups" (
    "id" UUID PRIMARY KEY NOT NULL,
    "group_name" VARCHAR(255) NOT NULL,
    "code" VARCHAR(100) NOT NULL,
    "status" status NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL
);


------------------------------------------------------------------
--- Level 1 -> Group
------------------------------------------------------------------
ALTER TABLE "groups" ADD UNIQUE (group_name, code);

------------------------------------------------------------------
-- INDEX
------------------------------------------------------------------

CREATE INDEX IF NOT EXISTS "idx_groups_group_name" ON "groups" (group_name, code);