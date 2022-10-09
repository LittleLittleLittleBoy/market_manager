package ocr

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/littlelittlelittleboy/market_manager/pkg/settings"
)

func GetImageContent(imageUrl string) string {
	valid := checkURL(imageUrl)
	if !valid {
		return ""
	}
	url := fmt.Sprint(settings.Config.Ocr.Endpoint, "vision/v3.2/read/analyze")
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"url":"%s"}`, imageUrl))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatalln(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", settings.Config.Ocr.Key)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	location := res.Header.Get("Operation-Location")
	if location == "" {
		body, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			fmt.Println(err)
			return ""
		}

		log.Fatalln(string(body))
		return ""
	}

	req, err = http.NewRequest("GET", location, nil)

	if err != nil {
		log.Fatalln(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", settings.Config.Ocr.Key)

	res, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return ""

}

func checkURL(url string) bool {
	if url == "" {
		log.Fatalln("image url is empty")
		return false
	}
	re := regexp.MustCompile(`(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`)
	result := re.FindAllStringSubmatchIndex(url, -1)
	if result == nil {
		log.Fatalln("image url check fail, url:", url)
		return false
	}

	return true
}
