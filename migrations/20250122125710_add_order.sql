-- +goose Up
-- +goose StatementBegin
create table if not exists public.orders
(
    id uuid primary key not null default uuid_generate_v4(),
    product_id uuid not null references public.products(id) on delete cascade,
    user_id uuid not null references public.users(id) on delete cascade,
    amount integer not null default 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete table if exists public orders;
-- +goose StatementEnd
