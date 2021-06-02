package main

import (
	"log"

	"github.com/jviney/goav/avcodec"
	"github.com/jviney/goav/avdevice"
	"github.com/jviney/goav/avfilter"
	"github.com/jviney/goav/avformat"
	"github.com/jviney/goav/avutil"
	"github.com/jviney/goav/swresample"
	"github.com/jviney/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
