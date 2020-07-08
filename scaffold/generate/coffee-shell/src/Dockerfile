FROM alpine:latest

COPY . .

RUN chmod +x set_umask.sh
RUN chmod +x main.sh
RUN mkdir app

ENTRYPOINT ["./set_umask.sh"]
CMD ["./shell.sh"]