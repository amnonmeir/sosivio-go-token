FROM golang:1.13 AS build

WORKDIR /compose/
COPY usermgmt_client.go usermgmt_client.go
RUN CGO_ENABLED=0 go build -o frontend usermgmt_client.go

FROM scratch
COPY --from=build /compose/usermgmt_client /usr/local/bin/usermgmt_client
CMD ["/usr/local/bin/usermgmt_client"]
