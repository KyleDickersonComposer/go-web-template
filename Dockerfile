FROM golang:1.24-alpine

# Install Node.js and npm
RUN apk add --no-cache nodejs npm git bash

WORKDIR /app

# Copy package files and install npm dependencies
COPY package.json package-lock.json ./
RUN npm ci

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source files
COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate

# Build frontend assets
RUN npx tailwindcss -i ./css/input.css -o ./dist/output.css
RUN npm run build

# Build Go application
RUN go build -o app .

EXPOSE 8080

CMD ["./app"]
