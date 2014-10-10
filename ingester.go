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
  req, err := http.NewRequest("GET", "http://localhost:5984/", nil)
  abortOnError(err)
  req.Header.Add("Accept", "application/json")
  
  resp, err := http.DefaultClient.Do(req)
  abortOnError(err)
  defer resp.Body.Close()
  
  body, err := ioutil.ReadAll(resp.Body)
  abortOnError(err)
  fmt.Printf("%s\n", body)
  
  type info struct {
      Version string `json:"version"`
      UUID string `json:"uuid"`
      Vendor map[string]interface{} `json:"vendor"`
  }
  var i info
  err = json.Unmarshal(body, &i)
  fmt.Printf("%s\n", i.Vendor["version"])
}