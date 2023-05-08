CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE blog_list
(
    id serial not null unique,
    title varchar(255) not null,
    tag varchar(255)
);

CREATE TABLE users_lists
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    list_id int references users(id) on delete cascade not null

);

