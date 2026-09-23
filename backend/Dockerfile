FROM golang:1.27-alpine AS builder
WORKDIR /src

# Enable if project relies on C libraries
# RUN apk add --no-cache gcc musl-dev
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Set CGO_ENABLED=1 if project relies on C libraries
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/backend ./cmd/backend


FROM scratch
# In case you run into any problems with the scratch image move to distroless
# FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=builder /out/backend .
ENV TZ=Europe/Moscow
ENTRYPOINT ["/app/backend"]