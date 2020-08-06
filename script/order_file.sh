#!/bin/bash
order_static()
{
  for file in /home/xsguo/static/static/*
  do
  if test -d $file
  then
  echo http://gpu604.aibee.cn:9090/v1/annotation/list?videoName=${file:26:48} >> ./video_url.csv
  fi
  done
}

rm -rf /home/xsguo/static/candidate

hdfscli download /bj_dev/user/store_solutions/prod/customer/SEPHORA/shanghai/qbwk/tracking/onboarding/20191025/mask15/label/candidate /home/xsguo/static
rm -rf  /home/xsguo/static/static
mv /home/xsguo/static/candidate /home/xsguo/static/static
rm -rf ./video_url.csv

order_static
