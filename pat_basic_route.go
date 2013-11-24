package main

import (
  "io"
  "github.com/gorilla/pat"
  "net/http"
  "log"
)

const Port = "3000"

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    log.Println("HelloServer")
    tostr := func() string {
      log.Println("tostr")
      if v := req.URL.Query().Get(":name"); v != "" { return v }
      return "[empty]"
    }
    io.WriteString(w, "hello, " + tostr() + "!\n")
}

func main() {
    m := pat.New()
    m.HandleFunc("/hello/{name}", HelloServer)

    http.Handle("/", m)
    log.Printf("Going to listen on port %s.\n", Port)
    err := http.ListenAndServe(":" + Port, nil)
    if err != nil { 
      log.Fatal("ListenAndServe: ", err)
    }
}
