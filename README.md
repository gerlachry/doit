# doit
A very basic todo command line application.  Uses Cobra as a CLI library and sqlite as a backend database instead of a file based approach.

The first time running the CLI will generate .doit.db sqlite database file in users home director by default.

## Examples
- list help
```
doit 
A bar bones todo cli application

Usage:
  doit [command]

Available Commands:
  add         Add a new todo task
  done        Mark a task as done using the ID
  help        Help about any command
  list        List the current outstanding tasks

Flags:
      --config string   config file (default is $HOME/.doit.yaml)
  -h, --help            help for doit
  -t, --toggle          Help message for toggle

Use "doit [command] --help" for more information about a command.

```
- add a new task
```
doit add "take out the trash" --project=home --priority=3
added task
doit list
+----+----------+---------+---------------------------+
| ID | PRIORITY | PROJECT |           TASK            |
+----+----------+---------+---------------------------+
|  2 | H        |         | add a readme              |
|  1 | L        |         | figure out github release |
|  3 | L        | home    | take out the trash        |
+----+----------+---------+---------------------------+
```
- list outstanding tasks for specific project
```
doit list --project=home
+----+----------+---------+--------------------+
| ID | PRIORITY | PROJECT |        TASK        |
+----+----------+---------+--------------------+
|  3 | L        | home    | take out the trash |
+----+----------+---------+--------------------+
```
