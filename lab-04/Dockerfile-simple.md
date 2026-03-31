FROM ubuntu:22.04

LABEL maintainer="My Name"

RUN apt update && apt install -y curl

CMD ["echo", "Hello from My Name Container"]