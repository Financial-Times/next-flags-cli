package main

import (
  "os"
  "fmt"
  "sort"
  "github.com/codegangsta/cli"
  "github.com/franela/goreq"
  "encoding/json"
)

type Flags struct {
	Version		float64
	Flags		[]Flag
}

type Flag struct {
    Name		string
	Description	string
    State 		bool
}

const CLR_G = "\x1b[32;1m"
const CLR_R = "\x1b[31;1m"
const CLR_N = "\x1b[0m"

// Sorting data in Go requires you to implement the sort.Interface. This
// interface requires three simple methods: Len, Less, and Swap.

type ByLength []Flag

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
    return s[i].Name < s[j].Name
}

func main() {

  app := cli.NewApp()
  app.Name = "next-flags"
  app.Action = func(c *cli.Context) {
    
	res, _ := goreq.Request{ Uri: "https://next-flags.ft.com/api/v1" }.Do()

    str, _ := res.Body.ToString()
	flags := Flags{}
	if err := json.Unmarshal([]byte(str), &flags); err != nil {
		panic(err)
	}    

	// output
	fmt.Println(fmt.Sprintf("version: %3.f", flags.Version))
	fmt.Println(fmt.Sprintf("number of flags: %d", len(flags.Flags)))
	fmt.Println(" ")
	
	sort.Sort(ByLength(flags.Flags))

	for i := 0; i < len(flags.Flags); i++ {
		if (flags.Flags[i].State == true) {
			fmt.Println(fmt.Sprintf(" %s%t%s	%s", CLR_G, flags.Flags[i].State, CLR_N, flags.Flags[i].Name))
	 	} else {
			fmt.Println(fmt.Sprintf(" %s%t%s	%s", CLR_R, flags.Flags[i].State, CLR_N, flags.Flags[i].Name))
		}
	}

  }

  app.Run(os.Args)
}
