FROM golang:1.17 AS build

WORKDIR /compose/

COPY login.go login.go

RUN mkdir -p /tokens
COPY tokens/token.go tokens/token.go

RUN mkdir -p /data
COPY data/data.go data/data.go

RUN mkdir -p /person
COPY person/person.go person/person.go

RUN mkdir -p /car
COPY car/car.go car/car.go

RUN mkdir -p /connect_to_server
COPY connect_to_server/makeConnection.go connect_to_server/makeConnection.go

RUN CGO_ENABLED=0 go build -o rest_api login.go

FROM scratch
COPY --from=build /compose/rest_api /usr/local/bin/rest_api
COPY --from=build /compose/rest_api/tokens /usr/local/bin/rest_api/tokens
COPY --from=build /compose/rest_api/data_base /usr/local/bin/rest_api/data_base
COPY --from=build /compose/rest_api/connect_to_server /usr/local/bin/rest_api/connect_to_server
CMD ["/usr/local/bin/rest_api"]
