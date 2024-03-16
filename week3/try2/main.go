package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
)

type response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Item []struct {
			SchoolID int    `json:"school_id"`
			Name     string `json:"name"`
		} `json:"item"`
		NumFound int `json:"numFound"`
	} `json:"data"`
	Location  string `json:"location"`
	Encrydata string `json:"encrydata"`
}

type school struct {
	SchoolID int    `json:"school_id"`
	Name     string `json:"name"`
}

var (
	ch      = make(chan response)
	schools = make([]school, 0)
	mu      sync.RWMutex
	num     int64
	wg      sync.WaitGroup
	end     = make(chan struct{})
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(2)
	go GetURL(cancel)
	go Recieve(ctx)
	go check(cancel) //检查数量是否达标
	wg.Wait()

	jsondata, err := json.MarshalIndent(schools, "", "   ")
	if err != nil {
		fmt.Println("JSON marshaling error:", err)
		return
	}

	if err := ioutil.WriteFile("schools.json", jsondata, 0644); err != nil {
		fmt.Println("Writing to file failed:", err)
		return
	}
}

func GetURL(cancel context.CancelFunc) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		url := fmt.Sprintf("https://api.zjzw.cn/web/api/?keyword=&page=%d&province_id=&ranktype=&request_type=1&size=20&top_school_id=[2461,436]&type=&uri=apidata/api/gkv3/school/lists&signsafe=9326e2339790781062a5aac6ac933f66", i)
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching URL:", err)
			continue
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}

		var r response
		if err := json.Unmarshal(body, &r); err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			continue
		}
		ch <- r
	}

}

func Recieve(ctx context.Context) {
	defer func() {
		wg.Done()
		close(ch)
	}()

	for {
		select {
		case r := <-ch:
			for _, item := range r.Data.Item {
				id := item.SchoolID
				name := item.Name
				var s = school{
					SchoolID: id,
					Name:     name,
				}
				go Addsch(s)
			}
		case <-ctx.Done():
			return
		}
	}
}

func Addsch(s school) {
	schools = append(schools, s)
	atomic.AddInt64(&num, 1)
	fmt.Printf("第%d个大学（%s）被录入\n", atomic.LoadInt64(&num), s.Name)
	if atomic.LoadInt64(&num) == 200 {
		end <- struct{}{}
	}
}
func check(cancel context.CancelFunc) {
	select {
	case <-end:
		fmt.Println("所有学校已录入完成")
		cancel()
		return
	}
}