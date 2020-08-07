FROM openjdk:14-alpine
COPY target/bff-micronaut-*.jar bff-micronaut.jar
EXPOSE 8080
CMD ["java", "-Dcom.sun.management.jmxremote", "-Xmx128m", "-jar", "bff-micronaut.jar"]
