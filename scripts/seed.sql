insert into public.categories(title, code)
select 'category#' || a::text as title, 'cat#' || a::text as code
from generate_series(1, 100) as a;

insert into public.products(title, category, price)
select 
    'product#' || a::text as title, b.id as category,
    random() * 100 + 20 as price
from 
(
    select a, (random() * 100)::integer as c
    from generate_series(1, 10000) as a
) as a
inner join public.categories as b on b.code = 'cat#' || a.c;

insert into public.users(last_name, first_name)
values
    ('Ivanov', 'Ivan'),
    ('Petrov', 'Petr'),
    ('Sidorov', 'Sidr');
