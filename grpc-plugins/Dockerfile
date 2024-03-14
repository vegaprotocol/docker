FROM ubuntu:22.04 AS grpc_plugins
ARG GRPC_VER=1.41.0


RUN apt update && DEBIAN_FRONTEND=noninteractive apt install -y --no-install-recommends \
	autoconf \
	build-essential \
	ca-certificates \
	cmake \
	git \
	libtool \
	pkg-config
RUN git clone -b "v$GRPC_VER" https://github.com/grpc/grpc
WORKDIR /grpc
RUN git submodule update --init
WORKDIR /grpc/cmake/build
RUN cmake ../..
RUN make -j "$(grep -c ^processor /proc/cpuinfo)"
RUN make install
RUN strip /usr/local/bin/protoc-* /usr/local/bin/grpc_*_plugin



FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive
ENV BUF_VER="1.0.0-rc12"

ARG UNAME=testuser
ARG UID=1000
ARG GID=1000

RUN groupadd -g $GID -o $UNAME
RUN useradd -m -u $UID -g $GID -o -s /bin/bash $UNAME
RUN chmod 1777 /tmp

RUN apt update \
	&& apt install -y --no-install-recommends \
		ca-certificates \
		curl \
        wget \
		git \
		jq \
		make \
		openssh-client \
	&& rm -rf /var/lib/apt/lists/*


COPY --from=grpc_plugins /usr/local/bin/protoc* /usr/local/bin/
COPY --from=grpc_plugins /usr/local/bin/grpc_python_plugin /usr/local/bin/

RUN wget -q -O /usr/local/bin/buf https://github.com/bufbuild/buf/releases/download/v$BUF_VER/buf-Linux-$(uname -m) \
	&& chmod 0755 /usr/local/bin/buf

WORKDIR /workspace

