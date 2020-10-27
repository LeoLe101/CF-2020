# Requirement:
- Golang: Version 1.15.3 or >1.11.0

# Build The Project
1. To build the project, run the following command in the folder System Assignment 2020: `GO111MODULE=on go build main.go`
2. The project will create a `main` file. To run the project, follow the following commands:
    * `./main -help` for commands description 
        * Example: `./main -help`
    * `./main -url=<URL>` to request the HTML from this URL  
        * Example: `./main -url=https://general-sde-2021.leolemain.workers.dev/`
    * `./main -url=<URL> -profile=<Number>` to request the HTML with the amount specified in `-profile` flag from the same URL
        * Example: `./main -url=https://general-sde-2021.leolemain.workers.dev/ -profile=5`
3. If the build version does not work, use the following commands instead:
    * `go run main.go -help` for commands description 
        * Example: `go run main.go -help`
    * `go run main.go -url=<URL>` to request the HTML from this URL  
        * Example: `go run main.go -url=https://general-sde-2021.leolemain.workers.dev/`
    * `go run main.go -url=<URL> -profile=<Number>` to request the HTML with the amount specified in `-profile` flag from the same URL
        * Example: `go run main.go -url=https://general-sde-2021.leolemain.workers.dev/ -profile=5`

# Images 
* Project on Google website: Please check the Result-Images folder. Please check the images with "Google" naming on it
* Project on CloudFlare Worker website: Please check the Result-Images folder. Please check the images with "CloudFlare" naming on it
