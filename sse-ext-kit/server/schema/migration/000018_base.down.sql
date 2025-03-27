alter table selling_point
    drop column if exists street,
    drop column if exists flat_number,
    drop column if exists apartment_number