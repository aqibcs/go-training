# CSV to JSON Converter

This Go program converts data from a CSV (Comma-Separated Values) file to JSON format.

## Prerequisite

- Golang

## Build

To build this project

```bash
./scripts/build.sh
```

## Run

You can run the project using the following command:

```bash
./scripts/run.sh
```

## Build and Run

To build and run the project use the following command:

```bash
./scripts/run.sh -b
```


# Hello Handler
The `HelloHandler` is HTTP request handler designed to handle requests to the `/hello` endpoint in our Go web application. it greets the user based on the name provided in the JSON request body and responds with a personalized welcome message in JSON fromat.

## API Endpoint
* **Endpoint: `/hello`**
* **Http Method: POST**
* **Request Format: JSON**

```JSON
{
    "name": "YourName"
}
```
* **Response Format: JSON** 
```JSON
{
    "code": 200,
    "message": "Welcome YourName!",
    "timestamp": "2023-10-23T12:00:00Z"
}
```
