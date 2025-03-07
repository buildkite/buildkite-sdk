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
RUN mise install node@latest && mise use --global node@latest
RUN mise install python@3.9.5 && mise use --global python@3.9.5
RUN mise install go@latest && mise use --global go@latest

# Install Python tools.
RUN pip install --no-cache-dir uv black

FROM base

WORKDIR /app

COPY . .

RUN mise trust && npm install

# Override the default command (irb).
CMD ["true"]
