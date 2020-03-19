package utils

import "container/list"

var largeWorkList = list.New()
var smallWorkList = list.New()
var shutdownQueueChan = make(chan string)
var workChan = make(chan string, 1000)
