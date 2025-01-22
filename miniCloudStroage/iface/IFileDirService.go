package iface

type IFileDirService interface {
	listDirs(root string) []string

	addDir(curPath string) error

	delDir(delPath string) error
}
