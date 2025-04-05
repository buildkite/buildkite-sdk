#!/bin/bash

# Initialize SDKMAN
source "$HOME/.sdkman/bin/sdkman-init.sh"

# Use Java 17 from SDKMAN
sdk use java 17.0.10-tem

# Use Groovy from SDKMAN
sdk use groovy 4.0.26

# Get the absolute path of the SDK directory
SDK_GROOVY_DIR=$(cd ../../sdk/groovy && pwd)

# Create output directory
mkdir -p ../../out/apps/groovy

# Run the Groovy script with the SDK groovy directory in the classpath
echo "Running Groovy script with classpath: $SDK_GROOVY_DIR"
groovy -cp "$SDK_GROOVY_DIR" main.groovy

# If successful, show the generated files
if [ $? -eq 0 ]; then
  echo "Success! Generated files:"
  ls -la ./out/apps/groovy/
fi
