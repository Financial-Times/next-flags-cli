package main

import (
  "os"
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/franela/goreq"
  "encoding/json"
)

type Response1 struct {
    Page   int
    Fruits []string
}

func main() {
  app := cli.NewApp()
  app.Name = "greet"
  app.Usage = "fight the loneliness!"
  app.Action = func(c *cli.Context) {
    
	res, _ := goreq.Request{ Uri: "https://next-flags.ft.com/api/v1" }.Do()
	//fmt.Println(res.Response, err)

	//byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
    str, _ := res.Body.ToString()
	if err := json.Unmarshal([]byte(str), &dat); err != nil {
		panic(err)
	}    
	//fmt.Println(dat)
	version := dat["version"].(float64)
	fmt.Println(fmt.Sprintf("version: %3.f", version))

  }

  app.Run(os.Args)
}
