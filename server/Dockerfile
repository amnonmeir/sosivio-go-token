FROM golang:1.17 AS build

WORKDIR /server/

COPY usermgmt_server/* .

RUN mkdir /usermgmt
COPY usermgmt/* usermgmt

RUN go mod download
RUN CGO_ENABLED=0 go build -o server

FROM scratch
COPY --from=build /server/usermgmt_server /usr/local/bin/server
COPY --from=build /server/usermgmt /usr/local/bin/server/usermgmt
CMD ["/usr/local/bin/server"]