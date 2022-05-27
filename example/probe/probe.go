package main

import (
	"fmt"

	fluentffmpeg "github.com/adefreitas/go-fluent-ffmpeg"
)

func main() {
	data, _ := fluentffmpeg.Probe("./video.avi")

	fmt.Println(data["format"])
}
