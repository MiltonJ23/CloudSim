package helper

type Stored_files struct {
}

type Connections struct {
	IP        string
	Namespace string
	NodeID    string
}

type Active_transfers struct {
	File       File
	SourceNode string
}

type File struct {
	ID   string
	Name string
	Size int64
}
