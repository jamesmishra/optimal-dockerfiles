# =======================================================================

# For the sake of security and reproducability, always specify
# the exact version of the Docker images you build on top of.

# This builder container has the Go compiler, which we will use
# to create a statically-linked binary. Files created in this stage
# will not be in the final image unless we copy them with the
# `COPY` command.
FROM golang:1.15.6-alpine3.12 as builder

# We copy our source code from the build server's filesystem into the
# container's filesystem.
COPY hello.go /app/

# We use the Go compiler in this image to create the statically-linked
# binary. Then we will not need the compiler.
# Note that the Alpine container has musl as its libc and the below
# Ubuntu container uses glibc. If you need to use cgo, then use
# the `golang:1.15.6-buster` image instead
WORKDIR /app
RUN CGO_ENABLED=0 go build hello.go



# =======================================================================

# Now we begin the final image that is used to build the containers.
# In this example, because we have created a statically-linked binary,
# we could have used Alpine or Busybox instead of Ubuntu to make a
# smaller container. However, if you are containerizing a Python
# library extensions depending on glibc, you may find it more convenient
# to use Ubuntu as your container's base image.
FROM ubuntu:focal-20201106
# We set this argument to prevent Apt commands from presenting us with
# interactive prompts to configure installed packages.
ARG DEBIAN_FRONTEND=noninteractive

# Don't run as root. Create a regular user.
ARG CONTAINER_UID=10000
ARG CONTAINER_GID=10000
ARG CONTAINER_USERNAME=anubis
ARG CONTAINER_GROUPNAME=anubis
RUN groupadd --gid=${CONTAINER_GID} --force ${CONTAINER_GROUPNAME} \
    && useradd --create-home --home-dir=/home/${CONTAINER_USERNAME} --shell=/bin/bash --uid=${CONTAINER_UID} --gid=${CONTAINER_GID} ${CONTAINER_USERNAME}

# The Docker ADD command can fetch files from a remote URL, which is
# handy to avoid having to install an HTTP client in your container.
# Here we download the tini init system.
ARG TINI_VERSION=v0.18.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static /tini
RUN chmod +x /tini

# Insert an Apt proxy in this container, but only if this build-time
# environment variable is defined.
ARG LOCAL_APT_CACHE_URL=
RUN if [ ! -z ${LOCAL_APT_CACHE_URL} ]; \
    then \
    # This first `find` command replaces HTTP Apt repository URL to run
    # through our apt-cacher-ng proxy.
    find /etc/apt/sources.list /etc/apt/sources.list.d/ \
    -type f -exec sed -Ei 's!http://!'${LOCAL_APT_CACHE_URL}'/!g' {} \; \
    # This second `find` commands replaces HTTPS Apt repository URLs
    # to run through out apt-cacher-ng proxy. If we don't do this,
    # then apt-cacher-ng will not be able to cache packages from
    # HTTPS repositories.
    && find /etc/apt/sources.list /etc/apt/sources.list.d/ \
    -type f -exec sed -Ei 's!https://!'${LOCAL_APT_CACHE_URL}'/HTTPS///!g' {} \; ;\
    fi

# Insert a PyPI proxy in this container, but ohly if this build-time
# variable is defined.
ARG LOCAL_PYPI_CACHE_URL=
ARG PYPI_CACHE_PROXY_CONFIG_FILE=/etc/xdg/pip/pip.conf
RUN if [ ! -z ${LOCAL_PYPI_CACHE_URL} ]; \
    then \
    mkdir -pv $(dirname ${PYPI_CACHE_PROXY_CONFIG_FILE}) \
    && echo "[global]" > ${PYPI_CACHE_PROXY_CONFIG_FILE} \
    && echo "index-url = ${LOCAL_PYPI_CACHE_URL}/root/pypi/+simple/" >> ${PYPI_CACHE_PROXY_CONFIG_FILE}; \
    fi

# Here we install our dependencies.
RUN apt-get update \
    # We run apt-get install with `--yes` to avoid asking for user input
    # and with `--no-install-recommends` to avoid installing extra
    # packages.
    && apt-get install --yes --no-install-recommends \
        # Many apps that make connections to the Internet
        # need a package of CA certificates. Our example app
        # doesn't, but your app probably does.
        ca-certificates \
        # We don't need python3 or pip either, other than to show
        # off how our PyPI caching server works.
        python3 \
        python3-pip \
    # Now we delete our Apt cache to shrink the size of our image.
    # However, this only works if the deletion happens in the same
    # RUN image as the `apt-get update` and `apt-get install`.
    # If you try to remove the cache in a separate RUN command,
    # you will not shrink the size of the image because the files
    # of each layer are separately stored.
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Let's pip install something to demonstrate our PyPI cacher.
# We use the `--no-cache-dir` to avoid increasing the size of this
# image by storing the pip cache.
RUN python3 -m pip --no-cache-dir install requests

# Here, we copy from the "builder" stage of the multi-stage build,
# as opposed to the host filesystem/.
COPY --from=builder /app/hello /bin/hello

# Set the init system to be the container entrypoint.
# The path provided in the Dockerfile `CMD` command will be passed
# as arguments to the `ENTRYPOINT` command.
ENTRYPOINT ["/tini", "--", "/bin/hello"]

# The `CMD` is passed to the entrypoint. It is common to place
# the path to a binary in the entry point and then the binary's
# arguments as the `CMD`.
# If you want to change the listen address, you can do this by
# overriding the `CMD` but leaving the entrypoint untouched.
CMD ["-listen-address=:8080"]

# The `EXPOSE` Dockerfile command is optional. Processes in the container
# can serve traffic on ports tha have not been marked with the `EXPOSE`
# command, but it is good documentation to use `EXPOSE`.
EXPOSE 8080

# Switch to our non-root user.
USER anubis
