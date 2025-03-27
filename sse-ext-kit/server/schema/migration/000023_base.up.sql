CREATE OR REPLACE FUNCTION notify_payment_row_change()
    RETURNS trigger AS $$
DECLARE
    payload JSON;
BEGIN
    payload := json_build_object(
            'operation', TG_OP,
            'id', COALESCE(NEW.id, OLD.id),  -- Use OLD.id for DELETE operations
            'status', COALESCE(NEW.status, OLD.status),  -- Use NEW.status for INSERT/UPDATE, OLD.status for DELETE
            'qrId', COALESCE(NEW.qr_id, OLD.qr_id),
            'type', COALESCE(NEW.type, OLD.type),
            'companyId', COALESCE(NEW.company_id, OLD.company_id),
            'accountNumber', COALESCE(NEW.account_number, OLD.account_number),
            'amount', COALESCE(NEW.amount, OLD.amount),
            'payerShortName', COALESCE(NEW.payer_short_name, OLD.payer_short_name),
            'payerPhoneNumber', COALESCE(NEW.payer_phone_number, OLD.payer_phone_number),
            'description', COALESCE(NEW.description, OLD.description),
            'serviceType', COALESCE(NEW.service_mapping_type, OLD.service_mapping_type),
            'transactionId', COALESCE(NEW.transaction_id, OLD.transaction_id),
            'table', TG_TABLE_NAME,
            'createdAt', COALESCE(NEW.created_at, OLD.created_at)
               );

    -- Notify with the constructed payload
    PERFORM pg_notify('payment_row_change_channel', payload::text);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER payment_row_change
    AFTER INSERT OR UPDATE OR DELETE
    ON remote_payment
    FOR EACH ROW
EXECUTE FUNCTION notify_payment_row_change();