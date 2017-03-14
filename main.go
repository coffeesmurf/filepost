package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {

	dataSourceTokenFlag := flag.String("token", "", "Bearer token used to post file content to the speficied URL.")
	folderPathFlag := flag.String("path", "./input", "Path to folder that will contain the .CSV files to process.")
	destinationURLFlag := flag.String("url", "", "Destination url where the POST request will be made.")

	flag.Parse()

	if *dataSourceTokenFlag == "" || *destinationURLFlag == "" {
		fmt.Println("You must specify a bearer token and a destination url.")
		fmt.Println("Use the -h option to learn more.")
		os.Exit(1)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	for {

		d, err := os.Open(*folderPathFlag)
		if err != nil {
			log.Fatal(err)
		}

		files, err := d.Readdir(-1)
		if err != nil {
			log.Fatal(err)
		}

		var csvFiles []string

		for _, fileInfo := range files {
			if strings.ToUpper(filepath.Ext(fileInfo.Name())) == ".CSV" {
				csvFiles = append(csvFiles, fileInfo.Name())
			}
		}

		if len(csvFiles) > 0 {

			fmt.Printf("Starting to upload %d files\n", len(csvFiles))

			startTime := time.Now()

			uploadFiles(*folderPathFlag, csvFiles, *dataSourceTokenFlag, *destinationURLFlag)

			fmt.Printf("%d files uploaded in %s\n", len(csvFiles), time.Since(startTime).String())
		}

		d.Close()
		time.Sleep(250 * time.Millisecond)
	}
}

func uploadFiles(folderPath string, files []string, dataSourceToken string, destinationURL string) {

	client := &http.Client{}

	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(len(files))

	var authorizationHeaderValue = "Bearer " + dataSourceToken

	for _, file := range files {
		go func(filePath string) {

			fmt.Printf("\tSarting %s\n", filePath)
			startTime := time.Now()

			defer os.Remove(filePath)
			defer waitGroup.Done()

			filesContent, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Println(err)
			}

			req, err := http.NewRequest("POST", destinationURL, bytes.NewBuffer(filesContent))
			if err != nil {
				fmt.Println(err)
			}

			req.Header.Add("Content-Type", "text/csv; chartset=utf-8")
			req.Header.Add("Authorization", authorizationHeaderValue)

			resp, err := client.Do(req)

			if err != nil {
				fmt.Println(err)
			}

			if resp.StatusCode != 200 {
				fmt.Println(string(resp.Status))
			}

			fmt.Printf("\tDone with %s in %s\n", filePath, time.Since(startTime).String())

		}(path.Join(folderPath, file))
	}

	waitGroup.Wait()
}
