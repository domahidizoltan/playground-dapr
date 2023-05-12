install: install-dapr-cli update-dapr-components

start-app: stop-app delete-logs
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

generate:
	go run -mod=mod entgo.io/ent/cmd/ent generate --target ./balanceservice/ent/generated  ./balanceservice/ent/schema
	go run -mod=mod entgo.io/ent/cmd/ent generate --target ./transactionservice/ent/generated  ./transactionservice/ent/schema

BALANCE_DB_DRIVER=sqlite3
BALANCE_DB_DATASOURCE=file:localdev/sqlitedata/main.db
seed:
	sudo chmod 0666 localdev/sqlitedata/main.db
	BALANCE_DB_DRIVER=$(BALANCE_DB_DRIVER) BALANCE_DB_DATASOURCE=$(BALANCE_DB_DATASOURCE) go run balanceservice/seed/main.go


GATEWAY_ADDR=localhost:3000
test-inittransfer:
	curl -v http://$(GATEWAY_ADDR)/inittransfer?srcAcc=$(SRCACC)\&dstAcc=$(DSTACC)\&amount=$(AMT)

BALANCE_ADDR=localhost:3012
test-lock-balance: 
	curl -X POST http://$(BALANCE_ADDR)/balanceCommand \
		-d '{"datacontenttype": "application/json", "data": {"command": "lockBalance", "tnx":"TEST$(shell date +%s)","amount":$(AMT), "srcAcc":"$(SRC)", "dstAcc":"$(DST)"}, "topic": "balance", "pubsubname": "balance"}'

# update-dapr-components:
# 	cp -R components ~/.dapr

# PUB_TNX_ADDR=localhost:3001
# publish-debit: 
# 	curl -X POST http://$(PUB_TNX_ADDR)/debitsource \
# 		-d '{"datacontenttype": "application/json", "data": {"tnx":"$(TNX)","amount":$(AMT), "srcAcc":"$(SRC)", "dstAcc":"$(DST)"}, "topic": "transfer", "pubsubname": "debit-source"}'
# publish-credit: 
# 	curl -X POST http://$(PUB_TNX_ADDR)/creditdest \
# 		-d '{"datacontenttype": "application/json", "data": {"tnx":"$(TNX)","amount":$(AMT), "srcAcc":"$(SRC)", "dstAcc":"$(DST)"}, "topic": "transfer", "pubsubname": "credit-dest"}'
