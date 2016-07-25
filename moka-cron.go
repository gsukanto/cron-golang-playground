package main

import (
  "encoding/json"
  "fmt"
)

func main() {
  // initialize the DbMap
  dbmap := initDb()
  defer dbmap.Db.Close()

  fmt.Println("soemthing")

  // var list *[]BusinessCategoryTypeView
  // list, _ := selectBusinessCategoryTypeView(dbmap)
  // // fmt.Println(list)
  // for _, e := range *list {
  //   b, _ := json.MarshalIndent(e, "", " ")
  //   println(string(b))
  // }

  var list *[]BusinessOutletEmailListView
  list, _ = selectBusinessOutletEmailListView(dbmap)
  // fmt.Println(list)
  for _, e := range *list {
    b, _ := json.MarshalIndent(e, "", " ")
    println(string(b))
  }


}
