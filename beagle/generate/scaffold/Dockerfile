FROM cimg/base:stable-20.04

USER root

COPY . /

RUN sed -i 's/\r//g' /set_umask.sh
RUN sed -i 's/\r//g' /run.sh
RUN chmod +x /set_umask.sh
RUN mkdir /app

ENV DOCKER_EXECUTION=true
ENV CURRENT_PWD=/app

WORKDIR /

ENTRYPOINT ["/set_umask.sh"]
CMD ["/run.sh"]