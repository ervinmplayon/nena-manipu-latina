# The ID
## Why the Bidder Factory logs immediately at app startup
```shell\
2025/07/02 22:56:43 [INFO] Bidder Strategy Factory: Selected FirstResponder
2025/07/02 22:56:43 [INFO] Bidder Strategy Factory: Selected RoundRobin
2025/07/02 22:56:43 [INFO] Starting ad server on :8080...
```
This is because the bidder factory logic us being executed during initialization, before any HTTP requests hit the server. The `publisherConfig` map on `config/publisher.go` executes the creation of the map on app startup. 