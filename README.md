# gohex

Products are stored in `db.sqlite` in the process working directory. That file is not committed. The first time the program starts, it creates `db.sqlite` and a `products` table (`id`, `name`, `price`, `status`) when either is missing. A later start leaves an existing table and its rows in place.

## Run locally

Use Go 1.27 or newer and a C compiler. The SQLite driver is `github.com/mattn/go-sqlite3`, which needs cgo.

```bash
go run . cli -a create -n "Notebook" -p 19.9
```

The command prints the new product id. The row is in `./db.sqlite`.

## Docker Compose

`docker compose up --build` builds the Go 1.27.3 image and bind-mounts this repository at `/go/src`. The container stays up with `tail -f /dev/null`. The first command below creates `db.sqlite` on that mount, so the file shows up in this directory on the host:

```bash
docker compose up --build -d
docker compose exec app go run . cli -a create -n "Notebook" -p 19.9
```
