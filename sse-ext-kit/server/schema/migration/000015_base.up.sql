ALTER TABLE mundea_contracts
    DROP COLUMN IF EXISTS pay_type,
    ADD COLUMN IF NOT EXISTS description varchar(255) NOT NULL default '',
    ADD COLUMN IF NOT EXISTS knp varchar(255) NOT NULL default '',
    ADD COLUMN IF NOT EXISTS kbe varchar(255) NOT NULL default '',
    ADD COLUMN IF NOT EXISTS oked varchar(255) NOT NULL default '';
