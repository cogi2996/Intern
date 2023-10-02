package main

import (
	"errors"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"

	"bytes"
	"mime/multipart"
	"path"
)

// https://codevoweb.com/how-to-upload-single-and-multiple-files-in-golang/
// https://chenyitian.gitbooks.io/gin-web-framework/content/docs/12.html

/*
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/appleboy/test.zip" \
  -H "Content-Type: multipart/form-data"

curl -X POST http://localhost:8080/upload \
  -F "upload[]=@/Users/appleboy/test1.zip" \
  -F "upload[]=@/Users/appleboy/test2.zip" \
  -H "Content-Type: multipart/form-data"

curl -X POST http://localhost:8080/upload/single -H "Content-Type: multipart/form-data" -F "image=@/Users/lap00004/Downloads/questions.jpeg"

curl -X POST http://localhost:8080/upload/multiple -H "Content-Type: multipart/form-data" -F "images=@/Users/lap00004/Downloads/panel_2.png" -F "images=@/Users/lap00004/Downloads/panel_3.png"
*/

const HOST = "localhost:8080"

func uploadResizeSingleFile(ctx *gin.Context) {

	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.UnixNano()) + fileExt
	filePath := "http://" + HOST + "/images/single/" + filename

	imageFile, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	src := imaging.Resize(imageFile, 1000, 0, imaging.Lanczos)
	err = imaging.Save(src, fmt.Sprintf("public/single/%v", filename))
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	ctx.JSON(http.StatusOK, gin.H{"filepath": filePath})
}

func uploadSingleFile(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	filePath := "http://" + HOST + "/images/single/" + filename

	out, err := os.Create("public/single/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"filepath": filePath})
}
func uploadResizeMultipleFile(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["images"]
	filePaths := []string{}
	for _, file := range files {
		fmt.Printf("UploadResizeMultipleFile: Filename %s ext %s size %d \n", filepath.Base(file.Filename), filepath.Ext(file.Filename), file.Size)

		fileExt := filepath.Ext(file.Filename)
		originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
		now := time.Now()
		filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
		filePath := "http://" + HOST + "/images/multiple/" + filename

		filePaths = append(filePaths, filePath)
		readerFile, _ := file.Open()
		imageFile, _, err := image.Decode(readerFile)
		if err != nil {
			log.Fatal(err)
		}
		src := imaging.Resize(imageFile, 1000, 0, imaging.Lanczos)
		err = imaging.Save(src, fmt.Sprintf("public/multiple/%v", filename))
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"filepaths": filePaths})
}
func uploadMultipleFile(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["images"]
	filePaths := []string{}
	for _, file := range files {
		fileExt := filepath.Ext(file.Filename)
		originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
		now := time.Now()
		filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
		filePath := "http://" + HOST + "/images/multiple/" + filename

		filePaths = append(filePaths, filePath)
		out, err := os.Create("./public/multiple/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		readerFile, _ := file.Open()
		_, err = io.Copy(out, readerFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"filepath": filePaths})
}
func SetupFolders() {
	if _, err := os.Stat("public/single"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("public/single", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	if _, err := os.Stat("public/multiple"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("public/multiple", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func fileUpload() {
	fileDir, _ := os.Getwd()
	fileName := "upload-file.txt"
	filePath := path.Join(fileDir, fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	r, _ := http.NewRequest("POST", "http://example.com", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	client.Do(r)
}
