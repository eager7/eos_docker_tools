.PHONY: all run clean distclean

repository:="plainchant"

default: all

all:
	docker build -t ${repository}/eos_docker_tools:v1 ./

builder:
	docker build --target builder -t builder:v1 ./

run:
	docker run -p 13141:13141 --name eos_docker_tools ${repository}/eos_docker_tools:v1

clean:
	docker stop eos_docker_tools
	docker rm eos_docker_tools

distclean:
	-docker stop `docker ps -a -q`
	-docker rmi ${repository}/eos_docker_tools:v1
	-docker rmi builder:v1
	-docker rmi $(docker images | grep "none" | awk '{print $3}')

