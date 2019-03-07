.PHONY: all run clean distclean

default: all

all:
	docker build -t plainchant/eos/eos_docker_tools:v1 ./

builder:
	docker build --target builder -t builder:v1 ./

run:
	docker run --name eos_docker_tools plainchant/eos/eos_docker_tools:v1

clean:
	docker stop eos_docker_tools
	docker rm eos_docker_tools

distclean:
	-docker rmi plainchant/eos/eos_docker_tools:v1
	-docker rmi builder:v1

