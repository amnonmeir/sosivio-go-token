FROM golang:1.17 AS build

WORKDIR /client/

COPY usermgmt_client/* .

RUN mkdir /usermgmt
COPY usermgmt/* usermgmt

RUN go mod download
RUN CGO_ENABLED=0 go build -o client

FROM scratch
COPY --from=build /client/usermgmt_client /usr/local/bin/client
COPY --from=build /client/usermgmt /usr/local/bin/client/usermgmt
CMD ["/usr/local/bin/client"]
