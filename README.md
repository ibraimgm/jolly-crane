# jolly-crane

## Compiling and running

You have two options:

    # just run directly
    go run main.go

    # to build and run
    go build ./... && ./jolly-crane

## Running with docker

    docker run --rm -p 8080:8080 ibraimgm/jolly-crane

## Available endpoints

- `GET /admin/check`: A simple return, just to check if the server is alive.
- `POST /hash`: Saves a new hash. Expect a Json body in the format `{"token":"some text"}`. Returns the saved hash.
- `GET /hashes/:hash`: Returns the token (and additional info) from a given hash.
- `GET /hashes`: Returns an array with all hashes stored.
- `GET /demo`: Inserts initial data to help demonstrate the API.
