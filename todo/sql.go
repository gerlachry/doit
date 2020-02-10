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

// TaskSelectSQL select sql for outstanding tasks
const TaskSelectSQL string = `
select 
  t.id, 
  t.name, 
  priority,
  coalesce(p.id,0) as project_id,
  coalesce(p.name, "") as project_name
from tasks t
left outer join projects p on p.id = t.project_id
where completed = 0 
`

// TaskSelectOrderBySQL order by clause for outstanding tasks
const TaskSelectOrderBySQL string = `
order by priority asc, t.created asc
`
