package downtools

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type DownImage struct {
	Wg *sync.WaitGroup
}

func (d DownImage) Down(url string) {
	
	re, _ := http.NewRequest("get", url, nil)
	cli := &http.Client{}
	res, err := cli.Do(re)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	newFile(res.Body, "jpg")
}

/*
imageType 为图片类型
*/
func newFile(r io.Reader, imageType string) {
	//随机
	rand.Seed(int64(time.Now().Nanosecond()))
	randStr := strconv.Itoa(rand.Int())
	f, err := os.Create("./Image/" + randStr + "." + imageType)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	io.Copy(f, r)
}
