FROM ubuntu:latest
COPY ./main  /bin/device-plugin
RUN apt-get install  util-linux bash
ENTRYPOINT ["/bin/device-plugin"]
