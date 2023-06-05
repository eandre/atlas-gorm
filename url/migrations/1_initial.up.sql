-- Placeholder migration for Atlas to ignore the schema_migrations table
-- when computing schema migrations.
CREATE TABLE IF NOT EXISTS schema_migrations (
    version BIGINT PRIMARY KEY,
    dirty boolean NOT NULL
);