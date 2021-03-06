FROM golang:1.12-alpine as builder

RUN apk --no-cache add git

RUN go get -d -v github.com/eager7/eos_docker_tools/

WORKDIR /go/src/github.com/eager7/eos_docker_tools/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o eos_docker_tools .

ENTRYPOINT [""]

FROM ubuntu:18.04 as prod

ARG release=latest

WORKDIR /eos_docker_tools/

COPY --from=0 /go/src/github.com/eager7/eos_docker_tools/eos_docker_tools .

# Install required packages
RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get -y install openssl ca-certificates curl wget make \
    && rm -rf /var/lib/apt/lists/*

# Install CDT from deb package
ADD install_deb.sh /

RUN /install_deb.sh $release && rm -f install_deb.sh

USER root

EXPOSE 13141

ENTRYPOINT ["./eos_docker_tools"]
