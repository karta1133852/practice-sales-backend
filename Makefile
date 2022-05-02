APP_NAME=practice-sales-backend
DB_NAME=db-test-img

build: setup
	GOARCH=amd64 GOOS=windows go build -o bin/${APP_NAME} main.go
# GOARCH=amd64 GOOS=linux go build -o ${APP_NAME}-linux main.go

build-linux: setup
	GOARCH=amd64 GOOS=linux go build -o bin/${APP_NAME} main.go

gin:
	gin -p 8080 -a 3000

run:
	bin/${APP_NAME}

setup:
	mkdir -p bin

clean:
	go clean
	rm -rf build

docker: clean-docker build-docker run-docker

build-docker:
	docker build -t ${APP_NAME} .

run-docker:
	docker run -it --name ${APP_NAME} --net=container:${DB_NAME} ${APP_NAME}

clean-docker:
	docker container rm -f ${APP_NAME}
	docker image rm -f ${APP_NAME}

docker-compose:
	docker-compose up -d --build

run-docker-compose:
	docker-compose up -d