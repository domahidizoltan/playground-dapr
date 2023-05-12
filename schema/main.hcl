schema "main" {}

table "balance" {
  schema = schema.main
  column "id" {
    null = false
    type = varchar(6)
  }
  column "balance" {
    null = false
    type = float
    default = 0
  }
  column "locked" {
    null = false
    type = float
    default = 0
  }
 primary_key {
    columns = [column.id]
  }
}

table "transaction" {
  schema = schema.main
  column "id" {
    null = false
    type = uuid
  }
  column "tnx" {
    null = false
    type = varchar(64)
  }
  column "source_acc" {
    null = false
    type = varchar(6)
  }
  column "dest_acc" {
    null = false
    type = varchar(6)
  }
  column "amount" {
    null = false
    type = float
  }
  column "datetime" {
    null = false
    type = varchar(32)
  }
 primary_key {
    columns = [column.id]
  }
}
