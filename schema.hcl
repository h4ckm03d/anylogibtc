table "transactions" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "datetime" {
    null = false
    type = timestamptz
  }
  column "amount" {
    null = false
    type = numeric(6,2)
  }
  primary_key {
    columns = [column.id, column.datetime]
  }
  index "transactions_datetime_idx" {
    on {
      desc   = true
      column = column.datetime
    }
  }
}
schema "public" {
}
