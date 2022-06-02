# syntax=docker/dockerfile:1
# build stage
FROM golang:1.18.3-alpine3.15 as builder

# set workdir
WORKDIR /src/pageshot

# install dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# build a binary
RUN CGO_ENABLED=0 go build -o bin/pageshot cmd/pageshot/main.go

# final image
FROM chromedp/headless-shell:102.0.5005.63

# set maintainer label
LABEL maintainer="Aleksandr Yakimenko <urlucidfall@gmail.com>"

# set the default server's port
ENV SERVER_PORT=8000

# update the apt-get lists, install ca-certificates, dumb-init, 
# delete the apt-get lists and update ca-certificates
RUN apt-get update && apt-get install --no-install-recommends -y ca-certificates=20210119 dumb-init=1.2.5-1 \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* \
    && update-ca-certificates

# set workdir
WORKDIR /opt/pageshot

# copy the binary from the builder stage
COPY --from=builder /src/pageshot/bin/pageshot pageshot

# expose port
EXPOSE ${SERVER_PORT}/tcp

# set an entrypoint to dumb-init in order to reap zombie processes
# https://github.com/chromedp/docker-headless-shell#using-as-a-base-image
ENTRYPOINT ["dumb-init", "--"]

# set an executable
CMD [ "/opt/pageshot/pageshot" ]