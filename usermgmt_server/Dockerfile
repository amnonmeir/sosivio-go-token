FROM golang:1.17 AS build

WORKDIR /compose/

COPY usermgmt/ .
COPY usermgmt_server/ .

WORKDIR /compose/usermgmt_server/
RUN CGO_ENABLED=0 go build -o usermgmt_server usermgmt_server.go

FROM scratch
COPY --from=build /compose/usermgmt_server /usr/local/bin/usermgmt_server
CMD ["/usr/local/bin/usermgmt_server"]
