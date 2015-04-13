# Flickr Downloader

A Flickr download tool write by golang.It will download photos from flicker with origin size.

## Download

You can download binary file at [Release Page](https://github.com/eternnoir/flickrdownloader-bin/releases).

## How to use

```
$ flickrdownloader-bin -u <WhichFlickrUrlYouWantDownload> -m <MaxPagesToDownload> -p <SaveTo> -d <DebugMode>
```
* -u : what url you want to download.
* -m : how many pages you want to download.
* -p : what path you want to save photos.Default is ***/tmp***
* -s : image size. Default is "o" means origin size.
	* o : Original
	* l : Large
	* m : Medium 500
	* z : Medium 640
	* c : Medium 800
* -d : show debug infomation.Default is ***false***

### Examples

Download user marksein's photos only two pages.

```
flickrdownloader-bin -u https://www.flickr.com/photos/marksein/ -m 2
```

Download large size set's photo to /home/user/photo.

```
./flickrdownloader-bin -u https://www.flickr.com/photos/taisg/sets/72157649581313124 -p /home/frank/tmp -s l       
```
