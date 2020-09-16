create table user
(
    openId    varchar(100) not null primary key unique,
    username  varchar(255),
    create_time timestamp
)