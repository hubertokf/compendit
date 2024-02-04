FROM golang:1.21 AS development

# Install Go Air
RUN go install github.com/cosmtrek/air@latest

# Install apt packages
RUN apt update \
    && apt install -y make

WORKDIR /go/src/compendit

CMD ["air"]

FROM development AS build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -tags timetzdata -a -installsuffix cgo -o /go/bin/ ./cmd/...

FROM scratch AS compendit-runtime

# Copy Certificate
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/bin/compendit /go/bin/

ENTRYPOINT ["/go/bin/compendit"]
