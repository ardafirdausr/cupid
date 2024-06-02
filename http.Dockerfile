FROM golang:1.21-alpine3.20 as builder

LABEL maintainer="Arda <arda@gmail.com>"

# Update alpine repository indexs and next install git, openssh-client, and gcc
RUN apk update && apk add --no-cache git openssh-client build-base

RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Resolve and download all dependencies.
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# RUN STAGE ############################################################
FROM alpine:3.16

# tzdata for timezone
RUN apk update && apk add tzdata
ENV TZ Asia/Jakarta

# Set working directory
WORKDIR /app

# Copy build file
COPY --from=builder /app/main .
EXPOSE 8000

# Command to run the executable
CMD ["./main"]