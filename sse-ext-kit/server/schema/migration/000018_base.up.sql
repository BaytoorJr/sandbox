alter table selling_point
    add column if not exists street varchar(255) not null default '',
    add column if not exists flat_number varchar(255) not null default '',
    add column if not exists apartment_number varchar(255) not null default '';