create table public.security
(
    id       uuid                      not null
        primary key,
    name     text                      not null,
    metadata jsonb default '{}'::jsonb not null
);

alter table public.security
    owner to postgres;

