create database if not exists test;
use test;
create table if not exists data_tables(id int primary key, data varchar(50));
insert into data_tables values (1, "test_1"),(2, "test_2"),(3, "test_3"),(4, "test_4");