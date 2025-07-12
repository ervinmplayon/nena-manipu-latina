# The ID
## Why the Bidder Factory logs immediately at app startup
```shell
2025/07/02 22:56:43 [INFO] Bidder Strategy Factory: Selected FirstResponder
2025/07/02 22:56:43 [INFO] Bidder Strategy Factory: Selected RoundRobin
2025/07/02 22:56:43 [INFO] Starting ad server on :8080...
```
This is because the bidder factory logic us being executed during initialization, before any HTTP requests hit the server. The `publisherConfig` map on `config/publisher.go` executes the creation of the map on app startup. 

## To be Transfered to Spell Book
```go
package main
import(
    "errors"
    "fmt"
)

func doSomething() error {
    return errors.New("file not found")
}

func main() {
    err := doSomething()
    if err != nil {
        // Wrap the original error with additional context aka Adding a 
        // higher-level context to an underlying error without losing access 
        // to the original error's details. 

        wrappedErr := fmt.Errorf("failed to process data: %w", err)
        fmt.Println("Wrapped Error: ", wrappedErr)

        // Check if the original error is present in the wrapped error
        if errors.Is(wrappedErr, err) {
            fmt.Println()"Original error found in the wrapped error"
        }
    }
}
```