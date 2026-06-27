# gRPC URL Shortener

Small Go gRPC service that stores URL mappings in memory.

## What it does

The service exposes two RPCs:

- `CreateLink` accepts a URL and returns a short code.
- `GetLink` accepts a short code and returns the original URL.

Current behavior is intentionally simple:

- Storage is in-memory only, so all data is lost on restart.
- Short codes are generated from the URL length, for example `link18`.
- gRPC reflection is not enabled.

## Run

```bash
go run ./cmd/link-service
```

The server listens on `localhost:50051`.

## API

Proto definition: [proto/link.proto](proto/link.proto)

```proto
service LinkService {
  rpc CreateLink(CreateLinkRequest) returns (CreateLinkResponse);
  rpc GetLink(GetLinkRequest) returns (GetLinkResponse);
}
```

Requests and responses:

- `CreateLinkRequest { string url = 1; }`
- `CreateLinkResponse { string short_url = 1; }`
- `GetLinkRequest { string short_url = 1; }`
- `GetLinkResponse { string url = 1; }`

## Example with grpcurl

Because reflection is disabled, point `grpcurl` at the proto file:

```bash
grpcurl -plaintext -import-path proto -proto link.proto \
	-d '{"url":"https://google.com"}' \
	localhost:50051 link.LinkService/CreateLink
```

```bash
grpcurl -plaintext -import-path proto -proto link.proto \
	-d '{"short_url":"link18"}' \
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
