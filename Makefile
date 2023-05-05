install: install-dapr-cli update-dapr-components

start-app:
	dapr run --run-file dapr.yaml &

stop-app:
	dapr stop --run-file dapr.yaml

start-dependencies:
	docker start dapr_redis dapr_zipkin dapr_placement
	docker-compose up -d
	dapr dashboard -p 9000

install-dapr-cli:
	wget -q https://raw.githubusercontent.com/dapr/cli/master/install/install.sh -O - | /bin/bash -s 1.10.0
	dapr init

delete-logs:
	rm -r .dapr/logs
	
# update-dapr-components:
# 	cp -R components ~/.dapr
