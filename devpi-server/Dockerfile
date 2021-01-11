FROM ubuntu:focal-20201106
ARG DEBIAN_FRONTEND=noninteractive

# Use the tini init system as the entrypoint.
ARG TINI_VERSION=v0.18.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static /tini
RUN chmod +x /tini

# Install necessary Debian packages
RUN apt-get update \
    && apt-get install --yes --no-install-recommends \
    # Install Python 3
    python3.8 \
    python3-pip \
    ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY requirements.txt /requirements.txt
RUN python3.8 -m pip --no-cache-dir install -r /requirements.txt

WORKDIR /devpi
RUN devpi-init --serverdir /devpi
EXPOSE 3141
ENTRYPOINT ["/tini", "--", "/usr/local/bin/devpi-server"]
CMD ["--host=0.0.0.0", "--port=3141", "--serverdir=/devpi"]
