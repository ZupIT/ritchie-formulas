FROM alpine:latest

COPY . /app

RUN chmod +x /app/set_umask.sh
RUN chmod +x /app/src/main.sh

WORKDIR /app/src

ENTRYPOINT ["/app/set_umask.sh"]
CMD ["./main.sh"]
