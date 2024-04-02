ARG GANACHE_MNEMONIC="ozone access unlock valid olympic save include omit supply green clown session"

FROM vegaprotocol/smartcontracts-base:v1.5.0

ARG TARGETARCH
ARG GANACHE_MNEMONIC

RUN pwd
RUN ls
COPY ./prebuild .
RUN ls
RUN cat prebuild
RUN /bin/sh prebuild