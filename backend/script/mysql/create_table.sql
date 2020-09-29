create table user
(
    id int not null primary key unique auto_increment,
    openId    varchar(100) not null unique,
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

create table category (
    id int not null primary key auto_increment unique,
    category_name varchar(100) not null,
    category_url varchar(255),
    user_id int null,
    is_private tinyint(1),
    deleted tinyint(1)
)
