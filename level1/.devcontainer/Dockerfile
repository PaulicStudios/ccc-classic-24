FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    python3 \
    python3-pip \
    python3-dev \
    build-essential \
    libssl-dev \
    libffi-dev \
    python3-setuptools \
    python3-venv \
    python3-wheel \
    curl \
    git \
    wget \
    zsh \
    nano \
    cmake \
    ripgrep \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get clean

RUN python3 -m pip install --upgrade pip

#change the default shell to zsh
RUN chsh -s /usr/bin/zsh

RUN curl https://sh.rustup.rs -sSf | zsh -s -- -y

RUN curl -sS https://starship.rs/install.sh | sh -s -- -b /usr/local/bin -y

RUN echo 'eval "$(starship init zsh)"' >> /root/.zshrc

COPY ./starship.toml /root/.config/starship.toml

RUN curl -LO https://github.com/neovim/neovim/releases/latest/download/nvim-linux64.tar.gz &&\
    rm -rf /opt/nvim && \
    tar -C /opt -xzf nvim-linux64.tar.gz && \
    echo 'export PATH="$PATH:/opt/nvim-linux64/bin"' >> /root/.zshrc

RUN git clone https://github.com/NvChad/starter ~/.config/nvim

RUN wget -q -O - https://git.io/vQhTU | bash

SHELL ["zsh"]

CMD ["zsh"]
