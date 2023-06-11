# playground-dapr (work in progress)


```
zipkin: localhost:9411
redis commander: localhost:8081
dapr dashboard: localhost:9000
adminer: localhost:8080
    System: SQLite 3
    Username:
    Password: pass
    Database: /db/main.db
```

Plan:

```mermaid
flowchart TB
    inputfile["inputfile\nnew transactions"]
    outputfile["outputfile\ncopleted transactions"]
    debitTnx["debitTnx\ndebit source"]
    creditTnx["creditTnx\ncredit destination"]

    start --> inputfile
    inputfile --> gateway
    gateway -.->|lock amount and init transfer\ntopic.balance| balance
    linkStyle 2 stroke:blue;

    balance -.->|topic.debit_transaction| debitTnx
    debitTnx -.->|topic.balance| balance
    linkStyle 3 stroke:orange;
    linkStyle 4 stroke:orange;

    balance -.->|topic.credit_transaction| creditTnx
    creditTnx -.->|topic.balance| balance
    linkStyle 5 stroke:green;
    linkStyle 6 stroke:green;

    gateway --> outputfile
    outputfile --> X[end]
    linkStyle 7 stroke:red
    linkStyle 8 stroke:red
```


TODO:
- setup k0s cluster
- install redis, nats, postgres, vault?, grafana? with helm
- port dapr app to k8s
- scale to zero with keda.sh
- check how to debug with Bridge to k8s/dapr
- check grafana loki/tempo
- check debezium

