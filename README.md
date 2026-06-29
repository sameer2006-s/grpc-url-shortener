# gRPC URL Shortener

A Go URL shortener with a gRPC backend and an HTTP API gateway, backed by PostgreSQL.

## Architecture

```
HTTP clients  →  api-service (:8080)  →  link-service (:50051, gRPC)  →  PostgreSQL
```

| Service | Role | Port |
|---------|------|------|
| `link-service` | gRPC core — create links, resolve codes, track visits | 50051 |
| `api-service` | HTTP gateway in front of the gRPC service | 8080 |
| `postgres` | Persistent storage | 5433 (host) |

## What it does

**gRPC (`LinkService`)**

- `CreateLink(url)` — returns a short code for a URL
- `GetLink(short_code)` — returns the original URL
- `VisitLink(short_code)` — increments click count and returns the redirect URL
- `GetStats(short_code)` — returns click count and creation time

**HTTP (`api-service`)**

| Endpoint | Description |
|----------|-------------|
| `GET /shorten?url=<url>` | Create a short link; returns the short code |
| `GET /get?code=<code>` | Resolve a short code to the original URL |
| `GET /visit?code=<code>` | Redirect to the original URL (301) and increment clicks |
| `GET /stats?code=<code>` | Return click count and `created_at` as JSON |

## Quick start (Docker)

1. Copy and configure environment variables in `.env`:

```bash
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=shortener

LINK_SERVICE_ADDR=link-service:50051
```

2. Start everything (Postgres, migrations, both services):

```bash
docker compose up --build
```

Postgres is exposed on `localhost:5433`. The HTTP API is at `http://localhost:8080` and gRPC at `localhost:50051`.

## Local development

Run Postgres (or use `docker compose up postgres migrate`), apply migrations, then start each service in a separate terminal.

**Environment variables**

| Variable | Used by | Example (local) |
|----------|---------|-----------------|
| `DB_HOST` | link-service | `localhost` |
| `DB_PORT` | link-service | `5433` |
| `DB_USER` | link-service | `postgres` |
| `DB_PASSWORD` | link-service | `your_password` |
| `DB_NAME` | link-service | `shortener` |
| `LINK_SERVICE_ADDR` | api-service | `localhost:50051` |

**Run link-service**

```bash
go run ./services/link-service
```

**Run api-service**

```bash
go run ./services/api-service
```

**Apply migrations manually**

```bash
migrate -path ./migrations \
  -database "postgres://postgres:your_password@localhost:5433/shortener?sslmode=disable" \
  up
```

Requires [golang-migrate](https://github.com/golang-migrate/migrate).

## Examples

**HTTP**

```bash
# Create a short link
curl "http://localhost:8080/shorten?url=https://google.com"

# Resolve a code
curl "http://localhost:8080/get?code=link-abc12345"

# Visit (redirect)
curl -L "http://localhost:8080/visit?code=link-abc12345"

# Stats
curl "http://localhost:8080/stats?code=link-abc12345"
```

**gRPC** (requires [grpcurl](https://github.com/fullstorydev/grpcurl))

Create a short link:

```bash
grpcurl -plaintext -import-path proto -proto link.proto \
  -d '{"url":"https://google.com"}' \
  localhost:50051 link.LinkService/CreateLink
```

Get the original URL:

```bash
grpcurl -plaintext -import-path proto -proto link.proto \
  -d '{"short_url":"link-abc12345"}' \
  localhost:50051 link.LinkService/GetLink
```

Visit a link:

```bash
grpcurl -plaintext -import-path proto -proto link.proto \
  -d '{"short_url":"link-abc12345"}' \
  localhost:50051 link.LinkService/VisitLink
```

Get stats:

```bash
grpcurl -plaintext -import-path proto -proto link.proto \
  -d '{"short_url":"link-abc12345"}' \
  localhost:50051 link.LinkService/GetStats
```

## API definition

See [proto/link.proto](proto/link.proto) for the full protobuf schema.

```proto
service LinkService {
  rpc CreateLink(CreateLinkRequest) returns (CreateLinkResponse);
  rpc GetLink(GetLinkRequest) returns (GetLinkResponse);
  rpc VisitLink(VisitLinkRequest) returns (VisitLinkResponse);
  rpc GetStats(GetStatsRequest) returns (GetStatsResponse);
}
```

## Project layout

```
services/
  link-service/     gRPC server entrypoint
  api-service/      HTTP gateway entrypoint
internal/
  grpc/             gRPC handlers
  service/          business logic
  repository/       PostgreSQL access
  model/            domain types
proto/              protobuf definitions
gen/proto/          generated Go code
migrations/         SQL migrations
```

## Regenerate protobufs

```bash
protoc \
  --go_out=./gen \
  --go_opt=paths=source_relative \
  --go-grpc_out=./gen \
  --go-grpc_opt=paths=source_relative \
  proto/link.proto
```

## Requirements

- Go 1.26+
- Docker & Docker Compose (for containerized setup)
- PostgreSQL 18 (via Docker or local install)
