FROM alpine:3.12

RUN mkdir /rit && \
    apk add --no-cache sudo curl bash nodejs && \
    curl -fsSL https://commons-repo.ritchiecli.io/install.sh | bash

COPY . /rit

RUN sed -i 's/\r//g' /rit/set_umask.sh && \
    sed -i 's/\r//g' /rit/run.sh && \
    chmod +x /rit/set_umask.sh

WORKDIR /app
ENTRYPOINT ["/rit/set_umask.sh"]
CMD ["/rit/run.sh"]
