FROM golang:1.11.10-alpine AS build
# FROM alpine:3.9 AS build
# FROM golang:1.11
EXPOSE 8080

WORKDIR /root/go
# ADD . $HOME/go

# CMD ["go", "run", "main.go"]
COPY . ./

ENV CGO_ENABLED=0 \
    GO111MODULE=auto
ENV HOME /root
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:$PATH
ENV PATH /usr/local/bin:$PATH
ENV PATH /usr/local/sbin:$PATH
ENV PATH /usr/bin:$PATH
ENV GOROOT /usr/local/go


# COPY go.mod ./
# COPY go.sum ./

# RUN grep -H "" /etc/*version ; grep -H "" /etc/*release
# RUN curl -o apk-tools-static-2.6.8-r1.apk http://dl-cdn.alpinelinux.org/alpine/v3.5/main/x86_64/apk-tools-static-2.6.9-r0.apk
# RUN curl http://dl-cdn.alpinelinux.org/alpine/v3.5/main/x86_64/apk-tools-static-2.6.8-r1.apk
# RUN tar -zxvf apk-tools-static-2.6.8-r1.apk
# RUN apk update
RUN apk add bash ca-certificates git gcc g++ libc-dev
# RUN go mod download
# RUN echo $PATH
# RUN apk update && apk add git && go get gopkg.in/natefinch/lumberjack.v2
# RUN apk add --no-cache bash
RUN apk update && apk add git
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gin-gonic/gin
RUN go get gopkg.in/gorp.v1

# COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o binary 

FROM alpine:3.7

WORKDIR /usr/local/bin/

COPY --from=build /root/go /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/"]

# RUN cd /usr/local/bin/ && go run main.go

# CMD ["go", "run", "main.go"]