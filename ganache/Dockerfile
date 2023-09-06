ARG GANACHE_MNEMONIC="ozone access unlock valid olympic save include omit supply green clown session"

FROM vegaprotocol/smartcontracts-base:v1.2.4

ARG TARGETARCH
ARG GANACHE_MNEMONIC


COPY ./prebuild /

COPY ./entrypoint.sh /entrypoint.sh
COPY ./supervisord.conf /supervisord.conf


RUN /prebuild
RUN apk add py3-pip && pip install supervisor
ENTRYPOINT [ "/entrypoint.sh" ]