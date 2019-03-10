#/bin/bash
# This is how we want to name the binary output
TARGET=eos_docker_tools
SRC=main.go
# These are the values we want to pass for Version and BuildTime
GITTAG=1.0.0
BUILD_TIME=`date +%Y%m%d%H%M%S`
# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${GITTAG} -X main.Build_Time=${BUILD_TIME} -s -w"

default:
	GO111MODULE=off go build ${LDFLAGS} -o ${TARGET} ${SRC}

mod:
	GO111MODULE=on go build ${LDFLAGS} -o ${TARGET} ${SRC}

clean:
	-rm ${TARGET}
