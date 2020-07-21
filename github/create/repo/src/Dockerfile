FROM alpine:latest

COPY . .

RUN apk add --update ca-certificates openssl git openssh-client bash curl

RUN chmod +x set_umask.sh
RUN chmod +x repo.sh
RUN mkdir app

WORKDIR /app

ENV EXECUTION_DOCKER=true

ENV CURRENT_PWD=/app

ENTRYPOINT ["/set_umask.sh"]
CMD ["/repo.sh"]