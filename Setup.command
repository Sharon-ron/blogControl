#!/bin/bash

cd "$(dirname "$0")"

git config --global user.name "Sharon-ron"
git config --global user.email "xiaoyunwang1991@sina.com"
git clone https://github.com/Sharon-ron/Sharon-ron.github.io
git config --global credential.helper osxkeychain

echo 'press enter to exit'
read