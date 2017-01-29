package crawlerclient

import ("io/ioutil"
        "net/http"
        "github.com/michaelokarimia/htmlparser"
)


func GetUrls(domain string) (responseText string){
  resp, err := http.Get("http://" + domain)
  if err == nil{


  if b , err := ioutil.ReadAll(resp.Body); err == nil {
    responseText = string(b)
  }

  }  else {
    responseText = "Failed to parse"
    return
  }

  var urls= htmlparser.Htmlparser(responseText, domain)

  var urllist = "urls found:"
  for i:= range urls {
    urllist += "\n" + urls[i]
  }

  return urllist
}
