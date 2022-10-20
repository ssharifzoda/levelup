Create table users (
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    registered timestamp not null default current_timestamp
);
create table item(
    id     serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    body text not null,
    registered timestamp not null default current_timestamp
);
create table items_list(
     id serial not null unique,
     user_id int references users(id) on delete cascade not null,
     item_id int references item(id) on delete cascade not null
);
create table citations(
    id serial not null unique,
    author varchar(255) not null,
    citation text not null
);
create table bad_habit(
    id serial not null unique,
    bad_habit varchar(255) not null,
    equivalent varchar(255) not null,
    status text,
    session time,
    registered timestamp not null default current_timestamp
);
create table bad_habits_list(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    bad_habit_id int references bad_habit(id) on delete cascade not null
);