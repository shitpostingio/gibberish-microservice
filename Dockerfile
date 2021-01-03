# First stage: build the executable.
FROM golang:buster AS builder

# It is important that these ARG's are defined after the FROM statement
ARG SSH_PRIV="nothing"
ARG SSH_PUB="nothing"
ARG GOSUMDB=off

# Create the user and group files that will be used in the running 
# container to run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'gibber:x:65534:65534:gibber:/:' > /user/passwd && \
    echo 'gibber:x:65534:' > /user/group

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/gitlab.com/shitposting/gibberish-microservice

# Import the code from the context.
COPY .  .

# Build the executable
RUN go install

# Finale stage
FROM golang:buster

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Copy the built executable
COPY --from=builder /go/bin/gibberish-microservice /home/gibber/gibberish

RUN chown -R gibber /home/gibber

# Set the workdir
WORKDIR /home/gibber

# Copy the knowledge.json
COPY ./knowledge.json .

# Perform any further action as an unprivileged user.
USER gibber:gibber

# Expose port 10002
EXPOSE 10002

# Run the compiled binary.
CMD ["./gibberish"]
