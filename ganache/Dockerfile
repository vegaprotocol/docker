ARG ganache_mnemonic="ozone access unlock valid olympic save include omit supply green clown session"

FROM vegaprotocol/smartcontracts-base:v1.0.1
ARG ganache_mnemonic
ENV GANACHE_MNEMONIC=${ganache_mnemonic}


COPY ./prebuild /

RUN /prebuild