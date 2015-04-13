package main

import (
	"flag"
	"fmt"
	"github.com/eternnoir/flickrdownloader"
	"strconv"
)

func main() {
	urlPtr := flag.String("u", "", "Photo Set Url")
	pathPtr := flag.String("p", "/tmp", "Download path.")
	maxPagePtr := flag.Int("m", 99999, "Max Pages to Download")
	debugPtr := flag.Bool("d", false, "Enable debug mode.")
	sizePtr := flag.String("s", "o", "What image size you want to download.'o' is means origin.'l' means large.")
	flag.Parse()
	if len(*urlPtr) < 1 {
		fmt.Println("Url can not be empty.Use -h to get some help.")
		return
	}
	fmt.Println("Start download photo from:" + *urlPtr)
	fmt.Println("The max number of download page is " + strconv.Itoa(*maxPagePtr))
	fmt.Println("Download image size is " + *sizePtr)

	downloader := flickrdownloader.InitDownloader(*debugPtr)
	downloader.SaveAllPhoto(*urlPtr, *pathPtr, *maxPagePtr, *sizePtr)
}
