-- +goose Up
-- +goose StatementBegin
create table if not exists public.users 
(
    id uuid primary key not null default uuid_generate_v4(),
    last_name varchar(255) not null,
    first_name varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists public.users;
-- +goose StatementEnd
