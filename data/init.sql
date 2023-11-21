create schema todo;

create table todo.todo (
    id uuid primary key not null,
    priority text not null,
    name text not null,
    body text not null,
    deadline time,
    createdAt time,
    updatedAt time
);