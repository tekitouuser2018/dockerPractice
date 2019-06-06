FROM golang:latest

WORKDIR $HOME/go
# ADD . $HOME/go

CMD ["go", "run", "main.go"]

ENV HOME /root
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH $HOME/go

RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gin-gonic/gin
RUN go get gopkg.in/gorp.v1



