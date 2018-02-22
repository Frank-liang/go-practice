package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CleanUrl(uri *url.URL, link string) string {
	switch {
	case strings.HasPrefix(link, "https") || strings.HasPrefix(link, "http"):
		return link
	case strings.HasPrefix(link, "//"):
		return uri.Scheme + ":" + link
	case strings.HasPrefix(link, "/"):
		return fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, link)
	default:
		p := strings.SplitAfter(uri.Path, "/")
		path := strings.Join(p[:2], "") //一般情况是这样 ,/static/img/logo.png
		return fmt.Sprintf("%s://%s%s%s", uri.Scheme, uri.Host, path, link)

	}
}
func cleanUrls(u string, urls []string) []string {
	var ret []string
	uri, _ := url.Parse(u)
	for i := range urls {
		ret = append(ret, CleanUrl(uri, urls[i]))
	}
	return ret
}

func fetch(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("src")
		if ok {
			urls = append(urls, link)
		}
	})

	return urls, nil
}

func downloadImgs(urls []string, dir string) error {
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			continue
			return errors.New(resp.Status)
		}

		fullname := filepath.Join(dir, path.Base(u))
		f, err := os.Create(fullname)
		if err != nil {
			return err
		}
		defer f.Close()
		io.Copy(f, resp.Body)
	}
	return nil
}

func maketar(dir string, w io.Writer) error {
	basedir := filepath.Base(dir)
	compress := gzip.NewWriter(w)
	defer compress.Close()
	tr := tar.NewWriter(compress)
	defer tr.Close()
	filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		// 写入tar的FileHeader
		// 以读取的方式打开文件
		// 判断目录和文件,如果是文件
		// 把文件的内容写入到body
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		p, _ := filepath.Rel(dir, name)
		//fmt.Printf("dir:%s, name:%s, p:%s\n", dir, name, p)
		header.Name = filepath.Join(basedir, p)
		tr.WriteHeader(header)
		if info.IsDir() {
			return nil
		}

		f, err := os.Open(name)
		if err != nil {
			return err
		}
		defer f.Close()
		io.Copy(tr, f)
		return nil
	})
	return nil
}

func fetchImages(w io.Writer, url string) {
	urls, err := fetch(url)
	if err != nil {
		log.Panic(err)
	}

	urls = cleanUrls(url, urls)
	tmpdir, err := ioutil.TempDir("", "spider")
	if err != nil {
		log.Panic(err)
	}
	//fmt.Println(tmpdir)
	//defer os.RemoveAll(tmpdir)
	err = downloadImgs(urls, tmpdir)
	if err != nil {
		log.Panic(err)
	}

	maketar(tmpdir, w)
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fetchImages(w, r.FormValue("u"))
}

func main() {
	http.HandleFunc("/", handleHTTP)
	http.ListenAndServe(":7071", nil)
}
