package main

import (
	"encoding/json"
	"encoding/xml"
	"exporter/exporter"
	"io/ioutil"
	"strconv"
)

func main() {
	fb := new(exporter.Facebook)
	twtr := new(exporter.Twitter)
	ldin := new(exporter.Linkedin)
	//err := export(twtr, "twtrdata.txt")
	//err := export(fb, "fbdata.txt")

	//if err != nil {
	//	panic(err)
	//}

	//ScrollFeeds(fb, twtr, ldin)

	feedf := make(map[string]string)
	for i := range fb.Feed() {
		feedf[strconv.Itoa(i+1)] = fb.Feed()[i]
	}
	f, _ := json.MarshalIndent(feedf, "", " ")
	_ = ioutil.WriteFile("facebook.json", f, 0644)
	fcb, _ := xml.MarshalIndent(fb.Feed(), "", " ")
	_ = ioutil.WriteFile("facebook.xml", fcb, 0644)

	feedt := make(map[string]string)
	for i := range twtr.Feed() {
		feedt[strconv.Itoa(i+1)] = twtr.Feed()[i]
	}
	t, _ := json.MarshalIndent(feedt, "", " ")
	_ = ioutil.WriteFile("twitter.json", t, 0644)
	tw, _ := xml.MarshalIndent(twtr, "", " ")
	_ = ioutil.WriteFile("twitter.xml", tw, 0644)

	feedl := make(map[string]string)
	for i := range ldin.Feed() {
		feedl[strconv.Itoa(i+1)] = ldin.Feed()[i]
	}
	in, _ := json.MarshalIndent(feedl, "", " ")
	_ = ioutil.WriteFile("linkedin.json", in, 0644)
	lin, _ := xml.MarshalIndent(ldin, "", " ")
	_ = ioutil.WriteFile("linkedin.xml", lin, 0644)

}

// ScrollFeeds prints all social media feeds
/*func ScrollFeeds(platforms ...exporter.SocialMedia) {
	for _, sm := range platforms {
		for _, fd := range sm.Feed() {
			fmt.Println(fd)
		}
		fmt.Println("=================================")
	}
}

func export(u exporter.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}

	for _, fd := range u.Feed() {
		n, err := f.Write([]byte(fd + "\n"))
		if err != nil {
			return errors.New("an error occured writing to file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}*/
