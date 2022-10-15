Create table users(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);
create table diary_list(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255)
);

CREATE table user_list(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    list_id int references diary_list(id) on delete cascade not null
);
create table diary_items(
    id     serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    body text not null
);
create table items_list
(
 id serial not null unique,
 item_id int references diary_items(id) on delete cascade not null,
 list_is int references diary_list(id) on delete cascade not null
);