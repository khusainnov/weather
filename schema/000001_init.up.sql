CREATE TABLE search_queries
(
    id      serial       not null,
    city    varchar(255) not null unique,
    counter int          not null
);