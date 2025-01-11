-- +goose Up
-- +goose StatementBegin
create extension if not exists "uuid-ossp";

create table if not exists public.categories
(
    id uuid primary key not null default uuid_generate_v4(),
    title varchar(255) not null,
    code varchar(255) not null
);

create table if not exists public.products 
(
    id uuid primary key not null default uuid_generate_v4(),
    title varchar(255) not null,
    category uuid not null references public.categories(id),
    price numeric(10, 2)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists public.products;
drop table if exists public.categories;
-- +goose StatementEnd
