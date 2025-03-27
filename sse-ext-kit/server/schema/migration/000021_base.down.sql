CREATE TABLE IF NOT EXISTS company_qr
(
    id varchar(255) not null primary key,
    code varchar(30000) not null,
    type varchar(255) not null default 'STATIC',
    created_at timestamp without time zone
);

ALTER TABLE selling_point
    ADD COLUMN IF NOT EXISTS qr_id varchar(255);


ALTER TABLE company_services
    ADD COLUMN IF NOT EXISTS qr_id varchar(255);

