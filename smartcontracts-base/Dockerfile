FROM ethereum/solc:0.8.8 AS solc088
FROM ethereum/solc:0.8.1 AS solc081

FROM node:20.3-alpine AS builder
RUN mkdir /app
WORKDIR /app
ENV PATH="$PATH:/app/node_modules/.bin"
RUN apk update
RUN apk add ca-certificates curl
RUN apk add gcc g++ git make python3
RUN \
	npm install -g node-gyp@latest --force
ADD package.json package-lock.json ./
RUN npm install

# # #

FROM trufflesuite/ganache:v7.9.2
LABEL org.opencontainers.image.source https://github.com/vegaprotocol/devops-infra
WORKDIR /app2
ENV PATH="$PATH:/app2/node_modules/.bin"

COPY --from=solc088 /usr/bin/solc /usr/bin/solc-v0.8.8
COPY --from=solc081 /usr/bin/solc /usr/bin/solc-v0.8.1

COPY --from=builder /app/package*.json ./
COPY --from=builder /app/node_modules ./node_modules

ADD contracts-solc-0.8.1 ./contracts-solc-0.8.1
ADD contracts-solc-0.8.8 ./contracts-solc-0.8.8
ADD scripts ./scripts
ADD run ./run