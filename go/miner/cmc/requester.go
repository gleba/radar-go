package cmc

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"radar.cash/core/hand"
	"time"
)

var rqCounter []int64

func opSec() int {
	now := time.Now().Unix()
	var op = 0
	for n, t := range rqCounter {
		if now-t < 60 {
			op++
		} else {
			rqCounter = append(rqCounter[:n], rqCounter[n+1:]...)
		}
	}
	//fmt.Println(op)
	return op
}
func request(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	rqCounter = append(rqCounter, time.Now().Unix())
	hand.Safe(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.66 Safari/537.36 Edg/80.0.361.40")
	req.Header.Set("Accept-Encoding", "gzip")
	opSec()
	//opNum := opSec()
	//if opNum > 180 && opNum < 220 {
	//	fmt.Println("↓ speed:", len(rqCounter), "req/min, now sleep for 3 seconds")
	//	time.Sleep(time.Duration(3) * time.Second)
	//} else if opNum > 230 {
	//	fmt.Println("↓ speed:", len(rqCounter), "req/min, now sleep for 10 seconds")
	//	time.Sleep(time.Duration(10) * time.Second)
	//}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("speed:", len(rqCounter), "req/min, now sleep for 10 seconds")
		fmt.Println(err)
		fmt.Println("request error ", len(rqCounter), "req/min, now s sleep for 60 seconds")
		time.Sleep(time.Duration(60) * time.Second)
		return request(url)
	}
	switch resp.StatusCode {
	case 500, 501, 502, 503, 504:
		fmt.Println(url)
		fmt.Println("Guru Meditation")
		return nil
	case 200:
	default:
		fmt.Println(url)
		fmt.Println(resp.Status)
		fmt.Println("speed:", len(rqCounter), "req/min, now sleep for 42 seconds")
		time.Sleep(time.Duration(42) * time.Second)
		return request(url)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		log.Println("BODY ERROR:", url)
	}
	zr, err := gzip.NewReader(bytes.NewReader(body))
	hand.Safe(err)
	defer zr.Close()
	var unpacked []byte
	unpacked, err = ioutil.ReadAll(zr)
	if err != nil {
		fmt.Println("unzip error")
		fmt.Println("speed:", len(rqCounter), "req/min, now sleep for 8 seconds")
		time.Sleep(time.Duration(8) * time.Second)
		return request(url)
	}
	return unpacked
}
