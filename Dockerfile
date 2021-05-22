FROM golang:latest
RUN mkdir /dictionary
WORKDIR /go/bin
# Using latest release binary
RUN wget https://github.com/looCiprian/anagram-cheater/releases/download/v1.0/anagram-cheater_linux-amd64 -O anagram-cheater
RUN chmod +x anagram-cheater
WORKDIR /go
ENTRYPOINT anagram-cheater --help
ENTRYPOINT /bin/bash