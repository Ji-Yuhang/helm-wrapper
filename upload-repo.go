package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func uploadRepoChart(c *gin.Context) {
	repoName := c.Request.PostFormValue("repoName")
	if repoName == "" {
		err := errors.New("repoName must be not empty")
		respErr(c, err)
		return
	}
	file, header, err := c.Request.FormFile("chart")
	if err != nil {
		respErr(c, err)
		return
	}
	filename := header.Filename
	t := strings.Split(filename, ".")
	if t[len(t)-1] != "tgz" {
		respErr(c, fmt.Errorf("chart file suffix must .tgz"))
		return
	}

	out, err := os.Create(helmConfig.UploadPath + "/" + filename)
	if err != nil {
		respErr(c, err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		respErr(c, err)
		return
	}

	// abspath, err := filepath.Abs(helmConfig.UploadPath + "/" + filename)
	absolutePath, err := filepath.Abs(out.Name())
	if err != nil {
		respErr(c, err)
		return
	}
	cmd := exec.Command("helm", "cm-push", absolutePath, repoName)
	fmt.Printf("absolutePath :\n%s\n", string(absolutePath))

	cmdOutByte, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(cmdOutByte))
		// log.Fatalf("cmd.Run() failed with %s\n", err)
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	cmdOut := string(cmdOutByte)
	fmt.Printf("cmd_out_str out:\n%s\n", cmdOut)

	if strings.Contains(cmdOut, "Error") {
		err_lines := strings.Split(cmdOut, "\n")
		need_response_err_lines := make([]string, 0)
		for _, line := range err_lines {
			if strings.Contains(line, "Error") || strings.Contains(line, "error") {
				fmt.Println("Find Error Line %s", line)
				need_response_err_lines = append(need_response_err_lines, line)
			}
		}
		respErr(c, errors.New(strings.Join(need_response_err_lines, "\n")))
	} else {
		respOK(c, cmdOut)
	}
}
