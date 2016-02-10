package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app      = kingpin.New("Options Parser", "Thanks to https://github.com/alecthomas/kingpin")
	aStringFlag = kingpin.Flag("strF", "shows a string Flag").String()
	aStringValue    = kingpin.Arg("strV", "shows a string Value").String()
	numFlag   = kingpin.Flag("numF", "shows a Number Flag" ).Int()
	numValue   = kingpin.Arg("numV", "shows a Number Value" ).Int()

	upload      = app.Command("upload", "Upload a file to S3")
	download      = app.Command("download", "Download a file to S3")

)

func main() {

	kingpin.Version("0.0.1")
	kingpin.Parse()
	fmt.Printf("Would run app with Flags: %s  %d and values  %s  %d \n",
										*aStringFlag, *numFlag, *aStringValue, *numValue)


	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Register user
	case upload.FullCommand():

		println("\nupload code goes here\n")

	// Post message
	case download.FullCommand():

		println("\nupload code goes here\n")
	}


}
func check(err error) {
	if err != nil {

		return
	}

}