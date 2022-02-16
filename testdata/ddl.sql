create table users (
    id integer primary key,
    name varchar(255) not null,
    division varchar (255) not null,
    joined_at timestamp not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp
)
