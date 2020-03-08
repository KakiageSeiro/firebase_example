drop table if exists todos;
create table if not exists todos
(
  id      int unsigned not null primary key auto_increment,
  user_id    int not null,
  text    varchar(256) not null
);

create table if not exists users
(
    id      int unsigned not null primary key auto_increment,
    name    varchar(128) not null
);
