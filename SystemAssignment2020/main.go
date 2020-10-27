package main

import (
	"SystemAssignment/utils"
	"flag"
	"fmt"
	"net/url"
)

// DEBUG - Global debug flag
var DEBUG bool = false

// httpsLink - Const
var httpsLink string = ":443"

func main() {
	// Name for -agruments, taking the input
	url := flag.String("url", "", "URL argument to get the url to get from. (EX: https://general-sde-2021.leolemain.workers.dev/)")
	help := flag.Bool("help", false, "HELP argument to get the description information about the CLI")
	profile := flag.Int("profile", -1, "PROFILE argument to get the number of request time. (EX: 10)")

	flag.Parse()

	if *help {
		helperDescription()
		return
	}

	if *url == "" {
		errHelper("URL is not provided. Please enter the URL before moving forward. Use command: 'go run main.go -help' for more information.", nil)
		return
	}

	cli := utils.ClientManager{
		URL:   *url,
		Times: *profile,
	}

	var parsErr error
	cli.URL, cli.Path, cli.Port, parsErr = Parser(*url)
	if parsErr != nil {
		return
	}

	cli.IsHTTPS = (cli.Port == httpsLink)
	cli.MakeRequest()

}

// Parser - Parse all url coming in
func Parser(requestURL string) (string, string, string, error) {
	result, err := url.Parse(requestURL)
	if err != nil {
		errHelper("Error while parsing URL string", err)
		return "", "", "", err
	}

	if DEBUG {
		fmt.Println("Parse Scheme: ", result.Scheme)
		fmt.Println("Parse Host: ", result.Host)
		fmt.Println("Parse Path: ", result.Path)
	}

	// Port number depending on the URL Scheme
	portNumber := ":80"
	if result.Scheme == "https" {
		portNumber = ":443"
	}

	return result.Host, result.Path, portNumber, err
}

// helperDescription - Print out all -help flag
func helperDescription() {
	fmt.Print("Leo Le's Golang CLI Description \n\n")
	fmt.Print("The commands are:\n\n")
	fmt.Println("	-url		a URL string that is going to be requested to")
	fmt.Println("	-help 		a description of the CLI usage")
	fmt.Println("	-profile 	a positive integer which is the amount of request time to the URL provided")
	fmt.Println()
}

// errHelper - Print out error for functions
func errHelper(message string, err error) {
	fmt.Println(message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}
