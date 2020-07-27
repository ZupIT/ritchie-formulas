FROM alpine:latest

RUN apk add --update bash findutils git curl

COPY . .

RUN chmod +x set_umask.sh
RUN chmod +x android.sh
RUN mkdir app

ENV DOCKER_EXECUTION=true
ENV CURRENT_PWD=/app

ENTRYPOINT ["/set_umask.sh"]
CMD ["/android.sh"]