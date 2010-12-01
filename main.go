package main

import(
    "fmt"
    "flag"
    "io/ioutil"
    "dml"
)

func main() {
	html := flag.Bool("embed", false, "Embeds the converted DML in an HTML body before writing it out.")
	flag.Parse()

	file, err := ioutil.ReadFile(flag.Arg(0))
        if err != nil {
		fmt.Println("Could not" + err.String())
        }

	if *html {
		fmt.Println(dml.ToHTML(flag.Arg(0), string(file)))
	} else {
		fmt.Println(dml.ParseDoc(string(file)))
	}
}
