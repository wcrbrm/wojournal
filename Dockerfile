FROM golang:1.12-alpine as build
RUN apk add --no-cache git
RUN go get -u "github.com/jmoiron/sqlx"
RUN go get -u "github.com/lib/pq"
WORKDIR /go/src/github.com/wcrbrm/wojournal
COPY *.go ./
RUN CGO_ENABLED=0 go build -a --ldflags "-s -w" -o /usr/bin/wojournal

FROM alpine:3.9
COPY --from=build /usr/bin/wojournal /root/
COPY web/ ./web/
EXPOSE 9091
WORKDIR /root/
CMD ["/root/wojournal"]