create table users (
	id serial primary key,
	username text not null unique,
	displayName text not null
);
