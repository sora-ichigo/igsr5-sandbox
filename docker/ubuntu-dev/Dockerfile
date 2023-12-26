FROM --platform=linux/amd64 ubuntu:20.04

RUN apt-get update && apt-get install -y \
  neovim \
  tmux \
  fzf \
  sudo \
  gcc


RUN useradd -m dev && echo "dev:dev" | chpasswd && adduser dev sudo
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
USER dev

WORKDIR /$HOME/work
