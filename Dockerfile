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
RUN mise install node@latest && \
    mise use --global node@latest

# Python version support: https://devguide.python.org/versions/
RUN mise install python@latest && \
    mise use --global python@latest

# Go version support: https://endoflife.date/go
RUN mise install go@latest && \
    mise use --global go@latest

# Ruby version support: https://www.ruby-lang.org/en/downloads/branches/
RUN mise install ruby@latest && \
    mise use --global ruby@latest

# Install Python tools.
RUN pip install --no-cache-dir uv black
RUN npm install -g nx

# Override the default command (irb).
CMD ["true"]
