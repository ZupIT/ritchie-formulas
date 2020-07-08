
FROM node:10

COPY . .

RUN chmod +x set_umask.sh

WORKDIR /app

ENTRYPOINT ["/set_umask.sh"]
CMD ["node /index.js"]
