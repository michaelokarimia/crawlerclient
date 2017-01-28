package crawlerclient

import ("io/ioutil"
        "net/http"
)


func Crawlerclient(domain string) (responseText string){
  resp, err := http.Get("http://" + domain)
  if err == nil{


  if b , err := ioutil.ReadAll(resp.Body); err == nil {
    responseText = string(b)
  }

  }  else {
    responseText = "Failed to parse"
  }
  return responseText
}
