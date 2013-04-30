#!/bin/bash

cd "$(dirname "$0")"

echo '输入文件名'
read filename
fname=`echo $filename | ./dist/NewPost`
open $fname

echo 'press enter to exit'
read
