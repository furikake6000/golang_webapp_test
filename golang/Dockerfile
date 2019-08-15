FROM centos:centos7.6.1810
MAINTAINER furikake6000

# update packages
RUN yum -y update

# install git
RUN yum -y install \
        gcc \
        curl-devel \
        expat-devel \
        gettext-devel \
        openssl-devel \
        zlib-devel \
        perl-ExtUtils-MakeMaker \
        autoconf \
        wget \
        make
WORKDIR /tmp
RUN wget https://www.kernel.org/pub/software/scm/git/git-2.19.0.tar.gz
RUN tar -xzvf git-2.19.0.tar.gz
WORKDIR /tmp/git-2.19.0
RUN make prefix=/usr/local all
RUN make prefix=/usr/local install

# install golang compiler
RUN wget https://storage.googleapis.com/golang/go1.12.7.linux-amd64.tar.gz
RUN tar -C /usr/local -xzvf go1.12.7.linux-amd64.tar.gz
ENV PATH $PATH:/usr/local/go/bin
WORKDIR /opt
RUN git clone https://go.googlesource.com/go
WORKDIR /opt/go
RUN git checkout go1.12
WORKDIR /opt/go/src
RUN GOROOT_BOOTSTRAP=/usr/local/go ./all.bash
ENV PATH $PATH:/opt/go/bin
RUN mkdir -p /root/go
ENV GOPATH /root/go
WORKDIR /root/go

# install golang packages
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gin-gonic/contrib/sessions
RUN go get github.com/gomodule/oauth1/oauth
