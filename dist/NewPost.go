package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func fail(err error) {
	log.Fatalln("错误！", err)
}

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fail(err)
	}
	defer os.Stdin.Close()
	title2 := string(b[:len(b)-1])

	t := time.Now().Local()
	title1 := t.Format("2006-01-02")
	if title2 == "" {
		title2 = t.Format("15:04:05")
	}
	title := title1 + "-" + title2

	fileHeader := `---
layout: post
title: ` + title2 + `
description: ""
category:
tags: []
---
{% include JB/setup %}

<!--1. title是文章标题，  tags是标签 多个标签用空格分开-->
<!--2. 以下是你的正文-->
`
	filename := "Sharon-ron.github.io/_posts/" + title + ".md"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		fail(err)
	}
	defer f.Close()
	_, err = io.WriteString(f, fileHeader)
	if err != nil {
		fail(err)
	}
	fmt.Println(filename)
}
