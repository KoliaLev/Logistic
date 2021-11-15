BEGIN;
CREATE TYPE sex as ENUM ('male', 'female', 'other');
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT,
    password TEXT,
    birthday TEXT,
    sex sex,
    created_at TIMESTAMP WITH TIME ZONE
);
COMMIT;