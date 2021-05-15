package colly

import "testing"

func TestLogin(t *testing.T) {
	uri := "http://localhost:9033/admin/signin"
	job := NewJob(uri, nil, nil)

	res, err := job.Login(uri, map[string]string{"username": "admin", "password": "admin"})
	if err != nil {
		panic(err)
	}
	cookie := res.Headers.Get("Set-Cookie")
	t.Log(res)
	t.Log(cookie)
}
