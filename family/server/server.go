package server

import (
	"archive/tar"
	"compress/gzip"
	"github.com/chentianyou/family/family/common"
	"github.com/chentianyou/family/family/route"
	"github.com/go-errors/errors"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

var ImageData string

func init() {
	route.RegisterRoute(route.Route{
		Name:        "ImageList",
		Method:      []string{"GET"},
		Pattern:     "/images/list",
		HandlerFunc: ImagesList,
	})
	route.RegisterRoute(route.Route{
		Name:"GetImage",
		Method:[]string{"GET"},
		Pattern:"/images/{imgName}",
		HandlerFunc: GetImage,
	})
}

func ImagesList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		file,err := os.Open(ImageData)
		if err != nil {
			log.Println(errors.New(err).ErrorStack())
			common.HttpResponse(w, common.MSG_OTHER_ERROR("images file not exists"))
		}
		defer file.Close()
		archive, err := gzip.NewReader(file)
		tr := tar.NewReader(archive)
		var imageList []string
		for {
			hdr, err := tr.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(errors.New(err).ErrorStack())
				common.HttpResponse(w, common.MSG_OTHER_ERROR("failed to read image file"))
				return
			}
			if strings.Contains(hdr.Name, ".jpg") {
				imageList = append(imageList,path.Base(hdr.Name))
			}
		}
		common.HttpResponse(w, imageList)
	}
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		imgName := vars["imgName"]
		if imgName == "" {
			common.HttpResponse(w, common.MSG_OTHER_ERROR("image name cannot empty"))
			return
		}
		file,err := os.Open(ImageData)
		if err != nil {
			log.Println(errors.New(err).ErrorStack())
			common.HttpResponse(w, common.MSG_OTHER_ERROR("images file not exists"))
		}
		defer file.Close()
		archive, err := gzip.NewReader(file)
		tr := tar.NewReader(archive)
		for {
			hdr, err := tr.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(errors.New(err).ErrorStack())
				return
			}
			if strings.Contains(hdr.Name, imgName) {
				data, err := ioutil.ReadAll(tr)
				if err != nil {
					log.Println(errors.New(err).ErrorStack())
					return
				}
				w.Header().Set("Content-Type", "image/jpeg")
				w.Header().Set("Content-Length", strconv.Itoa(len(data)))
				if _, err := w.Write(data); err != nil {
					log.Println("unable to write image.")
				}
				break
			}
		}
	}
}