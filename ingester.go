package main

//import "io"
import "os"
import "fmt"
import "log"
import "net/http"
import "io/ioutil"
import "encoding/json"

func abortOnError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err) 
    os.Exit(-1)
  }
}


// via https://groups.google.com/d/msg/golang-nuts/s7Xk1q0LSU0/vSvGnerlDZ4J
func Log(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
    handler.ServeHTTP(w, r)
  })
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
  abortOnError(err)
  fmt.Printf("%s\n", i.Vendor["version"])
  
  
  http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {
    err := req.ParseForm()
    abortOnError(err)
    
    singleValForm := map[string]string{}
    for key, values := range req.Form {
        singleValForm[key] = values[0]
    }
    
    s, err := json.Marshal(singleValForm)
    abortOnError(err)
    fmt.Printf("%s - %s\n%s\n\n", singleValForm["From"], singleValForm["Body"], s)
    
    fmt.Fprintf(w, "You said: %s\n", req.Form["Body"])
  })
	err = http.ListenAndServe(":8080", Log(http.DefaultServeMux))
  abortOnError(err)
}