FROM ubuntu:focal-20201106
VOLUME ["/var/cache/apt-cacher-ng"]
ARG DEBIAN_FRONTEND=noninteractive

# Use the tini init system as the entrypoint.
ARG TINI_VERSION=v0.18.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static /tini
RUN chmod +x /tini

RUN apt-get update \
    && apt-get install --yes --no-install-recommends \
    apt-cacher-ng \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

EXPOSE 3142
ENTRYPOINT ["/tini", "--", "/bin/bash", "-c"]
CMD ["chmod 777 /var/cache/apt-cacher-ng \
    && /etc/init.d/apt-cacher-ng start \
    && tail -f /var/log/apt-cacher-ng/*"]
