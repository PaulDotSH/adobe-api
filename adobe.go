package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type AdobeResp struct {
	//TokenType  string `json:"token_type"`
	Token      string `json:"access_token"`
	ValidUntil int64  `json:"expires_in"`
}

func (a *AdobeResp) GetToken(m *sync.Mutex) {
	m.Lock()
	cmd := exec.Command("node", "./jwt-api/app.js")
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	if err := cmd.Run(); err != nil {
		fmt.Println(stdoutBuf.String(), stderrBuf.String(), err.Error())
	}
	//fmt.Println(stdoutBuf.String(), stderrBuf.String())
	a.Token = stdoutBuf.String()

	a.Token = removeEverythingBefore(a.Token, "access_token: '")
	a.Token = "Bearer " + strings.TrimSpace(removeEverythingAfter(a.Token, "',"))
	a.ValidUntil = time.Now().Unix() + 85500000
	m.Unlock()
}
