package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

const (
	baseURL       = "https://www.gaokao.cn/school/search"
	schoolURL     = "https://www.gaokao.cn/school/"
	detailURL     = "/introDetails"
	numPages      = 10
	numSchools    = 20
	maxGoroutines = 10
)

var (
	schoolsData = make(map[string]string)
	mutex       sync.Mutex
)

func getPage(page int) ([]byte, error) {
	resp, err := http.Get(baseURL + "?p=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getSchoolInfo(schoolID string, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{}
	defer func() {
		<-sem
	}()

	detailResp, err := http.Get(schoolURL + schoolID + detailURL)
	if err != nil {
		fmt.Printf("获取学校 %s 详情页面时出错: %v\n", schoolID, err)
		return
	}
	defer detailResp.Body.Close()

	detailBody, err := ioutil.ReadAll(detailResp.Body)
	if err != nil {
		fmt.Printf("读取学校 %s 详情页面内容时出错: %v\n", schoolID, err)
		return
	}

	mutex.Lock()
	schoolsData[schoolID] = string(detailBody)
	fmt.Printf("成功获取学校 %s 详情页面内容\n", schoolID)
	mutex.Unlock()
}

func crawlPages(wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()

	for i := 1; i <= numPages; i++ {
		_, err := getPage(i)
		if err != nil {
			fmt.Printf("获取第 %d 页时出错: %v\n", i, err)
			continue
		}

		for j := 1; j <= numSchools; j++ {
			schoolID := strconv.Itoa((i-1)*numSchools + j)
			wg.Add(1)
			go getSchoolInfo(schoolID, wg, sem)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxGoroutines)

	wg.Add(1)
	go crawlPages(&wg, sem)

	wg.Wait()

	err := writeDataToFile("03.txt", schoolsData)
	if err != nil {
		fmt.Println("写入文件时出错:", err)
		return
	}

	fmt.Println("学校信息已保存到文件 03.txt 中")
}

func writeDataToFile(filename string, data map[string]string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := filepath.Join(dir, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for id, info := range data {
		_, err := file.WriteString(fmt.Sprintf("学校ID: %s\n", id))
		if err != nil {
			return err
		}
		_, err = file.WriteString(info + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
