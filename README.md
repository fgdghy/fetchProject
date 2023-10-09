# fetch-challenge

## Running the application
1. Run `go mod download` to download the dependencies required

3. Build the application by running `go build` and run the executable or simply just run `go run main.go`


## API Endpoints:
The application provides the following API endpoints:

`POST /receipts/process`: Endpoint for processing receipts. It expects a JSON payload representing a receipt and returns an ID for the receipt.

`GET /receipts/{id}/points`: Endpoint for retrieving the number of points awarded for a given receipt ID.


## Testing the API:
You can test the API using tools like `curl` or `Postman`. Refer to the API specification in the challenge for details on how to use the endpoints.