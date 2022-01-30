FROM golang:latest as builder

RUN CGO_ENABLED=0 GOOS=linux go get -u github.com/jackc/tern

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/tern /
COPY ./dist/rappelconso /
COPY ./dist/migrations /migrations

EXPOSE 8080

CMD ["/rappelconso"]
