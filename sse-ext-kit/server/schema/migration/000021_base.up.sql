DROP TABLE IF EXISTS company_qr;

CREATE TABLE IF NOT EXISTS company_qr (
  id uuid not null primary key,
    qr_id varchar(255) not null,
    qr_code varchar(30000) not null,
    qr_type varchar(255) not null,
    service_type varchar(255) not null,
    service_id uuid not null,
    created_at timestamp without time zone,
    CONSTRAINT fk_qr_service_mapping FOREIGN KEY (service_type, service_id)
    REFERENCES service_mapping (service_type, service_id)
);

ALTER TABLE selling_point
    DROP COLUMN IF EXISTS qr_id;

ALTER TABLE company_services
    DROP COLUMN IF EXISTS qr_id;
