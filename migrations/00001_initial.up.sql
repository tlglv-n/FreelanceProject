DO $$
DECLARE author_id UUID;

  BEGIN
    -- EXTENSIONS --
    CREATE EXTENSION IF NOT EXISTS pgcrypto;

    -- TABLES --
    CREATE TABLE IF NOT EXISTS customers (
        created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        id          UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
        full_name   VARCHAR NOT NULL,
        pseudonym   VARCHAR NOT NULL
    );

    CREATE TABLE IF NOT EXISTS hires (
        created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        id          UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
        customer_id   UUID NOT NULL REFERENCES customers (id),
        job_name        VARCHAR NOT NULL,
        amount       INT NOT NULL,
        description VARCHAR NOT NULL,
        position VARCHAR NOT NULL
    );

    CREATE TABLE IF NOT EXISTS workers (
        created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        id          UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
        full_name   VARCHAR NOT NULL,
        description VARCHAR NOT NULL,
        position VARCHAR NOT NULL
    );



  COMMIT;
END $$;