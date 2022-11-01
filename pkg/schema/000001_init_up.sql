Create table users (
    id bigserial primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    deleted_at timestamptz default current_timestamp
);
----------------------------------------------------------------------------------------------------------------
create table item(
    id     bigserial primary key,
    title varchar(255) not null,
    description varchar(255),
    body text not null,
    created timestamptz not null default current_timestamp
);

----------------------------------------------------------------------------------------------------------------
create table citations(
    id bigserial primary key,
    author varchar(255) not null,
    citation text not null
);
----------------------------------------------------------------------------------------------------------------
create table categories(
     id bigserial primary key,
     name varchar(50) not null
);
--------------------------------------------------------------------------------------------------------------

create table bad_habit(
    id bigserial primary key,
    habit_category_id int references categories on delete cascade,
    equivalent_id int references categories on delete cascade,
    created timestamptz not null default current_timestamp
);

-------------------------------------------------------------------------------------------
create table mental_course(
    id bigserial primary key,
    mental_category_id int references categories on delete cascade,
    created timestamptz not null default current_timestamp
);

-----------------------------------------------------------------------------------------------------------
create table body_course(
    id bigserial primary key,
    level_id int references categories on delete cascade,
    train_category_id int references categories on DELETE cascade,
    created timestamptz not null default current_timestamp
);
----------------------------------------------------------------------------------------------------------

create table users_space(
    id bigserial primary key,
    user_id int references users on delete cascade,
    item_id int references item on delete cascade,
    bad_habit_id int references bad_habit on delete cascade,
    mental_course_id int references mental_course on delete cascade,
    body_course_id int references body_course on delete cascade
);

























-- create table gender(
--     id bigserial primary key,
--     sex varchar(50) not null
-- );
-- create table family_status(
--     id bigserial primary key,
--     status varchar(50)
-- );
-- create table temperament(
--     id bigserial primary key,
--     name varchar(50) not null
-- );
-- create table public(
--     id bigserial primary key,
--     age int not null,
--     gender_id int not null references gender,
--     family_status_id int not null references family_status,
--     temperament_id int not null references temperament,
--     goal_to_life varchar(250) not null,
--     big_fear varchar(255),
--     created timestamptz not null default current_timestamp
-- );
-- create table public_list
-- (
--     id        bigserial primary key,
--     user_id   int references users (id) on delete cascade  not null,
--     public_id int references public (id) on delete cascade not null
-- );