FROM golang:1.13 AS build

WORKDIR /compose/
COPY usermgmt_client.go usermgmt_client.go
COPY usermgmt.proto usermgmt.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative usermgmt.proto
RUN CGO_ENABLED=0 go build -o frontend usermgmt_client.go

FROM scratch
COPY --from=build /compose/usermgmt_client /usr/local/bin/usermgmt_client
CMD ["/usr/local/bin/usermgmt_client"]
