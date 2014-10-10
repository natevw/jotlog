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
  req, err := http.NewRequest("GET", "http://ipcalf.com/", nil)
  abortOnError(err)
  req.Header.Add("Accept", "text/plain")
  
  resp, err := http.DefaultClient.Do(req)
  abortOnError(err)
  defer resp.Body.Close()
  
  body, err := ioutil.ReadAll(resp.Body)
  abortOnError(err)
  fmt.Printf("%s\n", body)
}