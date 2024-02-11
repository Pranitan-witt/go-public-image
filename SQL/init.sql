create database if not exists movieDB;

create table movieDB.Movies
(
    ID int auto_increment primary key,
    Title varchar(50),
    Director varchar(50)
);

insert into movieDB.Movies (Title, Director) values ('Avatar', 'Mr.A');
insert into movieDB.Movies (Title, Director) values ('Marvel', 'Stand Lees');
insert into movieDB.Movies (Title, Director) values ('ABC', 'Director');