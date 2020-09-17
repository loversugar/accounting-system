create table user
(
    openId    varchar(100) not null primary key unique,
    username  varchar(255),
    create_time timestamp
)

create table bill (
    id int(10) not null primary key auto_increment unique,
    open_id varchar(255) not null,
    username varchar(255),
    consumption decimal(6, 2),
    create_time timestamp
)

create table bill_category (
    bill_id int not null,
    category_id int not null,
    primary key (bill_id, category_id)
)
