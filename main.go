package main

import (
	"demo/downtools"
	"sync"
)

var url = []string{
	"https://img0.baidu.com/it/u=2631815445,1952611015&fm=253&fmt=auto&app=120&f=JPEG?w=1280&h=800",
	"https://up.deskcity.org/pic_source/2f/f4/42/2ff442798331f6cc6005098766304e39.jpg",
}

func main() {
	var wg sync.WaitGroup
	wg.Add(len(url))
	var d = downtools.D
	d.Wg = &wg
	for _, v := range url {
		go d.Down(v)
	}
	wg.Wait()

}
