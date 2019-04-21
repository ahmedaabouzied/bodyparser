# Golang Json Body Parser

A very tiny utility to parse JSON request body into a Map.

## Installation

    go get "github.com/ahmedaabouzied/bodyparser"

## Usage 

```go

package main 

import (
    "github.com/ahmedaabouzied/bodyparser"
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    body := bodyparser.JsonParse(r)
    fmt.Fprintf(w , "The Body had a key %s with a value : %s", "username", body["username"])
}

func main(){
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

```
Now if you send a request with a JSON body as follows

    {"username" : "My username"}

the server would respond with ``` My username ```

## Author

[Ahmed Abouzied](https://www.github.com/ahmedaabouzied)

Twitter [@ahmedaabouzied](https://twitter.com/ahmedaabouzied)

## Licence

MIT Licence , Do whatever the fuck you want with it.