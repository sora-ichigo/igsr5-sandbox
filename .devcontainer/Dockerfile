FROM mcr.microsoft.com/devcontainers/base:ubuntu

##  Perform as root user
##-----------------------------------------------
USER root

##  Perform as vscode user
##-----------------------------------------------
USER vscode
ENV USER_HOME /home/vscode

#  Set PS1
#-----------------------------------------------
COPY bash/prompt.bash $USER_HOME/bash/prompt.bash
RUN echo 'source $HOME/bash/prompt.bash' >> $USER_HOME/.bashrc
