package main

import(
    "fmt"
    "dml"
    "flag"
)

func main() {
	html := flag.Bool("embed", false, "Embeds the converted DML in an HTML body before writing it out.")
	flag.Parse()
	if *html {
		fmt.Println(dml.ToHTML(flag.Arg(0)))
	} else {
		fmt.Println(dml.ParseDoc(flag.Arg(0)))
	}
}
