FROM amd64/ubuntu:latest
ENV PATH=$PATH:/usr/local/go/bin
#ENV CGO_ENABLED=0
COPY . .
RUN apt update
RUN apt-get -y install wget
RUN wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin
CMD  go run server.go