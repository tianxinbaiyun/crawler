package goqeury

import "testing"

func TestGetBaiduHotSearch(t *testing.T) {
	GetBaiduHotSearch()

	//2021/05/14 17:19:10 0:国家卫健委派出专家组前往安徽
	//2021/05/14 17:19:10 1:网红主播在酒店水壶内撒尿?
	//2021/05/14 17:19:10 2:以军方致电加沙居民称导弹将炸你家
	//2021/05/14 17:19:10 3:恒河出现大量浮尸 印媒给出原因
	//2021/05/14 17:19:10 4:安徽确诊病例曾2次停留北京
	//2021/05/14 17:19:10 5:救援队断水驴友却烧水泡茶
}

func BenchmarkGetBaiduHotSearch(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GetBaiduHotSearch()
		}
	})

	//BenchmarkGetBaiduHotSearch-20    	     176	   6528833 ns/op
}
