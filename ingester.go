package main

import "os"
import "fmt"
import "net/http"
import "io/ioutil"

func abortOnError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err) 
    os.Exit(-1)
  }
}

func main() {
  resp, err := http.Get("http://ipcalf.com/")
  abortOnError(err)
  defer resp.Body.Close()
  
  
  body, err := ioutil.ReadAll(resp.Body)
  abortOnError(err)
  fmt.Printf("%s\n", body)
}