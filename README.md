# Optimal Dockerfile practices

This is a demo repository focused on demonstrating how to write a good Dockerfile.

In [`hello/Dockerfile`](https://github.com/jamesmishra/optimal-dockerfiles/blob/main/hello/Dockerfile), we set up a simple HTTP server using a Docker multi-stage build. We also use the `tini` init system and set up optional caching proxies for Apt and PyPI packages.

We set up a local Apt caching server at [`apt-cacher-ng/Dockerfile`](https://github.com/jamesmishra/optimal-dockerfiles/blob/main/apt-cacher-ng/Dockerfile) and a PyPI caching server at [`devpi-server/Dockerfile`](https://github.com/jamesmishra/optimal-dockerfiles/blob/main/devpi-server/Dockerfile).

These three Docker images are configured and built via the [`docker-compose.yml`](https://github.com/jamesmishra/optimal-dockerfiles/blob/main/docker-compose.yml) at the root of this repository.

The [`Makefile`](https://github.com/jamesmishra/optimal-dockerfiles/blob/main/Makefile) at the root of this repository can be used to build the containers [`make build`]--starting the caching proxies before building `hello`. The `make up` command starts the `hello` HTTP server on the local machine which would serve traffic at `http://localhost:8080/hello`.

All of the files are heavily commented, and more explanation of how these Docker features work is available [at this blog post I wrote](https://jamesmishra.com/2021/01/09/docker-tips/).
