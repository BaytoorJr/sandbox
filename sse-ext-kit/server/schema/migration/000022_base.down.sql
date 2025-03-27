ALTER TABLE remote_payment
    DROP COLUMN IF EXISTS qr_id,
    DROP COLUMN IF EXISTS transaction_id;