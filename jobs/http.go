package jobs

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Message output
type Message struct {
	Name       string
	Start      time.Time
	Gap        int64
	Body       []byte
	StatusCode int
}

// HTTPGet http get method
func HTTPGet(url string) (msg Message, err error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println(err)
	}
	var message Message
	// set name with url
	message.Name = url
	start := time.Now()
	message.Start = start.UTC().In(loc)
	resp, err := http.Get(url)
	end := time.Now()
	if err != nil {
		return message, err
	}
	// count time diff
	tcost := end.Sub(start).Milliseconds()
	message.Gap = tcost
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return message, err
	}

	message.Body = body
	message.StatusCode = resp.StatusCode
	return message, nil
}
