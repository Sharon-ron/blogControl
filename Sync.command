#!/bin/bash

cd "$(dirname "$0")"

t="update at "`date +"%D-%T"`
echo $t
cd Sharon-ron.github.io && git add . && git commit -am "$t" && git push origin master

echo 'press enter to exit'
read