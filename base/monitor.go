package base

import (
	"bufio"
	"os"
	"io"
	"strings"
)

func check(e error) {
	if e != nil || e == io.EOF {
		panic(e)
	}
}

func Environment() string{
	path := "/etc/go.env"
	_,err := os.Stat(path)
	if err != nil {
		return "dev"
	}
	f,err := os.Open(path)
	if err != nil {
		Logger.Warning("%s 打不开",path)
		return "dev"
	}
	defer f.Close()
	b :=  bufio.NewReader(f)
	line,err := b.ReadString('\n')
	for ; err == nil; line,err = b.ReadString('\n'){
		line = strings.Fields(line)[0]
		return line
	}
	return "dev"
}

func Lines() map[string]string{
	m := make(map[string]string)
	var data *os.File
	var err error
	if env == "prod"{
		//这里是生产，所以得需要注意
		data,err = os.Open("conf/mysql_servers.txt")
	}else{
		//这里是测试，所以可以改
		data,err = os.Open("conf/1.log")
	}
	check(err)
	defer data.Close()
	rd := bufio.NewReader(data)
	for {
		line,err := rd.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF || line == ""{
			break
		}
		temp := strings.Fields(line)
		var ip string = temp[0]
		var dbname string = temp[2]
		m[dbname] = ip
	}
	return m
}
