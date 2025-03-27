-- -- Drop the unique constraint on the combination of company_id and issue_date
ALTER TABLE license
DROP CONSTRAINT IF EXISTS license_company_id_key;

ALTER TABLE license
ADD COLUMN IF NOT EXISTS guid VARCHAR(40) CONSTRAINT guid_unique UNIQUE;