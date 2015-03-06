package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("just got a request")

		cmd := exec.Command("./update.sh")

		cmdOutput := &bytes.Buffer{}
		cmd.Stdout = cmdOutput

		err := cmd.Run()
		printError(err)
		printOutput(cmdOutput.Bytes())
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
