FROM golang:alpine

ENV PKGS "zip unzip git curl npm py3-pip openssl openssl-dev build-base autoconf automake libtool gcc-doc python3-dev neofetch make wget gcc ca-certificates llvm nano vim ruby-full ruby-dev libffi-dev libgcc libssl1.1 zlib"

RUN apk upgrade && \
    apk add --update $PKGS

### gh ###
RUN wget \
    https://github.com/cli/cli/releases/download/$(curl https://get-latest.onrender.com/cli/cli)/gh_$(curl https://get-latest.onrender.com/cli/cli/no-v)_linux_amd64.tar.gz \
    -O gh.tar.gz
RUN tar -xzf gh.tar.gz
RUN mv "gh_$(curl https://get-latest.onrender.com/cli/cli/no-v)_linux_amd64/bin/gh" /usr/bin
RUN rm -rf gh*

### pyenv ###
RUN pip install tld --ignore-installed six distlib --user
RUN curl https://raw.githubusercontent.com/pyenv/pyenv-installer/master/bin/pyenv-installer | bash

ENV PATH "$HOME/.pyenv/bin:$PATH"

RUN echo 'eval "$(pyenv init -)"' >> ~/.bashrc
RUN echo 'eval "$(pyenv virtualenv-init -)"' >> ~/.bashrc

RUN /bin/bash -c "bash"

### pipenv ###
RUN curl https://raw.githubusercontent.com/pypa/pipenv/master/get-pipenv.py | python3

### poetry ###
RUN curl -sSL https://install.python-poetry.org | python3 -

ENV PATH "/root/.poetry/bin:$PATH"

RUN echo 'eval "$(poetry env install -q)"' >> ~/.bashrc
RUN echo 'eval "$(poetry env shell -q)"' >> ~/.bashrc

RUN /bin/bash -c "bash"

### nodejs package managers ###
RUN npm i -g npm@latest yarn@latest pnpm@latest

### update bundler ###
RUN gem update bundler

### build ###

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o ./create-botway-bot

ENV PORT 7050

EXPOSE 7050

ENTRYPOINT ["/app/create-botway-bot"]
