# First stage: build the application
FROM golang:1.22 AS builder

LABEL Manish Kumar <mskumar07@gmail.com>

# Set the working directory
WORKDIR /build

# Copy the source code into the container
COPY . .

# Set environment values
ENV DB_CONNECTION_STRING="host=postgres user=postgres password=mysecretpassword dbname=exoplanet-data port=5432 sslmode=disable TimeZone=Asia/Shanghai"

# Download all dependencies
RUN go mod download

# Build the application
RUN go build -o ./planetsApis



# Second stage: create the final image
FROM gcr.io/distroless/base-debian12

# Set the working directory
WORKDIR /app

# Copy the built application from the builder stage
COPY --from=builder /build/planetsApis ./planetsApis
COPY --from=builder /build/.env .env

ENV DB_CONNECTION_STRING="host=postgres user=postgres password=mysecretpassword dbname=exoplanet-data port=5432 sslmode=disable TimeZone=Asia/Shanghai"

# Expose the necessary port
EXPOSE 8080

# Run go build
CMD ["/app/planetsApis"]