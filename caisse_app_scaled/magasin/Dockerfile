FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY out/magasin/ .

EXPOSE 8080

# Set environment variables
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=postgres

# Run the application
CMD ["./app"]