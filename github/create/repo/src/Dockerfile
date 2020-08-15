FROM alpine:latest

RUN apk add --update bash git curl

COPY . .

RUN chmod +x set_umask.sh
RUN chmod +x repo.sh
RUN mkdir app

ENV DOCKER_EXECUTION=true
ENV CURRENT_PWD=/app

ENTRYPOINT ["./set_umask.sh"]
CMD ["./repo.sh"]