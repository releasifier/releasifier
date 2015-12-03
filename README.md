

```
docker run -d -p 5432:5432 -v $pwd/bin/data:/var/lib/postgresql/data --restart=always --name postgres -e POSTGRES_PASSWORD=betame -e POSTGRES_USER=ali postgres
```

docker run -it --rm --link postgres:postgres postgres sh -c 'exec psql -h "$POSTGRES_PORT_5432_TCP_ADDR" -p "$POSTGRES_PORT_5432_TCP_PORT" -U postgres'
