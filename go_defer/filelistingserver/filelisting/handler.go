package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix  = "/list/"
//创建用户可以查看的报错信息
type userError string
//实现其userError报错函数
func (e userError) Error() string {
	return e.Message()
}
//实现其userError的信息返回函数
func (e userError) Message() string  {
	return string(e)
}
func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	//检查URL是否与定义的URL路径一致
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start with " + prefix)
	}
	//获取文件路径
	path := request.URL.Path[len(prefix):]
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	//取得文件所有内容
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	//写入到http页面之上
	writer.Write(all)
	return nil
}