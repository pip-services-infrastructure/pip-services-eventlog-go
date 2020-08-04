# Start with the golang v1.13 image
FROM golang:1.13

# Set a working directory
WORKDIR /app

# Copy the package files
COPY go.mod go.sum ./

# Install all go_modules
RUN go mod download

# Copy the entire project
COPY . .

# Set environment variables
ENV HTTP_PORT=8080
ENV MONGO_SERVICE_HOST=
ENV MONGO_SERVICE_PORT=27017

# Expose the port
EXPOSE "8080:8080"

# Specify the command to run
CMD [ "./obj/app" ]