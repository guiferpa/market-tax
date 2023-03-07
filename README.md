# market tax

- [Get started](#get-started)
  - [Build source code](#build-source-code)
  - [Executing binary](#executing-binary)
  - [Containerizing binary](#containerizing-binary)
  - [Executing container with binary](#executing-container-with-binary)
  
- [Tasks](#tasks)
  - [Running lint](#running-lint)
  - [Running unit tests](#running-only-unit-tests)
  - [Running integration tests](#running-all-tests-including-integration-tests)

- [Source code design pattern](#source-code-design-pattern)
  - [Tree overview](#tree-overview)
  - [Concepts of source code arch](#concepts-of-source-code-arch)

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

## Tasks

> :balloon: This project has `Makefile` as job runner

### Running lint
```sh
$ make lint
```

### Running only unit tests
```sh
$ make test
```

### Running all tests (including integration tests)
```sh
$ make integration

```

### Source code design pattern

#### Tree overview

```sh
.
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   └── cli
│       ├── main.go
│       └── main_test.go
├── domain
│   └── stock
│       ├── port.go
│       ├── usecase.go
│       └── usecase_test.go
├── go.mod
├── handler
│   └── interface
│       └── cli
│           ├── interface.go
│           └── interface_test.go
├── infra
│   └── storage
│       └── memory
│           └── storage.go
└── pkg
    └── storage
        └── storage.go
```

#### Concepts of source code arch

- **cmd**: This directory's responsible for app's entrypoint. In this case we have a integration for CLI but, for example, it could there is REST API 
integration too and both existing in the same source code.

- **handler**: This directory's responsible for app's user interface protocol. It's here that'll develop all rules for handle an input then pass to use cases/domain layer.

- **domain**: This directory's responsible for app's core, where it'll be the business rule. Given hexagonal arch, it's here that all ports communicate with use cases rules and never the inverse.

- **infra**: This directory's responsible for all app's external integration which it'll helpful to use cases process the app's input. In this case it was created a short data structure that'll localated in [pkg/storage](https://github.com/guiferpa/market-tax/tree/main/pkg/storage), this data structure's used for infra layer to instance a adapter and help use cases processing.

- **pkg**: This directory's responsible for all app's modules that hasn't fit with hexagonal components. In this case we have a module called [pkg/storage](https://github.com/guiferpa/market-tax/tree/main/pkg/storage). In this case it serves to manage datas from buy/sell stock contexts.
