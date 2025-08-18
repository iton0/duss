# duss

> [!CAUTION]
> This project is being actively developed
> and is subject to breaking changes

### Distributed URL Shortening Service

duss is a modern, distributed URL shortening service built with Go. This project is a hands-on exploration of a microservices architecture, focusing on performance, scalability, and maintainability.

---

### Features

- **URL Shortening**: Converts a long URL into a short, manageable link.
- **URL Redirection**: Redirects users from the short link to the original destination.
- **Microservices Architecture**: The system is composed of multiple independent services.

### Tech Stack

- **Backend**: Go
- **Web Framework**: Gin
- **Database**: PostgreSQL (for durable storage)
- **Caching**: Redis (for high-performance read lookups)
- **Containerization**: Docker & Docker Compose
- **Cloud Platform**: Railway (for deployment)

### Contributing

We welcome your contributions! To keep our issue tracker organized and effective, we've outlined a simple process:

#### Before Opening an Issue

Before you open a new issue, please start a conversation in our **[GitHub Discussions](https://github.com/iton0/duss/discussions/new/choose)** section. This is the best place to share:

* Feature requests
* Bug reports
* General questions about configuration or behavior

#### Opening an Issue

Once your topic has been discussed and approved in the discussions section, you can then open a formal issue.

### Future Improvements

- [ ] Add UI for checking whether to trust a long URL; this would only run once
  for untrusted long URLs and users can toggle whether a long URL is trusted or
    not
- [ ] Replace Docker with Podman
