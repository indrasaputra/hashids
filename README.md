# Hashids

[![Go Report Card](https://goreportcard.com/badge/github.com/indrasaputra/hashids)](https://goreportcard.com/report/github.com/indrasaputra/hashids)
[![Workflow](https://github.com/indrasaputra/hashids/workflows/Test/badge.svg)](https://github.com/indrasaputra/hashids/actions)
[![codecov](https://codecov.io/gh/indrasaputra/hashids/branch/main/graph/badge.svg)](https://codecov.io/gh/indrasaputra/hashids)
[![Maintainability](https://api.codeclimate.com/v1/badges/2cd8202174459c1b5348/maintainability)](https://codeclimate.com/github/indrasaputra/hashids/maintainability)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=indrasaputra_hashids&metric=alert_status)](https://sonarcloud.io/dashboard?id=indrasaputra_hashids)
[![Go Reference](https://pkg.go.dev/badge/github.com/indrasaputra/hashids.svg)](https://pkg.go.dev/github.com/indrasaputra/hashids)

Hashids is a package to convert ID into a random string to obfuscate the real ID from user.
In the implementation itself, the ID will still be an integer. But, when it is shown to the user,
it becomes a random string. The generated random string can be decoded back to the original ID.
This project uses [https://github.com/speps/go-hashids](https://github.com/speps/go-hashids) as the backend.

## Installation

```
go get github.com/indrasaputra/hashids
```

## Example

Let there be a struct:

```go
type Product struct {
    ID      hashids.ID  `json:"id"`
    Name    string      `json:"name"`
}
```

Then we have an instance of Product like this:

```go
product := &Product{
    ID: hashids.ID(66),
    Name: "Product's name",
}
```

When the `product` is marshalled into a JSON, the ID will not be a plain integer. It will become a random string like this:

```json
{
    "id": "kmzwa8awaa",
    "name": "product's name"
}
```

Upon decoding the ID, Hashids will decode back the random string to the original ID.

```go
var product Product
b := []byte(`{"id": "kmzwa8awaa","name": "product's name"}`)
json.Unmarshal(b, &product)
```

The code above will fill the product's attributes like this:

```cmd
{66 product's name}
```

## Limitation

For now, this package only support encoding to JSON, decoding from JSON, and encode as a string.