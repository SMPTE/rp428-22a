package main

/*

Written by Jack Watts for SMPTE Standards TC 27C Community.

*/

import (
	"flag"
	"fmt"
	"os"

	"github.com/SMPTE/rp428-22a/pkg/tt"
)

func main() {
	run()
}

// run available command line flags.
func run() {

	flag.BoolVar(&tt.Txt, "text", true, "- Inidcate that text profile is to be used.")
	flag.BoolVar(&tt.Img, "image", false, "- Inidcate that image profile is to be used.")
	flag.BoolVar(&tt.Track, "T", false, "- write MXF trackfile, requires '-d'")
	flag.BoolVar(&tt.Encrypt, "e", false, "- encrypt trackfile")
	flag.IntVar(&tt.Duration, "d", 24, "- set the duration of the track file.")
	flag.StringVar(&tt.Framerate, "p", "24", "- set the frame rate of the track file.")
	flag.IntVar(&tt.Display, "m", 0, "- set the DisplayType.'0'=MainSubtitle,'1'=ClosedCaption. (default '0')")
	flag.IntVar(&tt.Reel, "r", 1, "- set the ReelNumber, Default ='1'")
	flag.StringVar(&tt.Language, "l", "en", "- set the RFC 5646 Language subtag")
	flag.StringVar(&tt.Title, "t", "No Title", "- set the ContentTitleText value.")
	flag.StringVar(&tt.Template, "x", "", "- path to 428-7 XML to use as template")
	flag.StringVar(&tt.Output, "o", "", "- set the output path, Default is StdOut")
	flag.Parse()
	if len(flag.Args()) > 0 {
		fmt.Println("check command expression")
		os.Exit(1)
	}
	if err := tt.CreateXML(tt.Txt, tt.Img, tt.Track, tt.Encrypt, tt.Reel, tt.Display, tt.Duration, tt.Framerate, tt.Language, tt.Title, tt.Template, tt.Output); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}
