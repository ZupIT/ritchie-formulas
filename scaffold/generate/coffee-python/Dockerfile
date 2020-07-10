FROM python:3

COPY . /app

RUN chmod +x /app/set_umask.sh

WORKDIR /app/src

ENTRYPOINT ["/app/set_umask.sh"]
CMD ["python3 main.py"]