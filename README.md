deploy.go
=========

https://github.com/andrerocker/deploy.go

### Description

- Read a YAML and export this as a 'Command as a Service' :p


##### Example

Based on a simple yaml
```yaml
daemon:
  bind: 127.0.0.1
  port: 8888
  load:
    - /var/www/*/config/deploy.yml
      /home/*/routines/deploy.yml

commands:
  process:
    - get: ps -ef | grep {process}
      put: kill {process}
      delete: kill -9 {process}

  log:
    - get: tail -f {log}
```

You can do this
```console
$ curl http://server:8888/process/ruby
andrero+ 1337 42   5  11:18 pts/25   00:00:01 ruby bin/rails s
andrero+ 1338 42   29 11:18 pts/26   00:00:01 ruby bin/rails c

$ curl -X PUT http://server:8888/process/1338
$

$ curl http://server:8888/process/ruby
andrero+ 1337 42   5  11:18 pts/25   00:00:01 ruby bin/rails s
```

```console
$ curl http://server:8888/log/log/development.log
Started GET "/document/42" for 127.0.0.1 at 2014-11-30 03:20:32 -0200
Processing by DocumentController#show as HTML
  Parameters: {"id"=>"42"}
  User Load (0.4ms)  SELECT  "users".* FROM "users"  WHERE "users"."id" = 1337
  Doc Load (0.6ms)  SELECT  "docs".* FROM "docs"  WHERE "docs"."id" = 42 LIMIT 1
  Rendered document/show.html.erb within layouts/application (0.2ms)
Completed 200 OK in 97ms (Views: 94.6ms | ActiveRecord: 1.0ms)
```


##### TODO List

- Use a better implementation of argparse
- Add support to before filters
- Add debian package and bricky
- Add experients with commands using - and resquest body as a arbitrary content
- Add a better implementation of command (process never end bug)
