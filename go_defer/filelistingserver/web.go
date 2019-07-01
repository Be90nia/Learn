package main

import (
	"code/learn/go_defer/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)
//函数式编程
//创建一个错误处理函数
type appHandler func(writer http.ResponseWriter, request *http.Request) error
//创建用户可以查看的报错信息接口
type userError interface {
	error
	Message() string
}
//统一的错误处理
func errWrapper(
	handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//panic error
		defer func() {
			//使用recover进行 panic 的错误信息进行处理
			if r := recover(); r != nil{
				log.Printf("Painc: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			log.Printf("Error handling request: %s", err.Error())
			//如果错误信息是userError类型的,则返回错误代码给用户查看
			//接口的类型会自动识别类型
			//user error
			if userErr ,ok := err.(userError); ok{
				http.Error(writer,userErr.Message(),http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
