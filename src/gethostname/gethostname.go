package gethostname

import (
	"fmt"
	"os"
)

func GetHostname() string{
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println("本机名称：", name)
	return name

}
