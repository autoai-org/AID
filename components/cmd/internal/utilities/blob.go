package utilities

import (
	"embed"
)

type Asset struct {
	Data     string
	Filepath string
	Ready    bool
}

func (a *Asset) Read() {
	var file embed.FS
	data, err := file.ReadFile(a.Filepath)
	if err != nil {
		Formatter.Error("Cannot load file " + a.Filepath)
	}
	a.Ready = true
	a.Data = string(data)
}
