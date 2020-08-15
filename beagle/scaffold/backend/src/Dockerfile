FROM alpine:latest

RUN apk add --update bash findutils

COPY . .

RUN chmod +x set_umask.sh
RUN chmod +x backend.sh
RUN mkdir app

ENV DOCKER_EXECUTION=true
ENV CURRENT_PWD=/app

ENTRYPOINT ["./set_umask.sh"]
CMD ["./backend.sh"]