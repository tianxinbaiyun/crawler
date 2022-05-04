package chromedp

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	job *Job
)

func init() {
	job = NewJob(true)
	job.Debug = true
}

func TestStep_Do(t *testing.T) {
	s := &Step{
		URL:      "http://localhost:9033/admin/login",
		Selector: "",
		Action:   "Navigate",
		WantResp: false,
		Sleep:    5,
		Cron:     "",
	}
	resp, err := s.Do(job.Ctx)
	assert.NoError(t, err)
	job.Close()
	t.Log(resp)
}

func TestStep_Navigate(t *testing.T) {
	s := &Step{
		URL:      "http://localhost:9033/admin/login",
		Selector: "",
		Action:   "Navigate",
		WantResp: false,
		Sleep:    5,
		Cron:     "",
	}
	resp, err := s.Navigate(job.Ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	t.Log(resp)
	job.Close()
}

func TestStepLogin(t *testing.T) {
	defer job.Close()
	job.SetSteps([]*Step{
		{
			URL:      "http://localhost:9033/admin/login",
			Selector: "",
			Action:   "Navigate",
			WantResp: false,
			Sleep:    5,
			Cron:     "",
		},
		{
			URL:      "",
			Action:   "SendKeys",
			Selector: "#username",
			Value:    "admin",
			WantResp: false,
			Sleep:    5,
			Cron:     "",
		},
		{
			URL:      "",
			Action:   "SendKeys",
			Selector: "#password",
			Value:    "admin",
			WantResp: false,
			Sleep:    5,
			Cron:     "",
		},
		{
			URL:      "",
			Action:   "Click",
			Selector: "#sign-up-form > div:nth-child(4) > button",
			Value:    "",
			WantResp: false,
			Sleep:    5,
			Cron:     "",
		},
	})
	for _, step := range job.Steps {
		_, err := step.Do(job.Ctx)
		assert.NoError(t, err)
	}
	time.Sleep(time.Second * 20)

}
