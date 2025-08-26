### Getting Started
To get the project up and running locally, follow these steps.

#### Prerequisites:
- [Go](https://go.dev/doc/install) (**v1.25**)
- [Git](https://git-scm.com/downloads)
- [ARCHITECTURE.md](ARCHITECTURE.md)
- make *(recommended for testing)*
- Fork and clone the repo

#### Initialize Go Modules:
Use the provided init script to navigate into each service directory and initialize the Go modules.
```bash
./scripts/init.sh
```

### Testing
Tests can be run either by:
- changing into respective service directories and running  `go test ./... -v -coverage`
- using make (this can be run at root level of the project)
  - Please review the [Makefile](Makefile) for usage
