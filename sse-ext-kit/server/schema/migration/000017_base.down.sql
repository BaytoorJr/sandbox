DO $$
    BEGIN
        IF EXISTS (
            SELECT 1
            FROM pg_constraint
            WHERE conname = 'unique_profile_selling_point'
              AND conrelid = 'managers'::regclass
        ) THEN
            ALTER TABLE managers
                DROP CONSTRAINT unique_profile_selling_point;
        END IF;
    END $$;