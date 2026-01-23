## Initial Project Implementation/Design:

The expected flow of the URL Shortener:

**User -> Load Balancer -> Web Server (w/ URL Shortener) -> Database**

There will be users that _request the creation of a short URL_.

The **Load Balancer** will handle incoming requests of _creating and redirecting URLs_.

A **web server** set up with the **URL Shortener** service will handle the _logic of creating short URLs_, _storing URL mappings_, and _retrieving the original URL for redirection_.

## Tech Stack in Mind

The stack that will be used to build the URL shortener is:

- **Golang** as the primary language with **Fiber** as the web framework,
- A **Redis** container as a caching layer,
- **MySQL** database to store URLs and their relationships,
- **Docker** for containerized deployment, and
- **Nginx** to simulate **load balancing**.
