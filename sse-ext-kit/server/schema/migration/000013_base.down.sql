ALTER TABLE company_services
    ADD COLUMN IF NOT EXISTS service_type varchar(255),
    ADD COLUMN IF NOT EXISTS amount_receive_type varchar(255),
    DROP COLUMN IF EXISTS category,
    DROP COLUMN IF EXISTS name;