# BUILD

FROM golang:1.16-alpline as build

WORKDIR /app

COPY go.mod ./ 
COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /server

# DEPLOY

FROM gcr.io.distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]

