# Flickr Downloader

An Flickr download tool write by golang.It will download photos from flicker with origin size.

## How to use

```
$ flickrdownloader-bin -u <WhichFlickrUrlYouWantDownload> -m <MaxPagesToDownload> -p <SaveTo> -d <DebugMode>
```
* -u : what url you want to download.
* -m : how many pages you want to downloader.
* -p : what path you want to save photos.Default is ***/tmp***
* -d : show debug infomation.Default is ***false***

### Example1

Download user marksein's photos only two pages.

```
flickrdownloader-bin -u https://www.flickr.com/photos/marksein/ -m 2
```

