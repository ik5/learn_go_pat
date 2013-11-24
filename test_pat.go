package main

import (
  "io"
  "github.com/gorilla/pat"
  "net/http"
  "log"
)

const Port = "3000"

func tostr(req *http.Request) string {
  log.Println("tostr")
  if v := req.URL.Query().Get(":name"); v != "" {
    return v
  }

  return "[empty]"
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    log.Println("HelloServer")
    io.WriteString(w, "hello, " + tostr(req) + "!\n")
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
