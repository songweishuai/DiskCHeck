package diskCheck

import (
	"DiskCheck/error"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	. "os/exec"
)

type DiskInfo struct {
	DevName    string
	DiskName   string
	DiskStatus string
}

func getDiskInfo() ([]DiskInfo, int, error) {
	diskNum := 0

	cmd := Command("sh", "-c", "mount | grep '/dev/sd' | awk '{print $1,$3}' | sort")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("stdout pipe err:", err)
		return nil, diskNum, err
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		fmt.Println("cmd.Start err:", err)
		return nil, diskNum, err
	}

	diskInfo := make([]DiskInfo, 0, 8)
	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var m DiskInfo
		n, err := fmt.Sscanf(line, "%s %s", &m.DevName, &m.DiskName)

		if err == nil && n > 0 {
			diskNum++
			diskInfo = append(diskInfo, m)
		}
	}

	return diskInfo, diskNum, nil
}

func GetDiskStatus(c *gin.Context) {
	diskInfo, diskNum, err := getDiskInfo()
	if err != nil {
		myError.ReturnErrorMsg(c, err)
	}

	for i, _ := range diskInfo {
		disk := "smartctl -H " + diskInfo[i].DevName + " | grep 'result: ' | awk -F ': ' '{print $2}'"

		cmd := Command("sh", "-c", disk)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("stdout pipe err:", err)
			diskInfo[i].DiskStatus = "cmd.StdoutPipe error"
			continue
		}

		if err := cmd.Start(); err != nil {
			fmt.Println("cmd.Start err:", err)
			diskInfo[i].DiskStatus = "cmd.Start error"
			continue
		}

		opBytes, err := ioutil.ReadAll(stdout)
		if err != nil {
			fmt.Println("ioutil.ReadAll err:", err)
			diskInfo[i].DiskStatus = "ioutil.ReadAll error"
			continue
		}

		diskInfo[i].DiskStatus = string(opBytes)

		stdout.Close()
	}

	data, err := json.Marshal(diskInfo)
	if err != nil {
		myError.ReturnErrorMsg(c, err)
	}

	c.JSON(200, gin.H{
		"status": "error",
		"data":   string(data),
		"num":    diskNum,
	})
}
