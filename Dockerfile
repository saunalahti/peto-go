FROM golang:1.22.5-alpine AS builder

WORKDIR /go/src/peto-go

COPY go.* ./
RUN go mod download

# Copy only the necessary files and folders excluding those listed in .dockerignore
COPY . .

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -v -o /peto-go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /peto-go /app/peto-go

EXPOSE 3000

CMD ["./peto-go"]