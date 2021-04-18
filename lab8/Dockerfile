FROM golang:1.15-alpine AS build

WORKDIR /src/
COPY webserver.go /src/

COPY mongo.go /src/

COPY go.mod .
COPY go.sum .

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/webserver

RUN CGO_ENABLED=0 go build -o /bin/mongo

#FROM scratch
#COPY --from=build /bin/webserver /bin/webserver
#ENTRYPOINT ["/bin/server"]