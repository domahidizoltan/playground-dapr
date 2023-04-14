install: install-dapr-cli
start-app:
	dapr run --app-id app --app-port 3000 --dapr-http-port 8000 -- go run app/main.go

start-dependencies:
	docker start dapr_redis dapr_zipkin dapr_placement
	docker-compose up -d
	dapr dashboard

install-dapr-cli:
	wget -q https://raw.githubusercontent.com/dapr/cli/master/install/install.sh -O - | /bin/bash -s 1.10.0
	dapr init

install-dapr-generator:
	npm install -g yo
	npm install -g generator-dapr
