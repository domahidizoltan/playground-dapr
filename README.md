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
<details>
<summary>Plant UML</summary>

```puml
@startuml
inputfile : new transactions
outputfile : copleted transactions
debitTnx : debit source
creditTnx : credit destination

[*] --> inputfile
inputfile -left> gateway
gateway -[#blue]-> balance : lock amount and init transfer
balance -down[#blue]-> statestore : PENDING[1]
balance -left[dashed,#orange]> debitTnx : topic.debit_source

debitTnx -[dashed,#orange]> balance : topic.update_source_balance

balance -right[dashed,#green]> creditTnx : topic.credit_destination
creditTnx -[dashed,#green]> balance : topic.update_destination_balance
balance -[#green]> statestore : COMPLETED[4]

statestore -[#red]> gateway
gateway -[#red]-> outputfile
outputfile -[#red]-> [*]
@enduml
```
</details>

![plan](http://www.plantuml.com/plantuml/svg/TL9DQm8n4BtdLmIybQPGwAa74QgKGcizU5iMYScuXiPaoKwm_VUTrKrd5-d98U_3DszdqQ5Ec4zUkD1cF3WFyba6E4jCEdJQe8kX4p4ZeoQs7X3ib69Xxt0Rlebm6MKNSp8WJ09RWEjCU8Skw5udH7LNIwNcyk__HqcKXmFEPQCHplf73BzILREzpr2JQg-z3WR8sqVp9VKfve1I1qj-3gy93v14uIaRpu7bj3rIc9XwXyrglNnRlrQFDTFx09NLtH7i_IoIMmFrN8vsnTWwyt1vs0qRSnNgLbgSYpAtLYCCqjs02WwGN7Fa14q22EJ2fHQVwyjkN2sJrDJWtnXZUSd2KQgYdjRsylcnULzjLylggHQ2eLwGmDQttsy0g7--Ay7Z2AVZ40i8bWxq5m00)