# es-tasks-list

CLI tool to show ongoing ElasticSearch tasks in a fancy way

## Install

```
go get github.com/aleroyer/es-tasks-list
```

## Usage

```
Usage:
  es-tasks-list [flags]

Flags:
      --cancellable     Show if a task is cancellable
      --detailed        Show task's description
  -h, --help            help for es-tasks-list
  -s, --server string   ElasticSearch URL (required)
      --ssl             Use HTTPS to connect
      --start-time      Show task's start time
```

`-server` (`-s`) flag is mandatory. It can be provided in various form:
```
ip:port
domain.tld
domain.tld:port
```

You can filter on actions by giving string arguments. For example if you want to show only search tasks:
```
es-tasks-list -s localhost:9200 "*search"
```

## Output

```
$ es-tasks-list -s localhost:9200 "*search"
+---------------+--------------+----------------------------------+----------+
| NODE          | IP           | TASK_ID                          | DURATION |
+---------------+--------------+----------------------------------+----------+
| localhost     | 127.0.0.1    | xxxxxxxxxxxxxxxx:NNNNNNNNN       | 74.5ms   |
+---------------+--------------+----------------------------------+----------+
```
