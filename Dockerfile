FROM golang:bullseye

LABEL author="Paique"

RUN apt-get update -y

RUN DEBIAN_FRONTEND=noninteractive TZ=America/Sao_Paulo apt-get install -y curl wget unzip zip tzdata

RUN useradd -ms /bin/bash dcs-docker

USER dcs-docker
ENV  USER=dcs-docker HOME=/home/dcs-docker
WORKDIR     /home/dcs-docker

RUN mkdir -p /tmp/buildgo/dcs-api
COPY dcs-api /tmp/buildgo/dcs-api

RUN cd /tmp/buildgo/dcs-api \
    && go build -o $HOME/dcs-api-app \
    && chmod +x $HOME/dcs-api-app

CMD         ["/bin/bash", "-c", "$HOME/dcs-api-app"]
EXPOSE 8080