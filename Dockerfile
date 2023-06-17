FROM alpine:3.17.2

RUN addgroup -S auth4flow-core && adduser -S auth4flow-core -G auth4flow-core
USER auth4flow-core

WORKDIR ./
COPY ./auth4flow-core ./

ENTRYPOINT ["./auth4flow-core"]

EXPOSE 8000
