#!/bin/bash

cd "$(dirname "$0")"

echo '输入文章标题'
read filename
fname=`echo $filename | ./dist/NewPost`
open $fname

echo 'press enter to exit'
read
