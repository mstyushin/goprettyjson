package main

import (
	"bufio"
	"os"
	"io"
	"log"
	"encoding/json"
	"bytes"
	"fmt"
	"regexp"
)

// 1KB of input should be pretty enough
const maxBytesInput int64 = 1024*1024

func main() {
	formattedJson := new(bytes.Buffer)
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, maxBytesInput)
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
		}
		if len(buf) != 0 {
			reg, err := regexp.Compile("(\\\\n+)")
			if err != nil {
				log.Fatal(err)
			}
			buf = []byte(reg.ReplaceAllString(string(buf), ""))
			if err := json.Indent(formattedJson, buf, "", "    "); err != nil {
				log.Println("Unexpected end of json")
				if int64(len(buf)) == maxBytesInput {
					log.Fatalf("Possibly too large input: >%d bytes", cap(buf))
				}
			}
		}
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	fmt.Println(formattedJson.String())
}
