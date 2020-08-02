
FROM alpine:latest
USER root

COPY . .
    
RUN apk update
RUN apk fetch openjdk8
RUN apk add openjdk8

ENV JAVA_HOME=/usr/lib/jvm/java-1.8-openjdk
ENV PATH="$JAVA_HOME/bin:${PATH}"

RUN chmod +x set_umask.sh

WORKDIR /app

ENTRYPOINT ["../set_umask.sh"]


CMD ["java -jar ../Main.jar"]