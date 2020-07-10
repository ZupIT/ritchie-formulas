
FROM golang:alpine AS builder

ADD . /app
WORKDIR /app/src
RUN go build -o main -v main.go

FROM alpine:latest


COPY --from=builder /app/src/main main
COPY --from=builder /app/set_umask.sh set_umask.sh
RUN chmod +x main
RUN chmod +x set_umask.sh

WORKDIR /app
ENTRYPOINT ["/set_umask.sh"]
CMD ["/main"]