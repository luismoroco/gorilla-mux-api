APP=gorillaMuxAPI

BINARY=main.out
ROOT_FILE=main.go
DEV_DIR=tmp

.PHONY: dev 
dev: 
	bash -c "air"

.PHONY: format
format:
	bash -c "gofmt -w ." 

.PHONY: build
build: 
	go build -o ${BINARY} ${ROOT_FILE} 

.PHONY: test
test:
	go test -v ${ROOT_FILE} 

.PHONY: run
run: build
	./${BINARY}

.PHONY: clean
clean:
	go clean
	@if [ -f ${BINARY} ] ; then rm ${BINARY} && rm -rf ${DEV_DIR}/ ; fi

.PHONY: clean-docker 
clean-docker:
	sudo docker system prune 

.PHONY: setup-docker 
setup-docker:
	sudo docker-compose up -d 

.PHONY: down-docker 
down-docker:
	sudo docker-compose down