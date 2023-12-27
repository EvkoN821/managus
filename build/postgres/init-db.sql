create table Managers(
    id serial not null primary key,
    name varchar
);

create table Clients(
    id serial not null primary key,
    manager_id integer not null references Managers(id) on delete cascade,
    name varchar
);

create table Invoices(
    id serial not null primary key,
    client_id integer not null references Clients(id) on delete cascade,
    invoice_id varchar,
    cont_date date,
    exec_date date,
    sum_total int,
    handed varchar,
    accepted varchar,
    add_info varchar,
    basis_doc varchar
);

create table Users(
    id serial not null primary key,
    login varchar,
    pwd varchar,
    user_type int
);

insert into users (login, pwd, user_type) values ('admin', 'admin', 1), ('user', 'user', 0);
