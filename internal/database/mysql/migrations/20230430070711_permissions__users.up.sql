CREATE TABLE permissions__users (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    ptype char(1) NOT NULL,
    v0 varchar(255),
    v1 varchar(255),
    v2 varchar(255),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
