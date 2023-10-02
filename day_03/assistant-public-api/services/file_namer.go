package services

type IFileNamer interface {
	ToPublicFile(filePath string) string
	ToRelativeFile(filePath string) string
}

type FileNamer struct {
}

func NewFileNamer() IFileNamer {
	return &FileNamer{}
}

func (*FileNamer) ToPublicFile(relativeFile string) string {
	return Config.Domain + Config.FileStorage.URLPath + "/" + relativeFile
}

func (*FileNamer) ToRelativeFile(filePath string) string {
	return filePath[len(Config.FileStorage.Folder):]
}

func (*FileNamer) ToLocalFile(fileName string) string {
	return makeLocalFile(fileName)
}

func makeLocalFile(fileName string) string {
	return Config.FileStorage.Folder + fileName
}
