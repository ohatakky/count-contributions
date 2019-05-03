package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	resp, err := http.Get(fmt.Sprintf("https://github.com/users/ohatakky/contributions"))
	if err != nil {
		log.Fatal("http error", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("read error", err)
	}

	s := string(body)
	rep := regexp.MustCompile(`data-count.+/>`)
	res := rep.FindAllStringSubmatch(s, -1)

	// t := time.Now()
	// const layout = "2006-01-02"
	// var week [7]time.Time
	// for i := 6; i >= 0; i-- {
	// 	week[i] = t.AddDate(0, 0, -i)
	// }

	// var index_week_ago [7]int
	// for i, v := range res {
	// 	for ii, vv := range week {
	// 		if strings.Index(v[0], vv.Format(layout)) != -1 {
	// 			index_week_ago[ii] = i
	// 		}
	// 	}
	// }

	// 369 368 367 366 365 364 363

	var contributions [7]string
	index_today := 369
	c := 0
	for i, _ := range contributions {
		contributions[i] = res[index_today][0][12:14]
		if contributions[i][1] == '"' {
			contributions[i] = string(contributions[i][0])
		}
		contribution, _ := strconv.Atoi(contributions[i])
		c += contribution
		index_today--
	}

	fmt.Println(c)
}
