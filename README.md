## Receipt Processor

This is a REST API for processing receipts and calculating points based on receipt data written in Go, using the gin-gonic framework.

### Features
- Process receipts endpoint to calculate points. Calculates points based on the following:
  - One point for every alphanumeric character in the retailer name.
  - 50 points if the total is a round dollar amount with no cents.
  - 25 points if the total is a multiple of 0.25.
  - 5 points for every two items on the receipt.
  - If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
  - 6 points if the day in the purchase date is odd.
  - 10 points if the time of purchase is after 2:00pm and before 4:00pm.
- Get points endpoint to retrieve points for a receipt that is created
- Validation on receipts and items
- Unit tests for controllers, models, and services
- Dockerfile for containerization

### How to run

#### Docker
1. Clone the repository
2. Run `docker build -t receipt-processor .` to build the image
3. Run `docker run -p 8080:8080 receipt-processor` to run the container
4. The API will be available at `http://localhost:8080`

#### Local
1. Clone the repository
2. Run `go get` to install dependencies
3. Run `go run main.go` to run the API
4. The API will be available at `http://localhost:8080`

### Endpoints Available
- `POST /receipts/process` - Process a receipt and calculate points
- `GET /receipts/:id/points` - Get points for a receipt that was processed

### Example Requests
#### Process Receipt
```
curl --request POST \
  --url http://localhost:8080/receipts/process \
  --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODg4NTQ2NzIsImlkIjoxLCJpc3N1ZSI6MTY4ODI0OTg3MiwidXNlcm5hbWUiOiJleGFtcGxlIn0.oDYRkHfOrorDoYtomO0Z19cfcDyyp2NfrRJcLn7u8Hs' \
  --header 'content-type: application/json' \
  --data '{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
```

#### Get Points
```
curl --request GET \
  --url http://localhost:8080/receipts/:ID-FROM-ABOVE/points
```

### Testing
1. Run `go test ./tests` to run all tests

### Future Improvements
- Add a database to store receipts and points
- Add authentication and authorization
- Add more unit tests for full coverage