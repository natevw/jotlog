package main

import "os"
import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"

func abortOnError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err) 
    os.Exit(-1)
  }
}

func main() {
  req, err := http.NewRequest("GET", "http://ipcalf.com/", nil)
  abortOnError(err)
  req.Header.Add("Accept", "application/json")
  
  resp, err := http.DefaultClient.Do(req)
  abortOnError(err)
  defer resp.Body.Close()
  
  body, err := ioutil.ReadAll(resp.Body)
  abortOnError(err)
  
  var s string
  err = json.Unmarshal(body, &s)
  fmt.Printf("%s\n", s)
}