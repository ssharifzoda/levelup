Create table users (
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);
create table item(
    id     serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    body text not null
);
create table items_list(
     id serial not null unique,
     user_id int references users(id) on delete cascade not null,
     item_id int references item(id) on delete cascade not null
  );