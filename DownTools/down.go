package downtools

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var D DownImage

type DownImage struct {
	Wg *sync.WaitGroup
}

func (d DownImage) Down(url string) {
	defer d.Wg.Done()

	re, _ := http.NewRequest("", url, nil)
	fmt.Println(re)
	//这里需要设置请求头信息
	cli := &http.Client{}
	res, err := cli.Do(re)
	if err != nil || res.StatusCode != 200 {
		log.Fatalln(res.Status)
		return
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
