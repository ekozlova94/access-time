create table access_time
(
    id   integer  not null primary key autoincrement,
    ip   integer  not null,
    time datetime not null
);