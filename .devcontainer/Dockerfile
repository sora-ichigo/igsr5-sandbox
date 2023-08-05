FROM mcr.microsoft.com/devcontainers/base:ubuntu

##  Perform as root user
##-----------------------------------------------
USER root

#  Install packages
#-----------------------------------------------
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
        libxcb-randr0-dev libxcb-xtest0-dev libxcb-xinerama0-dev libxcb-shape0-dev libxcb-xkb-dev \
        libnss3-dev iptables libx11-dev libxtst6 sshfs libncurses5-dev libncursesw5-dev libpq-dev

#  Install packages (devlopment tools)
#-----------------------------------------------
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
        tmux neovim bat fzf ripgrep

#  Clean up apt-get cache
#-----------------------------------------------
RUN apt-get clean \
    && rm -rf /var/lib/apt/lists/*

##  Perform as vscode user
##-----------------------------------------------
USER vscode
ENV USER_HOME /home/vscode

#  Set PS1
#-----------------------------------------------
COPY bash/prompt.bash $USER_HOME/bash/prompt.bash
RUN echo 'source $HOME/bash/prompt.bash' >> $USER_HOME/.bashrc
