package main

import (
	"fmt"
	"time"

	"./mock"
	"./real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		from map[string]string) string
}

//接口组合
type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}
func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	r = &mock.Retriever{"this is a fake imooc.com"}
	mRetriever := mock.Retriever{"this is a fake imooc.com"}
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock Retriever")
	}

	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
	// fmt.Println(download(r))
	fmt.Println("Try a session")
	fmt.Println(session(&mRetriever))
	// fmt.Println(download(r))

}
func inspect(r Retriever) {
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
		fmt.Printf("%T %v\n", r, r)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
		fmt.Printf("%T %v\n", r, r)
	}
}
