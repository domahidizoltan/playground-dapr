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
    classDef blue stroke:blue;

    inputfile["inputfile\nnew transactions"]
    outputfile["outputfile\ncopleted transactions"]
    debitTnx["debitTnx\ndebit source"]
    creditTnx["creditTnx\ncredit destination"]

    start --> inputfile
    inputfile --> gateway
    gateway -.->|lock amount and init transfer\ntopic.balance| balance
    balance -->|"PENDING[1]"| statestore
    linkStyle 2 stroke:blue,fill:blue;
    linkStyle 3 stroke:blue;

    balance -.->|topic.debit_transaction| debitTnx
    debitTnx -.->|topic.balance| balance
    linkStyle 4 stroke:orange;
    linkStyle 5 stroke:orange;

    balance -.->|topic.credit_transaction| creditTnx
    creditTnx -.->|topic.balance| balance
    balance -->|"COMPLETED[4]"| statestore
    linkStyle 6 stroke:green;
    linkStyle 7 stroke:green;
    linkStyle 8 stroke:green;

    statestore --> gateway
    gateway --> outputfile
    outputfile --> X[end]
    linkStyle 9 stroke:red
    linkStyle 10 stroke:red
    linkStyle 11 stroke:red
```
