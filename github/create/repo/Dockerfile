FROM cimg/base:stable-20.04

USER root

RUN mkdir /rit
COPY . /rit
RUN sed -i 's/\r//g' /rit/set_umask.sh
RUN sed -i 's/\r//g' /rit/run.sh
RUN chmod +x /rit/set_umask.sh

RUN  echo "[user]" > /root/.gitconfig
RUN  echo "	name = Your Name" >> /root/.gitconfig
RUN  echo "	email = you@example.com" >> /root/.gitconfig
RUN  echo "[url \"git@github.com:\"]" >> /root/.gitconfig
RUN  echo "  insteadOf = https://github.com/" >> /root/.gitconfig

WORKDIR /app
ENTRYPOINT ["/rit/set_umask.sh"]
CMD ["/rit/run.sh"]
