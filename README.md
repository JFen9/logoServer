### Build docker image
In the project directory run

`docker build . -t logo`

### Run logo server
`docker run -p 8124:8124 logo`

The server will be listening on port 8124

### How to run unit tests
`go test ./service`
