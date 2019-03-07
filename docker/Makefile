.PHONY: all clean distclean

default: all

all:
	docker build -t plainchant/eos/eos_docker_tools:v1 ./

run:
	docker run -d --name eos_docker_tools plainchant/eos/eos_docker_tools:v1

clean:
	docker stop eos_docker_tools
	docker rm eos_docker_tools

distclean:
	docker rmi eos_docker_tools
