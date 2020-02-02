package todo

// InitDB Initialize a fresh database
const InitDB string = `
create table projects (
	id integer not null primary key,
	name text,
	created text default current_timestamp,
	modified text default current_timestamp
);

create table tasks (
  id integer not null primary key,
  name text,
  priority integer not null,
  completed integer,
  project_id integer references project(id),
  created text default current_timestamp,
  modified text default current_timestamp
)
`
