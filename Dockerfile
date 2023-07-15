FROM alpine:3.17.2

RUN addgroup -S forge4flow-core && adduser -S forge4flow-core -G forge4flow-core
USER forge4flow-core

WORKDIR ./
COPY ./forge4flow-core ./

ENTRYPOINT ["./forge4flow-core"]

EXPOSE 8000
