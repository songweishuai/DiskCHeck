package main

import (
	"DiskCheck/http"
	"fmt"
	"os"
)

func main() {
	//config.LoadServerConfig()

	/*create http web*/
	err := http.CreateHttpWeb()
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}
