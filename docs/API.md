# Pykitzoid API

The API exposes the existing linear regression calculations without creating
plots or accessing local CSV paths. Each request fits a model from the supplied
samples, returns its coefficients and R-squared, and optionally predicts new y
values. Requests and models are not persisted.

## Run locally

Install Go 1.21.3 or newer, then from the repository root:

```bash
cd "algorithms/Linear Regression"
go run ./cmd/api
```

The server listens at `http://localhost:8080`. Set `PORT` to change the port.
Run `go test ./...` in the same directory to run the existing algorithm tests and
the API tests. The existing Go CI workflow also checks the new packages.

## Postman

Import `docs/pykitzoid.postman_collection.json` into Postman. Its `baseUrl`
variable defaults to `http://localhost:8080`; replace it with a hosted URL after
deployment. The collection includes health, discovery, prediction, and invalid
input requests with response assertions. `docs/openapi.json` is also importable
as an OpenAPI 3.0 specification.

| Method | Path | Purpose |
| --- | --- | --- |
| GET | `/health` | Health check, returns `{"status":"ok"}` |
| GET | `/api/v1/algorithms` | List available algorithms and endpoints |
| POST | `/api/v1/linear-regression` | Fit a line and predict values |

```bash
curl -X POST http://localhost:8080/api/v1/linear-regression \
  -H "Content-Type: application/json" \
  -d '{"samples":[{"x":1,"y":3},{"x":2,"y":5},{"x":3,"y":7}],"predict":[4,5]}'
```

Expected response:

```json
{
  "algorithm": "linear-regression",
  "slope": 2,
  "intercept": 1,
  "r_squared": 1,
  "sample_count": 3,
  "predictions": [9, 11]
}
```

`predict` is optional; omitting it returns an empty predictions array. The API
uses the repository's `float32` calculations, so allow for rounding. Supply
2–10,000 samples with at least two distinct x values, and at most 10,000
prediction inputs. Every sample must contain numeric `x` and `y` fields.
Sample magnitudes are capped at `1e15` to protect the algorithm's accumulations;
very small scales may also need rescaling due to float32 precision. R-squared
is `null` when all observed y values are identical because the score is undefined.

Errors have the JSON shape `{"error":"Explanation"}`:

- `400`: invalid JSON, unknown fields, invalid sample count, missing/null values,
  constant x values, or numbers outside the float32 range.
- `404`: unknown endpoint.
- `405`: unsupported method, with an `Allow` header.
- `413`: body larger than 1 MiB.
- `415`: content type is not `application/json`.
- `422`: numeric scale cannot be calculated reliably; rescale the data.

## Deploy later

Deployment files are prepared; no public service has been created.
From the repository root, with Docker installed:

```bash
docker build -t pykitzoid-api .
docker run --rm -p 8080:8080 pykitzoid-api
```

The image contains a standalone Go binary running as a non-root user. It listens
on all interfaces and respects the hosting provider's `PORT` variable.

For a provider such as Render, create a Docker web service for your repository,
use the root `Dockerfile`, and set its health-check path to `/health`. After the
deployment succeeds, put the service's HTTPS URL into Postman's `baseUrl`.
See [Render's Docker guide](https://render.com/docs/docker) and
[port configuration](https://render.com/docs/web-services#port-binding).

This is a stateless demonstration API without authentication. Add gateway
authentication and rate limiting when extending it into a production service.

## Adding algorithms

1. Place reusable calculations in an importable Go package. The extracted
   `regression` package is shared by the original CLI demo and this API.
2. Add a handler with request validation in `api`, then register its versioned
   route and add its entry to the catalog in `api.NewHandler`.
3. Add numerical and HTTP tests, and update the OpenAPI spec and Postman collection.

Polynomial regression is not advertised as available: its existing implementation
has an empty weight-fitting function and syntax errors. Complete and test that
algorithm before exposing it through a new endpoint.
