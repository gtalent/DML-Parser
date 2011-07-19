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
	inpath := ""
	outpath := ""
	if flag.NArg() < 2 {
		fmt.Println("Insufficient arguments.")
		return
	}

	inpath = flag.Arg(0)
	outpath = flag.Arg(1)


	file, err := ioutil.ReadFile(inpath)
        if err != nil {
		fmt.Println("Could not" + err.String())
        }
	output := ""
	if *html {
		output = dml.ToHTML(flag.Arg(0), string(file))
	} else {
		output = dml.ParseDoc(string(file))
	}

	ioutil.WriteFile(outpath, []byte(output), 0)
}
