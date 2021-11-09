FROM golang:1.17 AS build

WORKDIR /compose/

COPY usermgmt.pb.go usermgmt.pb.go
COPY usermgmt_grpc.pb.go usermgmt_grpc.pb.go

FROM scratch
COPY --from=build /compose/usermgmt /usr/local/bin/usermgmt
CMD ["/usr/local/bin/usermgmt"]
