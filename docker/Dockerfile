FROM golang:1.12-alpine as builder

RUN apk --no-cache add git

RUN go get -d -v github.com/eager7/eos_docker_tools/

WORKDIR /go/src/github.com/eager7/eos_docker_tools/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o eos_docker_tools .

ENTRYPOINT [""]

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /eos_docker_tools/

COPY --from=0 /go/src/github.com/eager7/eos_docker_tools/eos_docker_tools .

CMD ["./eos_docker_tools"]