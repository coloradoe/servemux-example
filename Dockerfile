FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Rodrigo DelPer <rodrigo.perez@zeropw.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Run the executable
CMD ["./main"]
