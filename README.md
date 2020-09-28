# load_distributor

### To run load distributor
    go mod download
    go build .
    ./load_distributor

### To add a request (repeat to add more)
    curl localhost:8080/add_request

### To get stats
    curl localhost:8080/get_stats

### To run tests
    go test .