FROM ubuntu:16.04
RUN sed -i 's/archive.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list \
  && sed -i 's/security.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list \
  && sed -i 's/ports.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list
RUN apt update && apt install -y rtmpdump flvmeta && apt clean
WORKDIR /app/rtmp-recorder-worker
COPY . .
VOLUME /data
CMD [ "sh", "run.sh" ]
