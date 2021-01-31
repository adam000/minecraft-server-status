####
# Base Go build
####

FROM golang:1.15.8 as build
ENV GO111MODULE on

# Warm up the module cache.
# Only copy in go.mod and go.sum to increase Docker cache hit rate.
COPY go.mod go.sum /src/
WORKDIR /src
RUN go mod download

COPY . /src

WORKDIR /src/app

RUN go build -v -o app

WORKDIR /src/Batchcraft

RUN go build -v -o Batchcraft

####
# Final image
####

FROM gcr.io/distroless/base-debian10:debug

COPY --from=build /src/app/ /app/
COPY --from=build /src/Batchcraft/Batchcraft /bin/Batchcraft

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["./app"]
