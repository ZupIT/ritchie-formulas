@echo off
docker build . -t precommit -f preCommit.Dockerfile
docker run -it -v %cd%:/app -v %userprofile%\.precommit\.cache:/root/.cache precommit
