package main

import (
  "log"
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
  "gopkg.in/gorp.v1"
  // "encoding/json"
)

const (
  hostname = "localhost"
  username = "postgres"
  // password = "12345678"
  dbname = "bayarkilat"
  sslmode = "disable"
)

func selectBusinessCategory(dbmap *gorp.DbMap) (*[]BusinessCategory, error) {
  // Run your query
  query := "select * from business_categories order by id"

  // pass a slice to Select()
  businessCategories := []BusinessCategory{}
  _, err := dbmap.Select(&businessCategories, query)
  checkErr(err, "Failed select BusinessCategory")
  return &businessCategories, err
}

func selectBusinessCategoryTypeView(dbmap *gorp.DbMap) (*[]BusinessCategoryTypeView, error) {
  // Run your query
  query := "select bc.id BusinessCategoryId, bt.id BusinessTypeId, bc.name BusinessCategoryName, bt.name BusinessTypeName " +
      "from business_categories bc, business_types bt " +
      "where bc.business_type_id = bt.id"

  // pass a slice to Select()
  var results []BusinessCategoryTypeView
  _, err := dbmap.Select(&results, query)
  checkErr(err, "Failed select BusinessCategoryTypeView")
  return &results, err
}

func selectBusinessOutletEmailListView(dbmap *gorp.DbMap) (*[]BusinessOutletEmailListView, error) {
  query := "select b.id BusinessId, b.name BusinessName, array_agg(distinct o.id) OutletIds  " +
      "from businesses as b " +
      "join outlets as o on b.id = o.business_id " +
      "join settings as s on b.id = s.business_id " +
      "join setting_alert_recipients as sar on s.id = sar.setting_id " +
      "where b.is_deleted = false and s.daily_sales_summary = true group by b.id order by b.id;"
  // pass a slice to Select()
  var results []BusinessOutletEmailListView
  _, err := dbmap.Select(&results, query)
  checkErr(err, "Failed select BusinessOutletEmailListView")
  return &results, err
}

func initDb() *gorp.DbMap {
  // connect to db using standard Go database/sql API
  // use whatever database/sql driver you wish
  dbinfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s",
          hostname, username, dbname, sslmode)
  db, err := sql.Open("postgres", dbinfo)
  checkErr(err, "sql.Open failed")

  // construct a gorp DbMap
  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

  return dbmap
}

func checkErr(err error, msg string) {
  if err != nil {
    log.Fatalln(msg, err)
  }
}
