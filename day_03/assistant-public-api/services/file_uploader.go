package services

import (
	"image"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IFileUploader interface {
	// UploadFile save file and return originalFileName and file path
	UploadFile(ctx *gin.Context, fieldFile string) (string, string, error)
	// UploadImage save file and return originalFileName, file path and thumbnail
	UploadImage(ctx *gin.Context, fieldImage string) (string, string, string, error)
}

type FileUploader struct {
}

func NewFileUploader() IFileUploader {
	return &FileUploader{}
}

func (*FileUploader) UploadFile(ctx *gin.Context, fieldFile string) (string, string, error) {
	file, header, err := ctx.Request.FormFile(fieldFile)
	if err != nil {
		st, _ := status.New(codes.InvalidArgument, "Field image invalid.").WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field:       fieldFile,
				Description: "Field file is invalid.",
			},
			&errdetails.LocalizedMessage{
				Locale:  "vi",
				Message: "Field file is invalid",
			},
		)
		return "", "", st.Err()
	}

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + uuid.NewString() + fileExt
	filePath := makeLocalFile(filename)

	out, err := os.Create(filePath)
	if err != nil {
		st, _ := status.New(codes.Internal, err.Error()).WithDetails(
			&errdetails.ErrorInfo{
				Reason: "CREATE_FILE_FAILED",
				Domain: "",
			},
			&errdetails.LocalizedMessage{
				Locale:  "vi",
				Message: "Không thể lưu file.",
			},
		)
		return "", "", st.Err()
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		st, _ := status.New(codes.Internal, err.Error()).WithDetails(
			&errdetails.ErrorInfo{
				Reason: "SAVE_FILE_FAILED",
				Domain: "",
			},
			&errdetails.LocalizedMessage{
				Locale:  "vi",
				Message: "Không thể lưu file.",
			},
		)
		return "", "", st.Err()
	}

	return originalFileName + fileExt, filePath, nil
}

func (*FileUploader) UploadImage(ctx *gin.Context, fieldImage string) (string, string, string, error) {
	file, header, err := ctx.Request.FormFile(fieldImage)
	if err != nil {
		st, _ := status.New(codes.InvalidArgument, "Field image invalid.").WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field:       fieldImage,
				Description: "Field image invalid.",
			},
			&errdetails.LocalizedMessage{
				Locale:  "vi",
				Message: "Field image invalid",
			},
		)
		return "", "", "", st.Err()
	}

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + uuid.NewString() + fileExt
	filePath := makeLocalFile(filename)

	// full image

	out, err := os.Create(filePath)
	if err != nil {
		st, _ := status.New(codes.Internal, err.Error()).WithDetails(
			&errdetails.ErrorInfo{
				Reason: "CREATE_FILE_FAILED",
				Domain: "",
			},
			&errdetails.LocalizedMessage{
				Locale:  "vi",
				Message: "Không thể lưu file.",
			},
		)
		return "", "", "", st.Err()
	}

	_, err = io.Copy(out, file)
	if err != nil {
		out.Close()
		st, _ := status.New(codes.Internal, err.Error()).WithDetails(
			&errdetails.ErrorInfo{
				Reason: "SAVE_FILE_FAILED",
				Domain: "",
			},
			&errdetails.LocalizedMessage{
				Locale:  "vi",
				Message: "Không thể lưu file.",
			},
		)
		return "", "", "", st.Err()
	}
	out.Close()

	// thumbnail
	input, err := os.Open(filePath)
	if err != nil {
		st, _ := status.New(codes.Internal, err.Error()).WithDetails(
			&errdetails.ErrorInfo{
				Reason: "CREATE_FILE_FAILED",
				Domain: "",
			},
			&errdetails.LocalizedMessage{
				Locale:  "vi",
				Message: "Không thể lưu file.",
			},
		)
		return "", "", "", st.Err()
	}
	defer input.Close()

	imageFile, _, err := image.Decode(input)
	if err != nil {
		st, _ := status.New(codes.Internal, err.Error()).WithDetails(
			&errdetails.ErrorInfo{
				Reason: "DECODE_FILE_FAILED",
				Domain: "",
			},
			&errdetails.LocalizedMessage{
				Locale:  "vi",
				Message: "Không thể lưu file.",
			},
		)
		return "", "", "", st.Err()
	}
	src := imaging.Resize(imageFile, Config.FileStorage.ThumbnailWidth, 0, imaging.Lanczos)

	thumbnail := makeLocalFile(strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + uuid.NewString() + "-min" + fileExt)
	err = imaging.Save(src, thumbnail)
	if err != nil {
		st, _ := status.New(codes.Internal, err.Error()).WithDetails(
			&errdetails.ErrorInfo{
				Reason: "SAVE_FILE_FAILED",
				Domain: "",
			},
			&errdetails.LocalizedMessage{
				Locale:  "vi",
				Message: "Không thể lưu file.",
			},
		)
		return "", "", "", st.Err()
	}

	return originalFileName + fileExt, filePath, thumbnail, nil
}
