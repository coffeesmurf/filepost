# filepost

[![Build Status](https://travis-ci.org/coffeesmurf/filepost.png?branch=master)](https://travis-ci.org/coffeesmurf/filepost)

Golang application that monitors a folder for CSV files and uploads the content using a POST request.

## Usage

The following options can be specified: 

- path: The path to the folder being monitored (optional. Default: ./input)
- auth: Value used for the Authorization header of the POST request (mandatory).
- url: The URL used for the POST request (mandatory).

`$ filepost -auth="Bearer..." -path="~/Files" -url="https://...."`

The authorization header value can be read from a file (e.g. mytoken.txt) and passed in as an argument using: 

`$ filepost -auth=$(<mytoken.txt) ...`

## Notes 

- Only .csv and .json files are supported at the moment. The POST request uses a Content-Type of "text/csv" or "application/json".
- Files are removed from the folder once they have been processed, even if the upload is not successful.
- I am new at using Go. Feedback is welcome.