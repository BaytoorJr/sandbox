alter table selling_point
    add column if not exists address_match_with_company bool default false;