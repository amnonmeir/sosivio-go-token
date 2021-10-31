FROM golang:1.13 AS build

WORKDIR /compose/
COPY usermgmt_server.go usermgmt_server.go
COPY usermgmt.proto usermgmt.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative usermgmt.proto
RUN CGO_ENABLED=0 go build -o backend usermgmt_server.go

FROM scratch
COPY --from=build /compose/usermgmt_server /usr/local/bin/usermgmt_server
CMD ["/usr/local/bin/usermgmt_server"]
