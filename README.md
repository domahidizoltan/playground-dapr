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
inputfile -> gateway
gateway -down[#blue]-> statestore : PENDING[1]
gateway -up[#blue]-> balance : reserve pending amounts
balance -[#blue]> gateway : get account balance

gateway -[#orange]> debitTnx
debitTnx -[#orange]> balance : update source

balance -[#green]> creditTnx
creditTnx -[#green]> balance : update destination
balance -[#green]> statestore : COMPLETED[4]

statestore -[#red]> gateway
gateway -[#red]-> outputfile
outputfile -down[#red]-> [*]
@enduml
```
</details>

![plan](http://www.plantuml.com/plantuml/svg/PP8nQ_im4CNt-nI2__O73uLE7KgW2QNGfXco6GyNwMe4oib8fwRzzfNQo8hnb2Flz-uTJzv4mI3fxC3obEJ3Eb8FYkcY9217r68zH_19cghzv4Z8B1539oj7_ih0xwYYJq4Jw42c2dzprgDOnk83wFyFgkiUrPn_Sqd-UqIX2tx3zLTrnb-u_tToYOOiHq6XA3wKmmwx_VPb_zpV3GrFKuDFw91r8GD52f-a9c9ZULHuzeabGYMwgsdEpvwHHA7M1QoReEOKWm_8Ox7K9g0E2xKTaIQ3GhMds-mn475cv-vQYqrUhreKkrwtBElf_UtmkZlkjltz0D2aOiXXmwhwAkGYs-S0RRm-JMNokyA6sAIp-m40)