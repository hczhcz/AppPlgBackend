create table logins (
 id bigserial primary key,
 email text,
 phone text);

create table users (
 id bigint not null references logins(id) on delete cascade on update cascade,
 nickname text not null,
 gender text not null,
 description text);
