create table drivers
(
    id         char(36)                   not null
        primary key,
    name       varchar(255)               not null,
    license_id varchar(255)               not null,
    created_at datetime   default (now()) null,
    updated_at datetime   default (now()) null on update CURRENT_TIMESTAMP,
    deleted    tinyint(1) default 0       null,
    constraint id
        unique (id)
);

create table trucks
(
    id            char(36)                 not null
        primary key,
    Brand         varchar(255)             not null,
    Model         varchar(255)             not null,
    Year          int                      not null,
    license_plate varchar(10)              not null,
    created_at    datetime default (now()) not null,
    updated_at    datetime default (now()) not null,
    deleted       tinyint  default 0       null,
    constraint LicensePlate
        unique (license_plate)
);

create table links
(
    id         char(36)                 not null
        primary key,
    truck_id   char(36)                 not null,
    driver_id  char(36)                 not null,
    created_at datetime default (now()) not null,
    updated_at datetime default (now()) not null,
    deleted    tinyint  default 0       not null,
    constraint links_ibfk_1
        foreign key (truck_id) references trucks (id),
    constraint links_ibfk_2
        foreign key (driver_id) references drivers (id)
);

create index drive_id
    on links (driver_id);

create index truck_id
    on links (truck_id);