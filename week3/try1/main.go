package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mu sync.RWMutex
var num = 0

type school struct {
	School_id int    `json:"school_id"`
	Name      string `json:"name"`
}
type response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Item []struct {
			Admissions      string `json:"admissions"`
			Answerurl       string `json:"answerurl"`
			Belong          string `json:"belong"`
			Central         string `json:"central"`
			CityID          string `json:"city_id"`
			CityName        string `json:"city_name"`
			CodeEnroll      string `json:"code_enroll"`
			CollegesLevel   string `json:"colleges_level"`
			CountyID        string `json:"county_id"`
			CountyName      string `json:"county_name"`
			Department      string `json:"department"`
			Doublehigh      string `json:"doublehigh"`
			DualClass       string `json:"dual_class"`
			DualClassName   string `json:"dual_class_name"`
			F211            int    `json:"f211"`
			F985            int    `json:"f985"`
			Hightitle       string `json:"hightitle"`
			InnerRate       int    `json:"inner_rate"`
			IsRecruitment   string `json:"is_recruitment"`
			IsTop           int    `json:"is_top,omitempty"`
			Level           string `json:"level"`
			LevelName       string `json:"level_name"`
			Name            string `json:"name"`
			Nature          string `json:"nature"`
			NatureName      string `json:"nature_name"`
			OuterRate       int    `json:"outer_rate"`
			ProvinceID      string `json:"province_id"`
			ProvinceName    string `json:"province_name"`
			Rank            string `json:"rank"`
			RankType        string `json:"rank_type"`
			Rate            int    `json:"rate"`
			SchoolID        int    `json:"school_id"`
			SchoolType      string `json:"school_type"`
			TagName         string `json:"tag_name"`
			Type            string `json:"type"`
			TypeName        string `json:"type_name"`
			ViewMonth       string `json:"view_month"`
			ViewTotal       string `json:"view_total"`
			ViewTotalNumber string `json:"view_total_number"`
			ViewWeek        string `json:"view_week"`
		} `json:"item"`
		NumFound int `json:"numFound"`
	} `json:"data"`
	Location  string `json:"location"`
	Encrydata string `json:"encrydata"`
}

var total = make([][]school, 10)

func Note(resbody []byte, page []school, i int) {
	var dataslice response
	json.Unmarshal(resbody, &dataslice)
	//if err != nil {
	//		fmt.Println("Error unmarshalling")
	//}
	for _, item := range dataslice.Data.Item {

		id := item.SchoolID
		name := item.Name
		var s = school{
			School_id: id,
			Name:      name,
		}

		mu.Lock()
		page = append(page, s)
		num++
		fmt.Printf("当前记录了第%d个大学(%s)\n", num, name)
		mu.Unlock()

	}

	total[i-1] = page

	mu.RLock()
	fmt.Printf("第%d页的所有学校已全部录入\n", i)
	mu.RUnlock()

	wg.Done()
}

func main() {
	//url := "https://api.zjzw.cn/web/api/?keyword=&page=1&province_id=&ranktype=&request_type=1&size=20&top_school_id=[2461,436]&type=&uri=apidata/api/gkv3/school/lists&signsafe=9326e2339790781062a5aac6ac933f66"

	for i := 1; i <= 10; i++ {

		page := make([]school, 0)

		URL := fmt.Sprintf("https://api.zjzw.cn/web/api/?keyword=&page=%d&province_id=&ranktype=&request_type=1&size=20&top_school_id=[2461,436]&type=&uri=apidata/api/gkv3/school/lists&signsafe=9326e2339790781062a5aac6ac933f66", i)
		response, _ := http.Get(URL)
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(body))
		wg.Add(1)
		go Note(body, page, i)
	}
	wg.Wait()
	fmt.Println("已全部录入")

	jsondata, _ := json.MarshalIndent(total, "", "   ")

	// 将 JSON 数据写入文件
	if err := ioutil.WriteFile("schools.json", jsondata, 0644); err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}

	fmt.Println("JSON 数据已成功写入文件: schools.json")
}
