package iface

type IFileService interface {
	Upload(data []byte, name string, path string)

	Download(path string) ([]byte, error)
}
