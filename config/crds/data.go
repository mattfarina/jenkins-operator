// Code generated by go-bindata. DO NOT EDIT. @generated
// sources:
// config/crds/jenkins_v1alpha1_jenkinsinstance.yaml
// config/crds/jenkins_v1alpha1_jenkinsjob.yaml
package crds

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

var _jenkins_v1alpha1_jenkinsinstanceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\xcd\x6e\xdb\x48\x0c\xbe\xeb\x29\x08\x5f\x83\x68\xd7\x9b\xcb\x42\xb7\x60\xf7\x92\x1e\x82\xa0\x09\x72\xa7\x47\x8c\xcc\x78\x34\x9c\xce\x70\xdc\xe4\xed\x0b\x49\xb6\x6b\x5b\x3f\x56\x82\xb6\xba\x08\x22\xe7\x23\x3f\xfe\x8d\x88\x9e\x9f\x29\x44\x16\x57\x00\x7a\xa6\x37\x25\xd7\x7c\xc5\x7c\xf3\x6f\xcc\x59\xfe\xda\x2e\x57\xa4\xb8\xcc\x36\xec\xca\x02\xfe\x4b\x51\xa5\xfe\x4a\x51\x52\x30\xf4\x3f\xbd\xb0\x63\x65\x71\x59\x4d\x8a\x25\x2a\x16\x19\x80\x09\x84\x8d\xf0\x89\x6b\x8a\x8a\xb5\x2f\xc0\x25\x6b\x33\x00\x8b\x2b\xb2\xb1\x39\x03\x60\xc4\x69\x10\x6b\x29\x5c\xab\x88\xdd\x3b\x2c\x60\xb1\xcc\xff\x5e\x64\x00\x0e\x6b\x2a\xe0\x95\xdc\x86\x5d\x64\x17\x15\x9d\xa1\x98\xef\x04\xfb\xb7\x78\x0a\xa8\x12\xf2\x1a\x9b\x37\x97\x79\xc5\xba\x4e\xab\xdc\x48\x9d\x45\x4f\xa6\xf1\x56\x05\x49\xfe\x60\x6b\x16\xb4\xf3\xbf\xe3\xda\x05\xff\xa5\x83\xdd\xed\xa8\xb4\x1a\x6f\x53\x40\xdb\xa7\x99\x01\x44\x23\x9e\x0a\xb8\x6f\xcc\x78\x34\x54\x66\x00\x5b\xb4\x5c\xb6\xc9\xe9\x0c\x8b\x27\x77\xfb\x70\xf7\x7c\xf3\x68\xd6\x54\x63\x27\x04\xf0\xa1\xe1\xa6\xbc\xf7\xdf\x3c\x47\x95\x3a\xc8\x00\xf4\xbd\xf1\x11\x35\xb0\xab\x0e\xe2\x96\xee\xa5\x43\xc7\x15\x3b\x3d\x28\xab\x57\x32\x7a\x10\xef\x93\xb8\x7f\x86\xc8\xb5\x04\xcb\x9a\x1d\xd5\xc8\xf6\x54\x3e\xe2\xff\x00\x89\x64\x02\xe9\x7c\x8c\x73\xa2\x6d\x0a\xe3\x30\xe6\x8c\xfe\xae\xd7\x5e\xb8\x9a\xed\x82\xdc\x76\xb6\x69\x7a\x23\x93\x54\x42\x8f\xcc\x8b\x84\x1a\xb5\x00\x76\x7a\xf3\xcf\xa0\x35\x76\x4a\x15\x85\x13\x1d\xd7\x58\xd1\xb9\x29\x8f\xaa\x14\x5c\x01\xf9\x55\x91\x5f\xcd\x8d\x82\x5d\x15\x28\xf6\x78\x8d\x95\x0f\xa6\x73\x3b\x99\x84\x1d\xc9\xf5\x38\x68\x80\x5f\xf3\x44\x0a\x5b\x36\xbd\x80\x2f\xe2\xd4\xc6\xe1\xb6\xb9\x80\x1c\x8d\xc0\x8a\x39\x1a\xcb\x19\xb6\x1c\xe9\x77\x09\x1b\x2f\x96\xcd\xfb\x30\x6a\x25\x62\x09\xdd\x89\xce\xdb\x54\x71\x3f\xbb\xac\x54\x0f\xa4\x7c\xaa\x56\x00\x5c\x0e\x49\x2f\x64\x0e\x60\xdb\xbf\x45\x66\x62\x27\xea\xdf\xa9\x30\x04\x7c\x3f\xd1\x84\x15\x9a\x8f\x74\xa0\xb1\x29\x2a\x85\x20\xf6\xa3\x4d\x31\x4a\x6e\xa4\xc7\x7e\xd7\x1c\xb4\xbf\xad\x4f\xce\x41\x7b\xe6\xd7\x86\x8d\xc6\x48\x72\xbd\x39\x99\x8c\x3e\xa9\xd4\x0d\xe8\xb1\x33\x71\xdb\x99\x78\x92\x0d\x0d\x76\xcd\x78\xb7\xc3\xfe\x46\x7b\x48\xd6\x3e\xb6\x03\x3b\xdc\xc9\xc3\xed\x3f\x19\xdf\xb1\xba\xdf\x77\xf0\xf9\x4a\xfc\x51\x96\xe3\xe5\x53\x09\x43\xbf\x82\x89\xba\xbd\xca\x2a\xfa\x6d\x6f\xdc\x2e\x86\xbc\xc3\x9d\xff\xe7\x67\xc4\x36\xa2\x1a\x5e\x23\x14\x35\xc5\xd9\x8b\xc4\x47\xb7\x02\xcf\xb3\xcf\xfa\x35\xc6\x5e\x62\x47\x4e\x07\xfa\x96\x38\xd0\xc9\x4d\x7b\xdd\x59\x98\x0e\xf8\x4c\xb4\xbf\x74\x61\xbb\x44\xeb\xd7\xb8\xcc\x7e\xe6\x03\x8d\x21\xaf\x54\xde\x9f\x2f\x9c\x8b\xc5\xc9\x8e\xd9\x7e\x1a\x71\x25\x77\x37\x53\xb7\x55\xff\x08\x00\x00\xff\xff\xb0\x02\xd8\x74\xc2\x0b\x00\x00")

func jenkins_v1alpha1_jenkinsinstanceYamlBytes() ([]byte, error) {
	return bindataRead(
		_jenkins_v1alpha1_jenkinsinstanceYaml,
		"jenkins_v1alpha1_jenkinsinstance.yaml",
	)
}

func jenkins_v1alpha1_jenkinsinstanceYaml() (*asset, error) {
	bytes, err := jenkins_v1alpha1_jenkinsinstanceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "jenkins_v1alpha1_jenkinsinstance.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jenkins_v1alpha1_jenkinsjobYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x4d\x8f\xd3\x30\x10\xbd\xe7\x57\x8c\x7a\x44\x6a\xa0\xe2\x82\x72\x43\x70\x61\x0f\xab\x15\xac\xf6\x82\x16\x69\xe2\x0c\xad\x5b\xc7\x63\x3c\x93\xd0\x45\xfd\xf1\x28\x71\x53\xb6\x69\xb6\x44\x9b\x43\x62\xbf\xf9\x78\x2f\x9e\xf1\x60\xb0\x0f\x14\xc5\xb2\x2f\x00\x83\xa5\xbd\x92\xef\x76\x92\xef\x3e\x48\x6e\xf9\x6d\xbb\x2a\x49\x71\x95\xed\xac\xaf\x0a\xf8\xd4\x88\x72\xfd\x95\x84\x9b\x68\xe8\x33\xfd\xb4\xde\xaa\x65\x9f\xd5\xa4\x58\xa1\x62\x91\x01\x98\x48\xd8\x81\xf7\xb6\x26\x51\xac\x43\x01\xbe\x71\x2e\x03\x70\x58\x92\x93\xce\x07\xc0\xb0\xd7\xc8\xce\x51\x5c\x2a\xb3\x1b\x08\x0b\x58\xac\xf2\x77\x8b\x0c\xc0\x63\x4d\x05\x6c\xc9\xef\xac\x97\x2d\x97\x92\x1f\xd7\xc3\x97\x03\x45\x54\x8e\x79\x8d\xdd\xd7\x56\xf9\xda\xea\xa6\x29\x73\xc3\x75\x26\x81\x4c\x47\xb4\x8e\xdc\x84\x53\x9a\x59\xa1\x89\xfa\x28\x33\xfd\xf7\x4d\x0a\xbb\xe1\xb2\x07\x83\x6b\x22\xba\x33\x71\x19\x80\x18\x0e\x54\xc0\x6d\x17\x1c\xd0\x50\x95\x01\xb4\xe8\x6c\xd5\x9f\x46\x4a\xc7\x81\xfc\xc7\xbb\x2f\x0f\xef\xbf\x99\x0d\xd5\x98\x40\x80\x10\x3b\x45\x6a\x07\xd6\xee\x79\x56\x9a\x13\x06\xa0\x4f\x1d\x87\x68\xb4\x7e\x7d\x82\x7b\x91\xff\x73\x7a\x5e\xa2\x73\x47\x2e\xb7\x64\xf4\x04\x0f\x47\x37\x3c\x53\xe2\xfa\x0a\x46\xaa\xc8\xab\x45\x37\x32\x00\x58\xa5\xfa\x02\x7c\x39\xd3\x38\xdf\x94\x15\x20\xa0\x2a\x45\x5f\xc0\x8f\xef\xb8\xfc\xf3\xd8\xbd\x96\x8f\x6f\xfa\xf5\xa4\xff\xe4\x29\x4c\xd1\xf5\x8e\xd7\x29\x1b\xa1\xd8\x75\xc5\x1d\x8a\xfc\xe6\x58\x1d\x84\x4c\x24\xbd\xa7\xbd\x1e\x84\x62\x6b\x0d\xa1\x31\xdc\x78\x3d\xb4\xd8\x38\x4d\xed\x94\xd6\x18\x42\x64\x47\x69\xa3\xbc\x23\xff\x1a\xb9\x89\x70\x5a\xe6\xac\xd0\x71\xed\xc7\xe1\xa3\x2e\x98\x61\x4c\x26\x8c\x11\x9f\xce\x2c\xc7\x6b\x61\xbd\x28\x7a\x73\x71\xb4\x2f\xaa\xdd\x72\x59\xc9\x45\xf1\xaf\xb9\xef\xeb\xb9\xee\xd3\x9d\xae\xa8\x8d\xcc\xe9\xf5\xb0\x41\x99\xfb\x23\x91\x7e\x35\x36\xd2\xd9\x8d\x5c\xa6\x0c\xd7\x05\x8d\xa0\x76\x18\xcc\xed\x0a\x5d\xd8\xe0\x2a\xfb\xa7\x17\x8d\xa1\xa0\x54\xdd\x8e\x27\xd5\x62\x71\x36\xa1\xfa\xad\x61\x5f\xf5\x63\x5a\x8e\x93\xf8\x6f\x00\x00\x00\xff\xff\x49\x39\x58\xde\xf6\x05\x00\x00")

func jenkins_v1alpha1_jenkinsjobYamlBytes() ([]byte, error) {
	return bindataRead(
		_jenkins_v1alpha1_jenkinsjobYaml,
		"jenkins_v1alpha1_jenkinsjob.yaml",
	)
}

func jenkins_v1alpha1_jenkinsjobYaml() (*asset, error) {
	bytes, err := jenkins_v1alpha1_jenkinsjobYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "jenkins_v1alpha1_jenkinsjob.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"jenkins_v1alpha1_jenkinsinstance.yaml": jenkins_v1alpha1_jenkinsinstanceYaml,
	"jenkins_v1alpha1_jenkinsjob.yaml": jenkins_v1alpha1_jenkinsjobYaml,
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
	"jenkins_v1alpha1_jenkinsinstance.yaml": &bintree{jenkins_v1alpha1_jenkinsinstanceYaml, map[string]*bintree{}},
	"jenkins_v1alpha1_jenkinsjob.yaml": &bintree{jenkins_v1alpha1_jenkinsjobYaml, map[string]*bintree{}},
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

