package main

import (
	"flag"
	"fmt"
	"github.com/eternnoir/flickrdownloader"
)

func main() {
	urlPtr := flag.String("u", "", "Photo Set Url")
	pathPtr := flag.String("p", "/tmp", "Download path.")
	maxPagePtr := flag.Int("m", 99999, "Max Pages to Download")
	debugPtr := flag.Bool("d", false, "Enable debug mode.")
	flag.Parse()

	downloader := flickrdownloader.InitDownloader(*debugPtr)
	downloader.SaveAllPhoto(*urlPtr, *pathPtr, *maxPagePtr)
}
