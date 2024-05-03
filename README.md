# Wait For Response Docker action

This action makes HTTP requests to a given URL until the required response code is retrieved or the timeout is met.  Initially created to allow test containers to startup before executing tests against them.

## Inputs

### `url`

The URL to poll. Default "http://localhost/"

### `server`

The hostname or IP address to connect to. When specified, the action connects to its value and uses the hostname part of the url as the HTTP host header when performing the request.

### `useDefaultGateway`
Detect and the default gateway as the server to connect to. Useful in containerised environments. Takes precedence over host. Default: false

### `method`

The HTTP method to use. Default: HEAD

### `expectedCode`

Response code to wait for. Default: 200

### `expectedBody`

Check the response, if it is what we expected. Default ""

### `timeout`

Timeout before giving up in milliseconds. Default 30000

### `interval`

Interval between polls in milliseconds. Default: 200

## Example usage
```
uses: omame/wait_for_response_body@master
with:
  url: 'http://localhost:8081/'
  useDefaultGateway: true
  responseCode: 200
  timeout: 2000
  expectedResponse: "latest"
  interval: 500
```
