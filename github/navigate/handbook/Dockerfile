FROM alpine:3.12
USER root

RUN mkdir /rit
COPY . /rit
RUN chmod +x /rit/set_umask.sh

WORKDIR /app
ENTRYPOINT ["/rit/set_umask.sh"]
CMD ["/rit/run.sh"]
