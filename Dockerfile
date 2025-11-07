# Build from ruby:latest, as it's Debian-based, and keeps us from having to build Ruby from source.
FROM ruby:latest AS base

RUN apt-get update && apt-get install -y \
    curl \
    git \
    && rm -rf /var/lib/apt/lists/*

# Use Bash, and fail on any error.
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# Install Mise for tool-version management. (https://mise.jdx.dev).
ENV MISE_DATA_DIR="/mise"
ENV MISE_CONFIG_DIR="/mise"
ENV MISE_CACHE_DIR="/mise/cache"
ENV MISE_INSTALL_PATH="/usr/local/bin/mise"
ENV PATH="/mise/shims:$PATH"
RUN curl https://mise.run | sh

# Install Node.js, Python, and Go.

# Node.js version support: https://nodejs.org/en/about/previous-releases
RUN mise install node@20 && \
    mise install node@21 && \
    mise install node@22 && \
    mise install node@23 && \
    # Current LTS Version
    mise install node@24 && \
    mise install node@25 && \
    mise use --global node@24

# Python version support: https://devguide.python.org/versions/
RUN mise install python@3.10 && \
    mise install python@3.11 && \
    mise install python@3.12 && \
    mise install python@3.13 && \
    # Current Version
    mise install python@3.14 && \
    mise use --global python@3.14

# Go version support: https://endoflife.date/go
RUN mise install go@1.24 && \
    # Current Version
    mise install go@1.25 && \
    mise use --global go@1.25

# Ruby version support: https://www.ruby-lang.org/en/downloads/branches/
# Go version support: https://endoflife.date/go
RUN mise install ruby@3.2 && \
    mise install ruby@3.3 && \
    # Current Version
    mise install ruby@3.4 && \
    mise use --global ruby@3.4

# Install Python tools.
RUN pip install --no-cache-dir uv black
RUN npm install -g nx

# Override the default command (irb).
CMD ["true"]
