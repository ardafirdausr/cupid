FROM golang:1.21-alpine3.20 as builder

LABEL maintainer="Arda <arda@gmail.com>"

# Update alpine repository indexs and next install git, openssh-client, and gcc
RUN apk update && apk add --no-cache git openssh-client build-base

# go get uses git internally. The following one liners will make git and consequently go get clone your package via SSH.
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

# add credentials on build
ARG SSH_PRIVATE_KEY
RUN mkdir /root/.ssh/
RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa
RUN echo "StrictHostKeyChecking no " > /root/.ssh/config
RUN chmod 400 /root/.ssh/id_rsa

# make sure your domain is accepted
RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

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