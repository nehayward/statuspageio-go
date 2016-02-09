# StatusPageio


## Overview
StatusPageio is a simple Go Client accessing StatusPage.io API.

## Getting Started
1. Simple `go get github.com/nehayward/statuspageio-go`

### Authentication

If you already have the API Key and PageID for your user, creating the client is simple:

````go
api := statuspageio.NewStatusioPageApi("apiKey", "pageID")
````

### Example

```go
api := statuspageio.NewStatusioPageApi("apiKey", "pageID")
api.ComponentUpdate("componentid", "operational")

```


## Requirements
* [Go](https://github.com/golang/example)

## Credits
- [Nick Hayward](https://github.com/nehayward)
