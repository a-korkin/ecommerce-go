-- +goose Up
-- +goose StatementBegin
alter table public.products alter column price set default 0;
alter table public.products alter column price set not null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table public.products alter column price drop not null;
alter table public.products alter column price drop default;
-- +goose StatementEnd
