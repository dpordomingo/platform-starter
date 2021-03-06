// Code generated by go-bindata.
// sources:
// config/.csscomb.json
// config/.editorconfig
// config/.eslintrc.js
// hooks/pre-commit
// DO NOT EDIT!

package main

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configCsscombJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\x4d\x76\xdb\x30\x0c\x84\xf7\x3e\x05\x1f\xd7\xc5\x05\x7c\x96\x6e\x40\x6a\xe4\xf0\x05\x02\x54\x10\x8e\x9b\xf4\xf5\xee\x7d\x72\xaa\xb6\x52\x65\x6f\x39\x1f\x66\x88\x9f\x1f\xa7\x94\x52\xca\x8e\xc9\xde\x40\x98\xe6\x78\x27\xbf\x0a\x3a\xa2\xe7\x73\x0a\xbf\xe2\xcb\x27\xc2\x72\xe3\xf7\x4e\x1d\x53\xab\x26\xa6\x5b\x75\x79\x72\xaa\xdc\x91\xcf\x29\x8b\xdd\xe0\xf9\xb7\x54\xc4\xea\x2b\x35\x1d\xa0\xb1\x88\xf7\xc7\x4d\x59\x7f\x31\x8f\x17\xd6\x21\x9f\xd3\xc8\xd2\x57\x53\x08\x26\x68\x1c\xdb\xc2\x46\x52\xdc\xa4\x29\xb6\x5f\x11\xf0\xd0\xf4\x42\x1f\x70\xdb\x19\x7e\xbb\x5a\x60\xe9\x2b\xf7\xa6\x17\xc1\xea\xd5\xcd\x83\xcc\x07\x38\x8d\x2c\x52\xb8\xbe\x2e\x10\x97\xfa\x87\x98\xb9\x82\x0a\x46\x73\xd0\xda\x7f\xde\xaa\x3c\x06\xfc\xaf\x98\x1e\xd4\x4e\xa5\x29\x87\xf9\x01\xb3\x3a\x3c\x41\x0a\xe2\x06\x28\x0d\xa8\xc2\xce\xd1\x4c\xef\x0d\x7d\xd5\xc3\x38\x9b\xa1\xcb\x30\x8a\x73\xc5\xc3\xc4\xff\xa8\xbd\xdb\x27\xd6\x21\xa8\x61\x4e\x03\xa4\x4d\x2d\xe0\x4f\x92\x8f\xe1\xe3\x99\x88\xf5\x47\xf1\xe1\x6d\xa6\x3b\xbd\xbb\xc7\xe0\x42\xbd\x7d\xec\x96\x7f\xd5\x16\x82\xde\xd7\xed\xff\x23\xbd\x41\x07\x73\x9a\x1d\x63\xfb\x4e\x2c\xed\xa2\xeb\x79\x9c\x7e\x9e\x7e\x05\x00\x00\xff\xff\x3a\x6d\xab\x3d\x09\x03\x00\x00")

func configCsscombJsonBytes() ([]byte, error) {
	return bindataRead(
		_configCsscombJson,
		"config/.csscomb.json",
	)
}

func configCsscombJson() (*asset, error) {
	bytes, err := configCsscombJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/.csscomb.json", size: 777, mode: os.FileMode(436), modTime: time.Unix(1498730933, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configEditorconfig = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\x31\x4f\x33\x31\x0c\x86\xf7\xfc\x0a\x4b\xdd\x3e\xe5\xeb\x80\x18\x58\x3a\x41\x07\x06\x7e\x41\x75\x8a\xd2\x8b\x73\xe7\x92\x73\xaa\xd8\xa5\x94\xaa\xff\x1d\x25\x1c\x20\x55\xc0\x66\xbd\x79\xf3\xf8\x91\x17\xb0\x0e\xa4\xb9\xdc\x67\x8e\x34\xc0\x88\x69\x2f\x10\xf0\x05\x53\xde\x63\xa9\x63\x24\x46\xf0\x1c\x60\xf2\xc4\xea\x89\xa1\xcf\x2c\x24\x8a\xac\x66\x01\x7d\x0e\xc4\x03\x88\x9e\x12\x0a\x6c\x51\x8f\x88\x0c\x81\x62\xc4\x82\xac\x80\x0d\x2f\x8d\xf0\xf8\xb0\x16\xb3\x98\xa3\xbe\x6d\x5c\xe6\x32\x18\x53\x72\x56\x58\x81\x96\x03\x1a\xb3\xf9\xd7\x19\xe4\xe0\x72\x74\xa9\xee\x5e\x41\x8a\xa6\x1f\x7d\x11\xac\xa5\x83\xc6\xff\x77\x46\x0b\x4d\x4e\x8b\xa7\x44\x3c\xb8\xe3\x48\x8a\xb2\xf7\x3d\x7e\x52\x88\x05\x8b\xba\x48\xec\x93\x63\x3c\xce\xa4\xf9\x2d\x20\xab\x6b\xca\xb0\x82\xf6\xef\x2b\xa4\xb7\x9a\xdd\x56\x8d\xe5\xb8\x95\xee\x37\x52\xf4\x49\x9a\xec\xf2\xbc\x13\xab\x62\x77\xf2\x6a\x4f\x53\xb2\x3b\xc9\x7c\xe9\xae\x78\x37\x1f\xcd\x7a\x16\x3b\x85\x4b\xf7\x97\xff\x37\x79\xc8\xdd\xb5\xac\xfa\xad\x31\x9b\xf3\x93\x7f\xc6\x48\x09\x2d\x4c\xf3\x74\xf9\xb1\xfa\x1e\x00\x00\xff\xff\xda\x97\x97\x56\xdf\x01\x00\x00")

func configEditorconfigBytes() ([]byte, error) {
	return bindataRead(
		_configEditorconfig,
		"config/.editorconfig",
	)
}

func configEditorconfig() (*asset, error) {
	bytes, err := configEditorconfigBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/.editorconfig", size: 479, mode: os.FileMode(436), modTime: time.Unix(1498730953, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configEslintrcJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\xc1\x6a\xf3\x30\x10\x84\xef\x7e\x8a\x41\xe7\xf8\xf7\x7f\x76\xe9\xa1\x97\x9e\xfa\x06\x25\x14\xd9\x9a\x24\x02\x45\x2b\x76\xa5\x36\x10\xfc\xee\xc5\x26\xa4\x2d\x39\xec\x69\xe6\xfb\xd8\xdd\xb3\x84\x96\xf8\x8f\x97\x22\x5a\x0d\xcf\xb8\x76\x00\x2f\x95\x39\xd8\x08\xe7\xa3\x4e\x79\xea\x27\x6f\x74\xbb\x0e\x28\xa9\x1d\x63\xb6\x11\xef\x2e\x9e\x57\xc4\xed\xe0\x8a\xb2\xd6\x48\x75\xfb\xb5\xa2\x2d\xd1\xc6\xcd\x83\x9f\x6c\xb8\x97\x46\x38\xaa\x8a\x6e\x3e\xe0\xe6\x19\xb2\xf4\xbc\x54\xf5\x99\xd2\xac\x0f\x2c\xcc\x81\x79\x8e\x34\x37\xe2\xff\x43\xb7\x65\xa5\x49\xfa\x64\x78\x8c\xb7\xf5\x2d\x4a\xfe\x8d\x1e\x5a\x9e\xfb\xec\xcf\x7f\x7c\x9b\x28\x50\x6d\x16\x65\x1f\x7c\x3e\x26\x6e\x31\x86\x01\x13\x67\xdf\x8c\x90\x03\x3e\xac\x15\xea\x9d\x29\xa9\xd9\x3a\x6e\xfd\xc3\xed\x18\x5c\xe1\x53\x92\xaf\x57\xd1\x37\x91\xf2\x72\xa8\xd4\x7a\x92\x76\x3c\x55\x1b\x51\xb5\x11\xcb\xbe\x03\x96\x6e\x79\xea\xbe\x03\x00\x00\xff\xff\x8c\xcd\x00\x2e\x76\x01\x00\x00")

func configEslintrcJsBytes() ([]byte, error) {
	return bindataRead(
		_configEslintrcJs,
		"config/.eslintrc.js",
	)
}

func configEslintrcJs() (*asset, error) {
	bytes, err := configEslintrcJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/.eslintrc.js", size: 374, mode: os.FileMode(436), modTime: time.Unix(1498733572, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _hooksPreCommit = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x96\xdf\x6f\xdb\x36\x10\xc7\xdf\xf9\x57\xdc\x54\x01\x8e\x83\x52\xda\xf6\xd8\x36\x1b\x3a\xff\xc8\x0c\x24\xd1\x50\x1b\xdd\x43\x57\x04\x94\x44\x49\xe7\x51\xa4\x40\x52\x4e\x5d\xec\x8f\x1f\x44\x49\x76\xec\xd8\x71\x1d\x20\x4f\x16\x78\x3f\x78\xf7\xbd\x0f\x0f\x7e\xf3\x53\x18\xa3\x0c\x4d\x41\xc8\x6d\x34\x9e\x4d\x67\x93\xf1\xfd\x74\x76\x33\x99\x5f\xf9\x17\x39\x5a\x48\x31\xcb\x80\x52\xc9\x4a\x4e\x95\x14\x6b\x68\x8f\x35\x5f\xd1\x8a\x69\xc3\xe1\xcf\xc9\xc7\xf1\x10\xfe\x83\x5c\xf3\x0a\xe8\x0a\x06\x81\x59\xe5\x83\x9d\x83\x4a\xee\x1d\x2c\xab\xbd\x03\x23\x59\x35\x18\x12\xf2\x06\x46\x05\x4f\xfe\x85\xc4\x98\x44\x95\x31\x21\xa3\xf9\x7c\x14\xdd\xfe\x71\xe5\x5f\x3c\x14\x98\x14\xbd\x01\x7e\xfd\x0d\xc2\x94\xaf\x42\x59\x0b\x31\x24\x98\xc1\x97\x2f\x40\xbf\x83\xe7\x77\x01\x1e\x7c\xfd\xfa\x1e\x6c\xc1\x25\x01\x00\x48\x98\x85\x0f\x1f\xfe\x99\x44\x53\x32\xd1\x5a\xe9\x77\x9b\x44\x68\x40\x2a\x0b\x28\x8d\x65\x42\xf0\x34\x20\xe4\x6f\xb4\x05\xac\x99\x96\xef\x88\x0b\x6e\x3e\x21\x17\x2a\x66\x02\x58\x9a\x6e\x8b\x73\x8e\x77\x7f\xdd\x76\x7e\xb2\x2a\xfb\x3c\x40\xf3\xad\xdb\xc7\x58\x69\x8b\x32\x87\x44\x95\x25\xda\x80\x34\x65\x34\x01\xfc\x1b\x5a\xf8\x85\x64\xd8\x34\x6e\x0b\xae\x39\x94\x98\x17\x16\x62\x0e\x65\x2d\x2c\x56\x82\x83\xa9\xe3\x14\x35\x4f\xac\xd2\xc8\x0d\x3c\x34\x77\x6e\x8c\x7d\x17\x89\x92\x19\xe6\xb5\x66\x16\x95\x34\x24\x53\xba\x37\xdd\xb7\x26\x40\x09\xfe\x45\x86\x32\x85\x00\xdc\x34\x61\x10\x74\x2e\xc1\xd2\x28\x39\x18\xbe\x87\x54\xb5\x62\x75\x49\x29\xed\x62\x3d\x7f\x37\x99\x07\x94\x0a\x94\x16\xfc\x5d\x64\x5c\x74\x3b\x0b\xff\x77\xa0\x92\xc3\xcf\xbb\x63\x70\x4d\x27\x85\x02\xcf\x4d\x01\x32\xa5\x4b\x66\x9d\x36\xb5\x71\x0a\xb5\xf7\xbc\x05\xb6\xa7\x99\xb7\x8d\x6f\x45\x6b\x3e\x33\x24\xa9\x92\x7c\x8b\x0d\x4f\xd1\x2a\xdd\x16\x49\xc8\x64\x3c\x5b\x44\x9f\x46\xd1\xdd\x74\x76\x7d\xbf\x88\xa2\x9b\xf9\x06\xa3\xc7\x8e\xd4\x2a\x25\xcc\x71\xa2\x9e\xa6\x39\x09\xd7\x81\xf4\x2f\xe0\xec\x69\x96\x53\xc8\x1d\x8a\xf8\x11\xfa\x1a\x5e\x32\x14\xdc\x51\xb2\x3b\xd3\x0d\x15\x07\x7a\x4a\x9c\xe8\x9e\xdf\x84\x7a\xe4\x87\xa6\xff\x48\xab\x85\x03\x9e\x39\xe8\x4d\x47\x02\x4f\x5d\x1d\xe6\x2d\x54\x82\x33\xc3\x21\xc3\x6f\x4d\x82\x32\x80\x4f\xbc\xe4\x65\xcc\x35\xac\x55\x0d\x09\x93\x4d\x5d\xb6\x40\x03\xf1\x1a\x96\xb5\xb1\x1d\x41\xb6\xe0\x90\x29\x21\xd4\x43\xdf\x33\x93\x69\xa7\xd6\x91\x36\x9a\x3b\x2a\x66\x8b\xd0\xaa\xb0\x5c\x87\xae\x80\xf0\xf2\x32\xbc\x7c\x46\xbc\x63\x24\x4e\xe6\x37\xb3\xbb\xc5\x96\x33\xe3\xde\xc9\x71\xb6\x9c\xfb\x69\x9e\xda\x34\x2f\x61\xa8\x8d\x6c\x7f\xba\xf7\x4c\x19\xea\x58\xc6\x34\x6e\x04\xee\x2c\x95\xa8\x73\x94\x14\xcb\x4a\x69\xbb\x77\x58\x69\x6e\x2d\x72\x0d\xfd\xc7\x49\x0c\x5f\xe7\xd2\x57\xdc\xa3\x5d\xc5\x07\xd6\x68\x6b\x79\x6e\x8b\xb6\x1e\x3a\x09\x96\xe6\xd1\x12\xed\x32\xd2\x04\x3c\x7f\x27\x87\x07\x9e\x7f\x91\xa2\x76\xd1\x94\x3e\x31\x0f\x1d\x7a\xc1\xd2\x78\xe7\xae\x53\x03\x99\xaa\x65\x0a\x2c\x49\x94\x4e\xdd\x63\x50\x07\x3b\x3b\x7f\xbd\x46\x95\xc5\x12\xbf\x73\x98\x7f\xbe\x06\x2c\x59\xde\x2b\x68\x56\xb9\x22\xf3\xcf\xd7\xd1\x06\xf9\xe6\xe4\x38\xf0\x8d\xeb\x33\xb8\x9f\x58\x4d\x2f\xf8\x47\x32\xb8\x74\xff\x47\xb6\x83\x71\xf5\x51\x8a\xb2\xaa\xed\x55\xbf\xbb\xce\x56\x5a\xb5\x82\x34\x12\x9a\x55\xbe\xd5\xe2\x4c\x69\xff\x0f\x00\x00\xff\xff\xbc\x93\x20\x92\x7e\x09\x00\x00")

func hooksPreCommitBytes() ([]byte, error) {
	return bindataRead(
		_hooksPreCommit,
		"hooks/pre-commit",
	)
}

func hooksPreCommit() (*asset, error) {
	bytes, err := hooksPreCommitBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "hooks/pre-commit", size: 2430, mode: os.FileMode(509), modTime: time.Unix(1498740326, 0)}
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
	"config/.csscomb.json": configCsscombJson,
	"config/.editorconfig": configEditorconfig,
	"config/.eslintrc.js": configEslintrcJs,
	"hooks/pre-commit": hooksPreCommit,
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
	"config": &bintree{nil, map[string]*bintree{
		".csscomb.json": &bintree{configCsscombJson, map[string]*bintree{}},
		".editorconfig": &bintree{configEditorconfig, map[string]*bintree{}},
		".eslintrc.js": &bintree{configEslintrcJs, map[string]*bintree{}},
	}},
	"hooks": &bintree{nil, map[string]*bintree{
		"pre-commit": &bintree{hooksPreCommit, map[string]*bintree{}},
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

