FROM alpine:latest

RUN apk add --update ca-certificates openssl git openssh-client bash

COPY . .

RUN chmod +x set_umask.sh
RUN chmod +x fork.sh
RUN mkdir app

ENV CURRENT_PWD=/app

ENTRYPOINT ["/set_umask.sh"]

CMD ["/fork.sh"]
