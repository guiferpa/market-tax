# market tax

- [Get started](#get-started)
  - [Build source code](#build-source-code)
  - [Executing binary](#executing-binary)
  - [Containerizing binary](#containerizing-binary)
  - [Executing container with binary](#executing-container-with-binary)

## Get started

### Build source code
```sh
$ CGO_ENABLED=0 go build -v -o ./dist/cli ./cmd/cli/main.go
```

### Executing binary

```sh
$ ./dist/cli << EOF
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}, {"operation":"sell", "unit-cost":20.00, "quantity": 5000}]
[{"operation":"buy", "unit-cost":20.00, "quantity": 10000}, {"operation":"sell", "unit-cost":10.00, "quantity": 5000}]
EOF
```

### Containerizing binary

> :balloon: It's necessary has [docker](https://www.docker.com/get-started/) installed

```sh
$ docker build -t market-tax .
```

### Executing container with binary

```sh
$ docker run -it market-tax
[{"operation":"buy", "unit-cost":10.00, "quantity": 100}, {"operation":"sell", "unit-cost":15.00, "quantity": 50}, {"operation":"sell", "unit-cost":15.00, "quantity": 50}]
[{"operation":"buy", "unit-cost":20.00, "quantity": 10000}, {"operation":"sell", "unit-cost":10.00, "quantity": 5000}]

```
