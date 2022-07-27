CREATE TABLE search_queries
(
    id      serial       not null,
    name    varchar(255) not null unique,
    counter int          not null
);