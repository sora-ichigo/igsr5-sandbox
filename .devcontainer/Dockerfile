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
        tmux neovim fzf ripgrep

#  Clean up apt-get cache
#-----------------------------------------------
RUN apt-get clean \
    && rm -rf /var/lib/apt/lists/*

#  Install DB (postgresql)
#-----------------------------------------------
RUN echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list \
    && wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - \
    && apt-get update \
    && apt-get -y install postgresql-13 \
    && service postgresql start \
    && su postgres -c "psql --command \"CREATE ROLE vscode LOGIN SUPERUSER\"" \
    && sed -i 's/md5/trust/g'  /etc/postgresql/13/main/pg_hba.conf \
    && sed -i 's/peer/trust/g'  /etc/postgresql/13/main/pg_hba.conf

##  Perform as vscode user
##-----------------------------------------------
USER vscode
ENV USER_HOME /home/vscode

#  Set functions.bash
#-----------------------------------------------
COPY bash/functions.bash $USER_HOME/bash/functions.bash
RUN echo 'source $HOME/bash/functions.bash' >> $USER_HOME/.bashrc

#  Set .bash_history
#-----------------------------------------------
COPY bash/.bash_history $USER_HOME/.bash_history

#  Set PS1
#-----------------------------------------------
COPY bash/prompt.bash $USER_HOME/bash/prompt.bash
RUN echo 'source $HOME/bash/prompt.bash' >> $USER_HOME/.bashrc
