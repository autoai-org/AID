// Code generated for package docker by go-bindata DO NOT EDIT. (@generated)
// sources:
// internal/assets/dockerfile.tpl
// internal/assets/runner.tpl
package docker

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _internalAssetsDockerfileTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x41\x8f\xda\x30\x10\x85\xef\xfe\x15\x23\x1f\xf6\x50\xe1\x98\xaa\x97\x2d\xd7\xdd\x56\x42\xd5\x92\x28\x80\x10\xa2\xa8\x72\x92\x21\xb5\x70\x1c\x77\x6c\xa7\xa5\x69\xfe\x7b\x95\x40\xe1\x50\xfb\x60\xcf\xf7\xfc\x9e\xdf\xe7\x3c\x7d\x03\x77\x09\xdf\x5b\xbb\xf8\x90\x7c\x14\xde\xe8\x46\x14\xd1\x07\x24\xc6\xf2\xed\x0a\x94\x0b\xa2\xc6\x00\xd1\x55\x2a\x20\x3c\x3d\xc1\x57\x06\x00\x77\xae\xad\x0f\xca\x18\x10\x97\x9b\x30\xae\xdf\x46\x17\xef\x6b\x51\x61\x07\x75\x59\x42\xd9\xa8\x33\x42\x11\xb5\xa9\x04\x7a\x8f\x36\x68\x65\xfe\x8f\x2a\x0d\x2a\xfb\xc0\xd4\x80\xa0\x13\xc8\x4e\x91\x34\xba\x90\xca\x05\x69\xb4\x0f\x5e\xbe\x03\x19\x1a\x37\x1e\xa3\x36\x5d\x19\x63\x2f\x69\xb6\x87\x44\x82\x54\xce\x5d\xbb\xf7\x7d\x46\x98\x2d\xb3\x3f\x5e\x9d\x70\x18\xae\xd0\x69\xf7\x28\x4d\xd3\x6b\x49\xf8\x23\x6a\xc2\x06\x6d\xf0\x49\xf8\x15\xfe\xd9\xd7\x18\xa2\xbb\xbb\x77\x69\xfe\xe5\x75\x99\xdf\x3e\xf8\xb4\xda\xe4\xfb\x2c\x5d\xae\x36\x70\xe0\x75\xb4\xba\x6c\xc9\xf2\x23\x63\x2f\x6f\xaf\x70\xe0\x14\xad\x45\xfa\xd6\xf7\xeb\xd6\x74\x48\x56\x35\x38\x0c\x0b\xa5\x2b\x8f\xd4\x21\xf1\x19\x17\x05\x9f\x01\x9f\x27\xd3\x5e\x3c\xcf\x9f\xe7\x23\x3c\xf3\x19\x8f\xdd\x14\x96\xfc\x6c\xe9\x8c\xe4\x93\xed\x75\xde\x4d\x23\x3f\xfe\x0d\x00\x00\xff\xff\x16\xc6\x33\xfa\xb4\x01\x00\x00")

func internalAssetsDockerfileTplBytes() ([]byte, error) {
	return bindataRead(
		_internalAssetsDockerfileTpl,
		"internal/assets/dockerfile.tpl",
	)
}

func internalAssetsDockerfileTpl() (*asset, error) {
	bytes, err := internalAssetsDockerfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/assets/dockerfile.tpl", size: 436, mode: os.FileMode(420), modTime: time.Unix(1606505983, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalAssetsRunnerTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xb1\x6a\xc4\x30\x0c\x86\x77\x3d\x85\x48\x97\x04\x5a\x93\x31\x1c\x64\x2a\x74\xee\x5e\x8a\x10\x89\x93\x33\x8d\x63\x23\xe5\x0a\x87\xf1\xbb\x97\xd4\xf8\x4a\xf1\xf2\xfb\xe7\x93\x3e\x3d\x4d\x61\x76\xfb\x7a\xb9\x1d\xcb\xcb\x00\xce\xc7\x20\x07\x06\xad\x89\x65\x8d\x2c\x6a\xeb\x5f\xef\x0a\x8b\x04\x8f\x7e\x8b\xde\xa8\x95\x6f\x2b\x58\x59\x37\x97\x02\x20\xa8\x99\xae\xb3\x93\x56\xef\x6a\x22\x1f\xd7\x8f\xfe\xb3\x83\x32\x99\xd2\x3b\x4f\x5f\xbc\xda\x9c\x4d\x4a\x6f\x6e\xb3\x3b\x7b\x9b\x73\x5d\x93\xd2\xeb\xc6\xaa\xa5\x04\xd0\xb0\x9d\x8e\xf1\x7f\xdf\x76\x00\x0f\x9f\x79\x30\x25\x00\xb8\x05\x89\x4e\x92\x68\x1c\x1b\x22\xcf\x6e\x27\x6a\x2e\x80\x88\x7f\x77\x1a\xb9\xed\x6d\xd3\x9b\xdf\xd7\x3c\xe3\xa9\x1f\x87\x7e\xe8\x3b\xf8\x09\x00\x00\xff\xff\x67\xcd\xb7\xba\x17\x01\x00\x00")

func internalAssetsRunnerTplBytes() ([]byte, error) {
	return bindataRead(
		_internalAssetsRunnerTpl,
		"internal/assets/runner.tpl",
	)
}

func internalAssetsRunnerTpl() (*asset, error) {
	bytes, err := internalAssetsRunnerTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/assets/runner.tpl", size: 279, mode: os.FileMode(420), modTime: time.Unix(1606506002, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"internal/assets/dockerfile.tpl": internalAssetsDockerfileTpl,
	"internal/assets/runner.tpl":     internalAssetsRunnerTpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"internal": &bintree{nil, map[string]*bintree{
		"assets": &bintree{nil, map[string]*bintree{
			"dockerfile.tpl": &bintree{internalAssetsDockerfileTpl, map[string]*bintree{}},
			"runner.tpl":     &bintree{internalAssetsRunnerTpl, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
