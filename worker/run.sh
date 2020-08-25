#!/bin/sh

#set -ex

cd /data

check() {
  flvmeta -C $VIDEO_NAME
  if [ "$?" != "0" ]; then
    tmp_file="tmp_${VIDEO_NAME}"
    mv $VIDEO_NAME $tmp_file
    flvmeta $tmp_file $VIDEO_NAME
    rm $tmp_file
  fi
}

if [ -f "$VIDEO_NAME" ]; then
  echo "File \"$VIDEO_NAME\" exists"
else
  trap "check" EXIT
  rtmpdump -v -m 60 -r $VIDEO_URL -o $VIDEO_NAME
fi
