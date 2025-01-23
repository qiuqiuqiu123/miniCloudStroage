package iface

type IFileDirService interface {
	ListDirs(root string) []string

	AddDir(curPath string) error

	DelDir(delPath string) error
}
