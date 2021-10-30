FROM golang:1.13 AS build

WORKDIR /compose/
COPY usermgmt.proto usermgmt.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative usermgmt.proto

FROM scratch
COPY --from=build /compose/usermgmt /usr/local/bin/usermgmt
CMD ["/usr/local/bin/usermgmt"]
