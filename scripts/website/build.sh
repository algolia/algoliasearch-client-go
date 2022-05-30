#!/bin/bash

# treat website as independant yarn project
touch website/yarn.lock

# build doc specs
yarn website:build-specs

# install website deps and build
cd website && yarn install && yarn build
