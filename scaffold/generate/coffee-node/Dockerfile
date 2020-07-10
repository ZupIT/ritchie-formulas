
FROM node:10

COPY . /app

RUN chmod +x /app/set_umask.sh

WORKDIR /app/src

ENTRYPOINT ["/app/set_umask.sh"]
CMD ["node index.js"]
