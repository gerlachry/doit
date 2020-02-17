from multiarch/crossbuild

RUN wget https://dl.google.com/go/go1.13.8.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.13.8.linux-amd64.tar.gz
ENV PATH $PATH:/usr/local/go/bin

