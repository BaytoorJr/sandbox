DROP TABLE IF EXISTS managers;
DROP TABLE IF EXISTS selling_point;
DROP TABLE IF EXISTS company_services;
DROP TABLE IF EXISTS company_qr;
DROP TABLE IF EXISTS mundea_contracts;

CREATE TABLE IF NOT EXISTS selling_point
(
    id UUID not null primary key,
    company_id UUID not null,
    qr_id varchar(255),
    name varchar(255) not null default '',
    full_address varchar(255) not null default '',
    city varchar(255) not null default '',
    account_number varchar(255) not null,
    created_at timestamp without time zone default CURRENT_TIMESTAMP,
    updated_at timestamp without time zone default CURRENT_TIMESTAMP,
    FOREIGN KEY (company_id) REFERENCES company_info (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS company_qr
(
    id varchar(255) not null primary key,
    code varchar(30000) not null,
    type varchar(255) not null default 'STATIC',
    created_at timestamp without time zone
);

CREATE TABLE IF NOT EXISTS managers
(
    id UUID not null primary key,
    selling_point_id UUID not null,
    profile_id varchar(255) not null,
    created_at timestamp without time zone default CURRENT_TIMESTAMP,
    FOREIGN KEY (selling_point_id) REFERENCES selling_point(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS company_services
(
    id UUID not null primary key,
    company_id uuid not null,
    qr_id varchar(255),
    city varchar(255), 
    service_type varchar(255), 
    amount_receive_type varchar(255), 
    account_number varchar(255),
    otp_verified boolean,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    FOREIGN KEY (company_id) REFERENCES company_info (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS mundea_contracts
(
    id UUID not null primary key,
    company_colvir_code varchar(255) not null,
    cms_type varchar(255) not null default '',
    account_number varchar(255) not null default '',
    dcl_code varchar(255) not null default '',
    dep_code varchar(255) not null default '',
    form_code varchar(255) not null default '',
    pay_code varchar(255) not null default '',
    pay_type int not null default 0,
    trf_code varchar(255) not null default ''
)