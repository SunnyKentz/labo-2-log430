FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY out/centre_logistique/ .

# Expose port 8091
EXPOSE 8091
# Set environment variables
ENV DB_PORT=5433
ENV DB_USER=postgres
ENV DB_PASSWORD=postgres

# Run the application
CMD ["./app"]