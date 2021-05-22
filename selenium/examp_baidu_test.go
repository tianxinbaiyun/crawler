package selenium

import (
	"log"
	"testing"
)

func Example() {
	uri := "http://www.baidu.com"
	job := NewJob(uri)
	defer job.Close()

	title, err := job.Driver.Title()
	if err != nil {
		log.Printf("Failed to get page title: %s", err)

	}
	log.Printf("Page title: %s\n", title)

	//elems, err := job.Driver.FindElements(selenium.ByCSSSelector, ".s-hotsearch-content .hotsearch-item")
	//if err != nil {
	//	log.Printf("Failed to find element: %s\n", err)
	//	return
	//}
	//
	//for _, elem := range elems {
	//	text, _ := elem.Text()
	//	log.Printf("text %s ", text)
	//	return
	//}

	cookie, _ := job.Driver.GetCookies()
	log.Println(cookie)
}

func TestExample(t *testing.T) {
	Example()
}

func BenchmarkExample(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Example()
		}
	})
	//BenchmarkExample-20    	       1	2403773979 ns/op
}
