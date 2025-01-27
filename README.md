# Jeremiah Kwagh - Receipt Processor Challenge

Note: This is the submission by Jeremiah Kwagh of the Receipt Processor Challenge for Fetch.

This program was built using Go and provides 2 endpoints:
- A POST end point to process a receipt and generate a new unique ID
- A GET end point to retrieve the points total for a given receipt using its ID created by the POST end point

## Technologies Used
- Go
- Chi Router
- UUID
- CORS middleware

## API Endpoint

### Process Receipt
```
POST /v1/receipts/process
```
Processes a receipt in JSON format, stores in memory, and returns a newly created unique ID.

Example request body:
```json
{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}
```

### Get Points
```

GET /v1/receipts/{id}/points
