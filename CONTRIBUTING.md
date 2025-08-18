### Getting Started
To get the project up and running locally, follow these steps.

#### Prerequisites:
- [Go](https://go.dev/doc/install) (**v1.25**)
- [Git](https://git-scm.com/downloads)
- [ARCHITECTURE.md](ARCHITECTURE.md)
- make *(recommended for testing)*

#### 1. Clone the Repository:
```bash
git clone https://github.com/iton0/duss.git
cd duss
```
#### 2. Initialize Go Modules:
Use the provided init script to navigate into each service directory and initialize the Go modules.
```bash
./scripts/init.sh
```

### Testing
Tests can be run either by:
- changing into respective service directories and running  `go test ./... -v -coverage`
- using make
  - Please review the [Makefile](Makefile) for usage
