# Insurance-api

The Insurance-api transform the user profile in a risk profile and evaluate the insurance rules such as Auto, Home, Life and Disability.

The API is using Golang 1.17v

## Dependencies

- Gorilla Mux v1.8.0
- Validator v9.31.0

# Technical decisions

- Isolate the web framework.
- Single-Responsibility principle for creating new structs.
- Dependency Inversion principle, the high-level module must not depend on the low-level module, but they should depend on abstractions.
- Standardizing API errors.

# Getting started

## Locally

### Requisites
- Golang >=1.17

To run the API on port 8080 (by default)
```shell
make run-local
```

To run all tests
```shell
make run-tests
```

## Docker

To run the API on port 8080
```shell
make docker-up
```

to run **tests on Docker**:

```shell
make docker-tests
```

# How to use

The API documentation is in ./docs/openapi

## `POST /evaluation`

The body request to **evaluate a risk profile** is:

```shell
{
  "age": 35,
  "dependents": 2,
  "house": {"ownership_status": "owned"},
  "income": 2000000,
  "marital_status": "married",
  "risk_questions": [1, 1, 0],
  "vehicle": {"year": 2018}
}
```

Required fields:
- Age (an integer equal or greater than 0).
- Dependents (an integer equal or greater than 0).
- Income (an integer equal or greater than 0).
- MartialStatus ("single" or "married").
- RiskQuestions (an array with 3 booleans).

Users can have 0 or 1 house. When they do, it has just one attribute: ownership_status, which can be "owned" or "mortgaged".

Users can have 0 or 1 vehicle. When they do, it has just one attribute: a positive integer corresponding to the year it was manufactured.

### The output

The response could be an error HTTP 400 with the error messages:
```shell
{
    "message": [
        "validation error on RiskProfileRequest.House.OwnershipStatus",
        "validation error on RiskProfileRequest.Vehicle.Year"
    ]
}
```

And could be an HTTP 200 OK with the insurance suggestion plans:
```shell
{
    "auto": "regular",
    "disability": "economic",
    "home": "economic",
    "life": "regular"
}
```