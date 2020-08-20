FROM circleci/python:latest

USER root

RUN mkdir /app
WORKDIR /app

RUN apt-get update --fix-missing
RUN apt-get install -y default-jdk
RUN apt-get install -y shellcheck
RUN pip install pre-commit

ENTRYPOINT ["pre-commit", "run", "--all-files"]
