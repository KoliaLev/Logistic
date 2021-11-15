BEGIN;
-- ALTER TABLE users
-- ADD COLUMN transport TEXT;
CREATE TABLE transports (
    id SERIAL,
    user_id INT,
    transport_type TEXT PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_users FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE transport_types (
    id SERIAL PRIMARY KEY,
    name_type TEXT,
    speed INT,
    volume INT,
    acceleration INT,
    CONSTRAINT fk_transports FOREIGN KEY(name_type) REFERENCES transports(transport_type) ON DELETE CASCADE
);
COMMIT;