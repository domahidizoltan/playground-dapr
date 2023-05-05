schema "main" {}

table "balance" {
  schema = schema.main
  column "id" {
    null = false
    type = varchar(6)
  }
  column "balance" {
    null = true
    type = float
  }
 primary_key {
    columns = [column.id]
  }
}