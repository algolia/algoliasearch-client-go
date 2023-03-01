ARG NODE_VERSION=18.14.2
ARG JAVA_VERSION=11.0.18
ARG PHP_VERSION=8.1.16

# PHP is so complicated (and long) to install that we use the docker image directly
FROM php:${PHP_VERSION}-bullseye

ARG NODE_VERSION
ARG JAVA_VERSION

ENV DOCKER=true

# use bash for subsequent commands
SHELL ["/bin/bash", "--login", "-c"]

# PHP composer
COPY --from=composer:latest /usr/bin/composer /usr/local/bin/composer

RUN apt-get update && apt-get install -y \
    curl \
    zip \
    unzip \
    # python is used by nvm to install some packages
    python3 \
    && rm -rf /var/lib/apt/lists/*

# Javascript (node)
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash
RUN nvm install ${NODE_VERSION}
RUN npm install -g yarn 

# Java
RUN curl -s "https://get.sdkman.io" | bash
RUN source "$HOME/.sdkman/bin/sdkman-init.sh"
RUN sdk install java ${JAVA_VERSION}-zulu

# Java formatter
ADD https://github.com/google/google-java-format/releases/download/v1.15.0/google-java-format-1.15.0-all-deps.jar /tmp/java-formatter.jar

WORKDIR /app

CMD bash
