# STEP 1: build executable binary

# Pull golang image
FROM golang:1.18-alpine as build

# Additional Label
LABEL maintainer="Michael Stevan Lapandio<michaellapandio04@gmail.com>"

# Add a work directoryer
WORKDIR /app
# Install make
RUN apk add --no-cache bash make gcc libc-dev

# Cache and install dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# Copy app files
COPY . .
RUN cp -rf ./.env.example ./.env
# Build app
RUN go build -o agmc-day-7

# step 2: build a small image
FROM alpine:3.16.0
RUN apk add bash build-base gcompat
COPY --from=build /app/agmc-day-7 .
# Expose port
EXPOSE 8000
CMD ["/agmc-day-7"]