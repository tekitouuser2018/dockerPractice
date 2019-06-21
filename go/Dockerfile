FROM golang:1.11.10-alpine

EXPOSE 8080

WORKDIR $HOME/go
# ADD . $HOME/go

CMD ["go", "run", "main.go"]

ENV HOME /root
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH $HOME/go

RUN git --version
# RUN git clone https://github.com/git/git
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gin-gonic/gin
RUN go get gopkg.in/gorp.v1



