FROM ruby:2.6

COPY . .

RUN chmod +x set_umask.sh

WORKDIR /app

ENTRYPOINT ["/set_umask.sh"]
CMD ["./index.rb"]
