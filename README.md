# filepost

[![Build Status](https://travis-ci.org/coffeesmurf/filepost.png?branch=master)](https://travis-ci.org/coffeesmurf/filepost)

Golang application that monitors a folder for CSV files and uploads the content using a POST request.

## Usage

The following options can be specified: 

- path: The path to the folder being monitored (optional. Default: ./input)
- token: A Bearer token value used for the POST request (mandatory).
- url: The URL used for the POST request (mandatory).

`$ filepost -token="my token" -path="~/CSVFiles" -url="https://...."`

The token can be read from a file (e.g. mytoken.txt) and passed in as an argument using: 

`$ filepost -token=$(<mytoken.txt) ...`

## Notes 

- Only .csv files are supported at the moment. The POST request uses a Content-Type of "text/csv"
- Files are removed from the folder once they have been processed, even if the upload is not successful.
- I am new at using Go. Feedback is welcome.
