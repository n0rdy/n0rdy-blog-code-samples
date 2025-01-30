CREATE SCHEMA schema_202501;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE schema_202501.users
(
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ssn_hash  TEXT NOT NULL UNIQUE,
    user_info TEXT NOT NULL
);

CREATE INDEX users_ssn_hashed_idx ON schema_202501.users (ssn_hash);