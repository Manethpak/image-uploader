# Build stage for Svelte frontend
FROM node:20.16.0-alpine AS frontend-builder
WORKDIR /app/frontend

# Copy frontend source
COPY frontend/package*.json ./
RUN yarn

COPY frontend/ ./
RUN yarn build


# Build stage for Go backend
FROM golang:1.19.13-bookworm AS backend-builder
WORKDIR /app/backend

# Copy backend source
COPY backend/go.* ./
RUN go mod download

COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o server .


# Final stage
FROM alpine:3.18
WORKDIR /app/backend

# Copy built artifacts
COPY --from=frontend-builder /app/frontend/dist /app/frontend/dist
COPY --from=backend-builder /app/backend/server /app/backend/server

ENV GIN_MODE=release
ENV PORT=8080

# Expose port
EXPOSE 8080

VOLUME [ "./public" ]

# Run the server
CMD ["./server"]