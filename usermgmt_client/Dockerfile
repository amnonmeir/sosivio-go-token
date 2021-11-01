FROM golang:1.17 AS build

WORKDIR /compose/

COPY usermgmt/* ./
COPY usermgmt_client/* ./
WORKDIR /compose/usermgmt_client/

RUN CGO_ENABLED=0 go build -o usermgmt_client usermgmt_client.go

FROM scratch
COPY --from=build /compose/usermgmt_client /usr/local/bin/usermgmt_client
CMD ["/usr/local/bin/usermgmt_client"]
