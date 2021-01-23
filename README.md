## go-httpserver
web server for serving static site.

### Usage
run go-httpserver with default options (go-httpserver will serving current directoy and running on port 8080)
```
go-httpserver
```

run go-httpserver with spesific port
```
go-httpserver -port 9000
```

run go-httpserver and serving `static` directory
```
go-httpserver -dir static
```

run go-httpserver with Docker
```
docker run -d -v <directory>:/static -p 8080:8080 nothinux/go-httpserver:latest
```

### LICENSE
view [LICENSE](https://github.com/nothinux/go-httpserver/blob/main/LICENSE) file