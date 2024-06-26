CREATE TABLE makves_data (
    internal_id INT,
    id INT,
    uid TEXT,
    domain TEXT,
    cn TEXT,
    department TEXT,
    title TEXT,
    who TEXT,
    logon_count INT,
    num_logons7 INT,
    num_share7 INT,
    num_file7 INT,
    num_ad7 INT,
    num_n7 INT,
    num_logons14 INT,
    num_share14 INT,
    num_file14 INT,
    num_ad14 INT,
    num_n14 INT,
    num_logons30 INT,
    num_share30 INT,
    num_file30 INT,
    num_ad30 INT,
    num_n30 INT,
    num_logons150 INT,
    num_share150 INT,
    num_file150 INT,
    num_ad150 INT,
    num_n150 INT,
    num_logons365 INT,
    num_share365 INT,
    num_file365 INT,
    num_ad365 INT,
    num_n365 INT,
    has_user_principal_name INT,
    has_mail INT,
    has_phone INT,
    flag_disabled INT,
    flag_lockout INT,
    flag_password_not_required INT,
    flag_password_cant_change INT,
    flag_dont_expire_password INT,
    owned_files INT,
    num_mailboxes INT,
    num_member_of_groups INT,
    num_member_of_indirect_groups INT,
    member_of_indirect_groups_ids TEXT,
    member_of_groups_ids TEXT,
    is_admin INT,
    is_service INT
);

COPY makves_data FROM '/var/lib/postgresql/csv/ueba.csv' (FORMAT csv, DELIMITER ',', HEADER);
