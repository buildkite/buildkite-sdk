# Build from ruby:latest, as it's Debian-based, and keeps us from having to build Ruby from source.
FROM ruby:latest AS base

RUN apt-get update && apt-get install -y \
    curl \
    git \
    zip \
    unzip \
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
RUN mise install python@latest && mise use --global python@latest
RUN mise install go@latest && mise use --global go@latest

# Install Python tools.
RUN pip install --no-cache-dir uv black
RUN npm install -g nx

# Install SDKMAN! for Java-based languages.
ENV SDKMAN_DIR="/usr/local/bin/sdkman"
RUN curl -s https://get.sdkman.io | bash

# Install Java and Groovy.
RUN bash -c "source $SDKMAN_DIR/bin/sdkman-init.sh && \
    sdk install java 17.0.10-tem && \
    sdk install groovy 4.0.26"

ENV JAVA_HOME="$SDKMAN_DIR/candidates/java/current"
ENV GROOVY_HOME="$SDKMAN_DIR/candidates/groovy/current"
ENV PATH="$JAVA_HOME/bin:$GROOVY_HOME/bin:$PATH"

# Override the default command (irb).
CMD ["true"]
