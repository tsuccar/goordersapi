FROM golang:1.21

ARG USERNAME=golang-app
ARG USER_UID=1002
ARG USER_GID=$USER_UID

WORKDIR /app

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME 

ENV PATH="/home/${USERNAME}/.local/bin:${PATH}"


USER $USERNAME  

CMD ["/bin/bash"]
