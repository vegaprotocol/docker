package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	passFilePath, ok := os.LookupEnv("PASSWORD_FILE_PATH")
	if !ok || passFilePath == "" {
		panic("missing PASSWORD_FILE_PATH env variable")
	}

	args := os.Args[1:]

	passFileBytes, err := ioutil.ReadFile(passFilePath)
	if err != nil {
		panic(err)
	}

	// always make sure that we can communicate over stdin
	args = append(args, "--stdio-ui")
	cmd := exec.Command("clef", args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdout.Close()

	go func() {
		stdOutScan := bufio.NewScanner(stdout)
		for stdOutScan.Scan() {
			text := stdOutScan.Text()

			fmt.Println("request:", text)

			var req struct {
				ID int `json:"id,omitempty"`
			}

			var rpcResp string
			if strings.Contains(text, "enter the password") ||
				strings.Contains(text, "enter a password") ||
				strings.Contains(text, "Please enter password") {
				if err := json.Unmarshal(stdOutScan.Bytes(), &req); err != nil {
					panic(err)
				}

				rpcResp = fmt.Sprintf(`{ "jsonrpc": "2.0", "id":%d, "result": { "text":"%s" } }`, req.ID, passFileBytes)
			}

			if strings.Contains(text, "ui_approveNewAccount") {
				if err := json.Unmarshal(stdOutScan.Bytes(), &req); err != nil {
					panic(err)
				}

				rpcResp = fmt.Sprintf(`{ "jsonrpc": "2.0", "id":%d, "result": { "approved":true } }`, req.ID)
			}

			if rpcResp != "" {
				fmt.Println("response:", rpcResp)
				_, err := stdin.Write(bytes.NewBufferString(rpcResp).Bytes())
				if err != nil {
					panic(err)
				}
			}
		}
	}()

	fmt.Println("running: clef", args)

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}
