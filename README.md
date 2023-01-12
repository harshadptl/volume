# Volume Assignment

## Problem Statement
https://docs.google.com/document/d/1NbxiClee0xhqxPmfyHGkxe8JkZNQ2zg2EgEySyJK4jY/edit

## Build

### Requirements

- Go 1.19

### Command
    $ make build

### Test
    $ make test ./...

## Run
    $ bin/server

## Usage

```bash

$ curl -v localhost:8080/v1/flights/sort -d '[["YYZ","BOM"],["DEL","YYZ"]]' 
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /v1/flights/sort HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.79.1
> Accept: */*
> Content-Length: 29
> Content-Type: application/x-www-form-urlencoded
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Thu, 12 Jan 2023 01:17:48 GMT
< Content-Length: 51
< Content-Type: text/plain; charset=utf-8
< 
{"success":true,"src":"DEL","dest":"BOM","err":""}

```