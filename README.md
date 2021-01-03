# Gibberish Microservice

REST-powered microservice for analyzing text and returning if it is likely gibberish or not.
The repository also contains a client to perform requests to the service.

It is important to notice that this service provides no support for authentication or caching. It is also completely stateless, making it ideal to be used in the backend. A possible "frontend" implementation can be found in [Analysis API](https://gitlab.com/shitposting/analysis-api).

## Endpoints

- Gibberish endpoint: `<bind-address>/gibberish`
- Health check: `<bind-address>/healthy`

## Returned data

The data returned by the server is in the form:

```go
type GibberishResponse struct {
    IsGibberish bool
}
```

## Environment options

- Service bind address and port: `"GB_BIND_ADDRESS"` (defaults to `localhost:10002`).
- Path to knowledge base file: `GB_KNOWLEDGE_PATH` (defaults to `knowledge.json`).
