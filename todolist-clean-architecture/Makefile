load_test:
	echo "GET http://:8080/limit" | vegeta attack -rate=10/s -duration=1s | vegeta report

build:
	go build \
		-ldflags "-X main.buildcommit=`git rev-parse --short HEAD` \
		-X main.buildtime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" \
		-o app

maria:
	docker run -p 3306:3306 --name mariadb \
	-e MARIADB_ROOT_PASSWORD=password -e MARIADB_DATABASE=myapp -d mariadb:latest

image:
	docker build -t todo:test -f Dockerfile .

container:
	docker run -p 8080:8080 --env-file ./local.env --link mariadb:db \
	--name myapp todo:test