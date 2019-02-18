package main

import (
	"bytes"
	"time"

	astisub "github.com/asticode/go-astisub"
)

func main() {
	s1, _ := astisub.OpenFile("/Users/xliu/Documents/51CTO_Golang/learnGo/src/srtFileDemo/example.ttml")
	s2, _ := astisub.ReadFromSRT(bytes.NewReader([]byte("00:01:00.000 --> 00:02:00.000\nCredits")))

	s1.Add(-2 * time.Second)

	s1.Merge(s2)

	s1.Optimize()

	s1.Unfragment()

	s1.Write("/Users/xliu/Documents/51CTO_Golang/learnGo/src/srtFileDemo/example.srt")
	var buf = &bytes.Buffer{}
	s2.WriteToTTML(buf)
}
