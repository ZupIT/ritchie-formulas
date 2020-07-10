FROM maven:3.6.3-jdk-8 AS builder

ADD . /app
WORKDIR /app
RUN mvn clean install


FROM alpine:latest
USER root

COPY --from=builder /app/target/Main.jar /app/Main.jar
COPY --from=builder /app/set_umask.sh /app/set_umask.sh

RUN apk update
RUN apk fetch openjdk8
RUN apk add openjdk8

ENV JAVA_HOME=/usr/lib/jvm/java-1.8-openjdk
ENV PATH="$JAVA_HOME/bin:${PATH}"

RUN chmod +x /app/set_umask.sh

WORKDIR /app

ENTRYPOINT ["./set_umask.sh"]

CMD ["java -jar Main.jar"]