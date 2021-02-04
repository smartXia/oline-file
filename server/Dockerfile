FROM golang

RUN apt update
RUN apt -y install ffmpeg

ADD ./online_file /opt/online_file
ADD cmd/db/migration /opt/migration
# ADD ./lib /opt/lib

ENV ENVIRONMENT production

WORKDIR /opt/

CMD ["/opt/online_file"]
