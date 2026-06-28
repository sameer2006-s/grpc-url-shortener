# gRPC URL Shortener

A simple Go gRPC service for shortening URLs using PostgreSQL.

## What it does

- `CreateLink(url)` - Returns a short code for a URL
- `GetLink(short_code)` - Returns the original URL

## Setup

Set `DATABASE_URL` in `.env`:
```bash
DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
```

## Run

```bash
go run ./cmd/link-service
```

The server listens on `localhost:50051`.

## API

```proto
service LinkService {
  rpc CreateLink(CreateLinkRequest) returns (CreateLinkResponse);
  rpc GetLink(GetLinkRequest) returns (GetLinkResponse);
}

message CreateLinkRequest {
  string url = 1;
}

message CreateLinkResponse {
  string short_url = 1;
}

message GetLinkRequest {
  string short_url = 1;
}

message GetLinkResponse {
  string url = 1;
}
```

See [proto/link.proto](proto/link.proto) for details.

## Example

Create a short link:
```bash
grpcurl -plaintext -import-path proto -proto link.proto \
  -d '{"url":"https://google.com"}' \
  localhost:50051 link.LinkService/CreateLink
```

Get the original URL:
```bash
grpcurl -plaintext -import-path proto -proto link.proto \
  -d '{"short_url":"link-xxx"}' \
  localhost:50051 link.LinkService/GetLink
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
