# golang-api
An exercise to gain exposure with the Go language.

## Run Instructions:
Make sure you have Go Tools installed first. 
1. Clone this project into your $GOPATH src directory (i.e. "/User/youruser/go/src/")
2. Navigate the cloned project's directory.
3. Run the command "go build" to compile the application.
4. Execute the program generated, the default command being "./golang-api".
5. To execute the unit tests, run the command "go test -v ./...".

## API Endpoints
1. */hash* - Accepts an HTTP POST value from a form field called "password", base64 encode it, SHA512 hash it, and then return the resulting value.
2. */shutdown* - An HTTP GET here will tell the server to shutdown gracefully (i.e. the server will reject new requests and shutdown when all current requests are resolved).
3. */stats* - An HTTP GET here will return statistics the number of requests and the average process time of all HTTP GET requests made to the */hash* endpoint. The stats are stored in runtime memory, so they will only be retained for as long as the API is running.