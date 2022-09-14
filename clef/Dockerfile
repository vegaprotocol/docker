FROM golang:1.19-alpine as build

ARG TAG="clef-add-new-mime-type"

RUN apk add --no-cache make gcc musl-dev linux-headers git

COPY main.go main.go
RUN go build -o clef-runner main.go

RUN git clone https://github.com/vegaprotocol/go-ethereum.git && cd /go/go-ethereum && \
    if [[ -n $TAG ]]; then git checkout $TAG; fi && env GO111MODULE=on go run build/ci.go install ./cmd/clef


FROM alpine:latest as runtime

RUN apk add --no-cache ca-certificates
COPY --from=build /go/go-ethereum/build/bin/clef /usr/local/bin/clef
COPY --from=build /go/clef-runner /usr/local/bin/clef-runner

COPY rules.js /app/config/rules.js
COPY 4byte.json /app/config/4byte.json
COPY password /app/config/password
COPY data /app/data
COPY entrypoint.sh /entrypoint.sh

RUN /entrypoint.sh init

EXPOSE 8550

ENTRYPOINT ["/entrypoint.sh"]