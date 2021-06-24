package utils

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"strings"
)

//LicenseToSha256 return license sha256
func LicenseToSha256(value string) string {
	h := sha256.New()
	_, err := h.Write([]byte(value))
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", (h.Sum(nil)))
}

//ReadPomComments read license from pom comments
func ReadPomComments(pomData string) string {
	reader := strings.NewReader(pomData)
	closer := ioutil.NopCloser(reader)
	scanner := bufio.NewScanner(closer)
	var buffer bytes.Buffer
	var line string
	var startBuffer bool
	var stopBuffer bool
	var comment string
	for scanner.Scan() {
		line = scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "<!--") {
			startBuffer = true
			line = strings.Replace(line, "<!--", "", -1)
		}
		if strings.HasPrefix(strings.TrimSpace(line), "-->") {
			stopBuffer = true
			line = strings.Replace(line, "-->", "", -1)
		}
		if startBuffer && !stopBuffer {
			if len(strings.TrimSpace(line)) > 0 {
				buffer.WriteString(line)
			}
		}
		if stopBuffer {
			comment = buffer.String()
			break
		}
	}
	err := closer.Close()
	if err != nil {
		return ""
	}
	return comment
}
