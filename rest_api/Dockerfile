FROM golang:1.17 AS build

WORKDIR /compose/

COPY login.go login.go

RUN mkdir /tokens
COPY token.go token.go

RUN mkdir /data
COPY data.go data.go

RUN mkdir /person
COPY person.go person.go

RUN mkdir /car
COPY car.go car.go

RUN mkdir /connect_to_server
COPY makeConnection.go makeConnection.go

RUN go mod download
RUN CGO_ENABLED=0 go build -o rest_api login.go
ENV 
FROM scratch
COPY --from=build /compose/rest_api /usr/local/bin/rest_api
COPY --from=build /compose/rest_api/tokens /usr/local/bin/rest_api/tokens
COPY --from=build /compose/rest_api/data_base /usr/local/bin/rest_api/data_base
COPY --from=build /compose/rest_api/connect_to_server /usr/local/bin/rest_api/connect_to_server
CMD ["/usr/local/bin/rest_api"]
