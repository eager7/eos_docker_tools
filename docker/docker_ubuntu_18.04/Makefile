.PHONY: all run clean dist_clean

repository:="plainchant/eos_docker_tools:ubuntu-18.04"

default: all

all:
	docker build -t ${repository} ./

builder:
	docker build --target builder -t builder:v1 ./

run:
	docker run -p 13141:13141 --name eos_docker_tools ${repository}

clean:
	docker stop eos_docker_tools
	docker rm eos_docker_tools

dist_clean:
	-docker rmi $$(docker images | grep "none" | awk '{print $$3}')
	-docker stop `docker ps -a -q`
	-docker rmi ${repository}
	-docker rmi builder:v1

exec:
    docker exec -it ${repository} sh

test:
	docker images | grep "none" | awk '{print $$3}'
