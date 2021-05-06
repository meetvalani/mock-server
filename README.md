# Mock Server

## Get custom generated responses, status codes, headers, and content-type.

## Use:
To run pre-built server directly with docker.

    docker pull meetvalani/mock-server
    docker run -d -p 8080:80 meetvalani/mock-server
### To run:
    go mod download
    go build main.go
    ./main --host <ip>:<port>
### Ex:
    ex: ./main --host 0.0.0.0:8080

## APIs:
### Get all saved mocks.
    API: /get
    METHOD: GET

    Sample Response:
        [
            {
                "id": 1,
                "method": "get",
                "endpoint": "/test",
                "responseCode": 200,
                "httpResponseContentType": "application/json",
                "httpHeaders": {
                    "header1": "value1",
                    "header2": "value2"
                },
                "httpResponseBody": "{
                    "body": "data"
                }"
            }
        ]
### Get saved mocks by ID.
    API: /get/{id}
    METHOD: GET

    Sample Response:
        {
            "id": 1,
            "method": "get",
            "endpoint": "/test",
            "responseCode": 200,
            "httpResponseContentType": "application/json",
            "httpHeaders": {
                "header1": "value1",
                "header2": "value2"
            },
            "httpResponseBody": "{
                "body": "data"
            }"
        }
### Add new mock.
    API: /add
    Method: POST
    Request Body(JSON):
        {
            "id": 1,
            "method": "get",
            "endpoint": "/test",
            "responseCode": 200,
            "httpResponseContentType": "application/json",
            "httpHeaders": {
                "header1": "value1",
                "header2": "value2"
            },
            "httpResponseBody": "{}"
        }

    Sample Response:
        {
            "id": 1,
            "method": "get",
            "endpoint": "/test",
            "responseCode": 200,
            "httpResponseContentType": "application/json",
            "httpHeaders": {
                "header1": "value1",
                "header2": "value2"
            },
            "httpResponseBody": ""
        }
### Delete saved mock.
    API: /delete/{id}
    Method: DELETE

    Sample Response:
        {
            "message": "deleted successfully if present."
        }
### Update saved mock.
    API: /update
    Method: PUT
    Request Body(JSON):
        {
            "id": 1,
            "method": "get",
            "endpoint": "/test",
            "responseCode": 200,
            "httpResponseContentType": "application/json",
            "httpHeaders": {
                "header1": "value1",
                "header2": "value2"
            },
            "httpResponseBody": "{}"
        }
    
    Sample Response:
        {
            "id": 1,
            "method": "get",
            "endpoint": "/test",
            "responseCode": 200,
            "httpResponseContentType": "application/json",
            "httpHeaders": {
                "header1": "value1",
                "header2": "value2"
            },
            "httpResponseBody": "{
                "body": "data"
            }"
        }

## Using saved mock:
Call any saved mock api with correct method and it will return you the exact response body, response code, content-type, headers.

    ex:
        curl --location --request GET 'http://<ip>:<port>/test'
        curl --location --request GET 'http://<ip>:<port>/ab'
        curl --location --request GET 'http://<ip>:<port>/<any saved endpoint path>'