FROM golang:1.13 AS build

WORKDIR /compose/
COPY usermgmt_server.go usermgmt_server.go
RUN CGO_ENABLED=0 go build -o usermgmt_server usermgmt_server.go

FROM scratch
COPY --from=build /compose/usermgmt_server /usr/local/bin/usermgmt_server
CMD ["/usr/local/bin/usermgmt_server"]