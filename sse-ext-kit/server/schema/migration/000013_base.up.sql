ALTER TABLE company_services
    DROP COLUMN IF EXISTS service_type,
    DROP COLUMN IF EXISTS amount_receive_type,
    ADD COLUMN IF NOT EXISTS category varchar(255) not null,
    ADD COLUMN IF NOT EXISTS name varchar(255) not null;