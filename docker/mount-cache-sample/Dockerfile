# Dockerfile
FROM node:14

WORKDIR /app

COPY package.json .
COPY yarn.lock .

RUN --mount=type=cache,target=/app/.yarn-cache \
    yarn install --cache-folder /app/.yarn-cache

COPY . .

