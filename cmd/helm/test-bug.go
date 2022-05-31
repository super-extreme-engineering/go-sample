// main.go
package main

import {
    "encoding/json",
    "io/ioutil",
    "net/http",
    "os",
    "text/template"
}

func main() {
    tmpl := template.Must(template.ParseFiles("index.html"))
    name := "John"
   
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      data := PageData {
          name: name,
      }
       
      tmpl.Execute(w, data)
   
    })
     
    http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}