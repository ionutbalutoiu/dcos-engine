// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// ../../../parts/windowsupgradescripts/winBootstrapUpgradeScript.ps1
// ../../../parts/windowsupgradescripts/winNodeUpgradeScript.ps1
package dcosupgrade

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

var _windowsupgradescriptsWinbootstrapupgradescriptPs1 = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\x6b\x6f\xdb\x3a\x12\xfd\xae\x5f\x31\xd0\x35\x6a\x1b\x0d\x1d\xa7\x0b\xec\x0d\xbc\xf0\x02\x69\x1e\x17\x59\x78\x93\xc0\x76\x9a\xc5\xd6\x85\xc1\x4a\x63\x85\x08\x45\xaa\x24\x15\xc7\x4d\xfd\xdf\x2f\xa8\x87\x45\xd9\x72\xd2\x97\x3f\x04\xb1\x34\x9c\x39\x73\x66\x38\x3c\xf4\x82\x71\x83\x0a\xa6\x2c\x46\x6d\x68\x9c\xc0\xb3\xff\xb1\xd5\xf9\x0b\x0d\x39\xa3\x06\x81\x5c\x48\x15\x53\x03\xb2\xfb\x09\x5a\x73\x7f\xed\x79\x8b\x54\x04\x86\x49\x01\x77\x8a\x19\x24\x23\x19\x75\x5a\x31\x6a\x4d\x23\xec\x7a\xcf\x1e\x00\x40\x2b\xd6\x11\x0c\xa1\x7c\x0c\xdf\x2a\xf7\xd9\xfb\x7c\xe5\x75\x6a\x92\xd4\x64\xc6\x9e\xeb\x77\x8c\x46\xad\x4e\x53\xc5\x3b\xad\x54\xf1\x03\x68\x25\xd4\xdc\x97\xae\x17\x52\x75\x5a\x0c\x86\x70\xf4\x2f\x68\x31\x20\x1c\xe1\xa8\x6f\xff\x7d\xfb\xb6\x0b\xb9\x89\xfd\x18\xb5\x72\xbe\xd9\xcf\x1b\x08\x52\xc5\x7b\xf8\x84\x40\xc8\x03\x62\x42\x39\x7b\x44\x62\x58\x8c\xf0\x0e\xc8\x62\xa4\x27\x40\x88\xb2\xb1\xe1\x5d\x1f\x88\xcc\xe3\x82\xc5\x50\x73\xc4\x16\xd0\x69\x8d\x4e\x26\xd3\xf3\xff\x5d\x4e\x4f\xaf\xcf\xce\x81\xe0\x17\xe8\x77\xb7\xe2\x55\x79\x8e\x64\x04\xfe\x99\x5c\x0a\x2e\x69\x88\x61\xe6\x11\x98\xb0\xf0\xa9\x31\x18\x27\x46\xfb\x3b\x4b\x15\x9a\x54\x89\xda\xe3\xf5\xe6\xdb\x1a\x02\x6a\x82\x7b\x27\x62\xf5\x6e\xc2\x11\x93\xce\xbb\xae\x57\x3d\x36\xf7\x4a\x2e\xc1\xbf\xa0\x8c\x63\x08\x46\x42\x58\x80\xc9\xa0\xf8\x35\xf2\x4f\x15\x52\x83\x67\x32\x78\x40\x35\x31\x54\x99\x4e\x6b\xc1\x38\x5e\xd1\x18\x0f\xa0\xc5\x65\x74\x00\xad\x47\xc9\xd3\xb8\x2a\x76\x20\x85\x41\x61\x60\x08\x7e\xb6\x82\x4c\x15\x15\x3a\x50\x2c\x31\x40\x72\x0e\xb9\x8c\x80\xd0\x24\x41\x11\xe6\xa9\x4e\xd0\x90\xd3\x62\x1d\xb9\xc9\x6c\xca\x30\x40\x3e\x50\x9e\xe2\xc6\xef\x26\x33\x27\x50\xbb\xd6\x41\x1d\xff\xe3\x73\x7f\xfd\x09\x9e\x8f\xd6\x3e\x90\x05\x34\x75\xef\x41\x01\x8e\x89\x08\xc2\x2c\x3b\xb0\xee\x28\x13\xa8\xfc\x6e\x7b\x13\xe4\x24\x0c\x7f\x05\x98\xff\xa6\xf0\x9e\x35\x9a\x4a\x85\xed\xa9\x18\x48\x08\x84\x08\x34\x4b\xa9\x1e\x20\x48\xb5\x91\xb1\xa0\x96\x1d\x38\xee\x1f\xff\x73\x70\xdc\x07\xf2\x58\xf2\x0a\x22\x62\xe2\x69\x70\x74\xdc\xff\x87\xff\xbb\x60\xb5\xbd\x86\xae\x15\x58\x75\xed\xcf\xf0\x59\xf5\x93\xcd\xb3\x20\x95\xc5\x34\x42\x3f\xef\xbe\x89\x91\x89\xd3\x0c\xd9\xb3\xf3\x27\x66\xe0\xc8\x5b\x7b\x3f\x55\xc1\x34\x08\x50\xeb\x45\xca\xf9\x0a\xb4\x2d\x27\x86\x0d\xd5\xf4\xb6\x03\x67\x41\xfb\x5e\x5e\xe5\x1f\xa1\x72\xed\x79\xf5\x49\xe2\xec\xe8\x4d\x3b\xa5\x49\xa4\x68\x88\x16\xc1\x82\x45\xa9\xa2\x76\x23\x55\xa5\x6b\xbd\x97\xd2\x68\xa3\x68\x72\x3b\x1e\xd9\x16\xb9\xbb\xbc\x9a\xbf\xbf\xbe\x9e\x4e\xa6\xe3\x93\x9b\xf9\xed\x78\xe4\xd8\x16\xbe\xce\x98\xb2\x96\xa7\x83\xd9\xc9\xd7\x54\xe1\x19\x35\x74\x56\xbc\x9a\x5d\x9d\xdf\xcd\x3f\x9c\x8f\x27\x97\xd7\x57\xce\xc2\x08\x85\x8d\x9f\x2f\xfc\x8f\x64\xa2\x48\xcd\x71\xe8\x17\x36\xce\x2a\x2e\xa3\xcc\x6c\xef\x92\x30\x90\x7a\x1e\xa1\x40\x45\x0d\xce\xf3\x0c\x7b\x5c\x46\xbb\x90\x6f\x15\x7f\xd9\x55\xf1\xff\x3c\x1b\x39\x8d\x84\xa2\x29\xf8\x84\x3b\x26\x42\xb9\xd4\xf0\xb9\xa4\x0e\x84\x0c\xd1\x0e\xff\x92\xed\x0a\xc0\x18\x63\xf9\x88\xe4\xd2\x60\x0c\x64\x8c\x41\xaa\x74\xde\x36\x01\x02\x39\x57\x4a\xaa\x93\x7c\xb4\x4d\x18\x47\x61\xf8\xca\x16\x9f\x09\x5b\xe6\x0a\xde\xc6\xdb\x15\x2e\x0b\x57\xf6\xef\x74\x95\x20\x9c\x31\x85\x81\x91\x6a\xb5\xf1\x9a\x27\x57\x51\x5e\x91\x91\xec\x50\xe0\x14\xc6\x2f\xf8\x5b\xd1\x98\x57\xf8\x83\x64\xab\xd0\xb9\x15\x59\x32\xd1\x73\x6a\x9d\xaf\xca\x23\xd4\xd6\x06\x83\x99\x3d\x44\x66\x45\xa0\x19\x4b\x48\x88\x06\x03\xd3\x4b\xf4\x91\xdf\x08\x33\x08\x6b\xb9\xbf\x04\xff\xd5\x66\x58\xe6\xa5\xea\x19\xaa\x7a\x4f\x5f\xdd\xb2\x14\xa7\xf8\xd6\x06\xc8\x13\xa8\xa0\xc4\x61\x36\x2a\x0f\x03\x68\x5b\x1a\x6e\x94\x8c\x14\x8d\xe1\x82\x71\xd4\xb3\x3f\xc9\xff\x59\x32\xfb\xf3\xab\x35\xf1\x01\xa1\x37\xfb\x0e\x0c\x40\xb4\x84\x6f\xf0\x9a\xb7\x27\x20\x9a\x01\x31\x86\xaa\x6a\xfc\xbf\x32\x26\xdd\xcf\xce\xa9\x9a\x0a\x43\x15\xfc\x10\x49\xeb\x8a\x88\x37\xfb\x92\x4b\xf4\x11\x10\x52\x3e\x25\x76\x23\x90\xa2\x2a\xa4\x38\x65\x4f\x6f\xc7\xe3\xb2\x4f\xe0\xdf\x9b\x6d\xfd\x5b\xb2\xca\x66\xfb\x1e\x5c\x8d\x89\xfc\x01\x17\x68\xc5\x49\x39\x16\x0b\x8c\xb7\xe3\x51\xd5\x67\x71\x26\x5f\x86\x30\x41\x8e\x81\x21\x13\xa3\xec\xc6\x2f\x5a\xae\x1c\x4a\xf6\xab\x41\x25\xc0\xbf\xb2\x9b\x7f\xd7\xdf\xc0\x07\x72\x4a\x35\x4e\x50\x68\x66\xd8\x23\xd6\x12\x26\x57\xd2\x14\x91\x5e\x48\xf4\xbf\x4c\x6b\x1b\x7b\x4f\x88\x4c\xa7\x15\x80\xfc\x06\xb5\x95\x69\xb9\x21\x74\xf2\x38\xbd\x11\x13\x08\x44\x61\xc2\x69\x80\xd0\xde\x87\xbb\x7d\xd0\x6e\x77\x7b\x53\xc5\xe2\x4e\xb7\x01\x73\xaa\xf8\x0b\x88\xdf\xd3\xf0\x47\xd1\x3a\xb5\xb1\x47\xe3\xce\x91\xe9\x0c\x00\x25\xed\x09\x0b\xc3\xd2\x26\xd1\x40\xbe\xd4\x1b\xa9\xb0\x79\x45\xf0\xda\x40\x89\x65\x36\x93\x33\xa0\x51\x3d\xb2\x00\x37\x11\x76\x15\x6f\x4d\x39\x3d\x30\xce\x37\xb6\x0d\xbc\x37\x9d\xc3\xb5\x40\xbe\x9b\xf4\xb8\xd2\x27\x9b\x9c\x61\xc9\xcc\x7d\xbe\xc6\x9d\x8b\xc1\x60\x96\x5b\xba\xeb\xa5\xe0\x2b\x08\x32\x65\xec\xa8\x37\xb6\x00\x66\x20\x94\xa8\x41\x48\x03\xf8\xc4\xb4\x23\xbe\x68\xc5\x61\x29\xfc\xb8\x86\x6f\xa0\xf3\x96\xd7\x45\xcb\x27\x65\x8f\x6f\xfc\xfa\x75\xb6\x69\x2f\x90\xa9\x95\x2a\xd9\x25\x63\xf3\x6e\x97\xfb\x1a\x7f\x1b\xad\x99\x83\x26\x24\x54\xec\x11\xd5\xd0\xb7\x01\x80\x10\x99\x18\x7b\x1c\xc5\xbd\x62\x4d\x61\x5f\x0e\x2a\x7d\xcf\xe2\x5e\xc8\x34\xfd\xcc\x71\x1e\x51\x83\x4b\xba\x0a\x85\x1e\x1a\x95\xa2\xdf\x88\xb5\x86\xf9\xfb\x46\x4d\xf9\xd9\x19\x39\x3b\x4c\xd7\x79\xdc\x0d\xba\x6e\x1e\xa8\x0e\x1f\x9f\x53\xc6\xf7\x68\x70\xe3\x08\x6e\xa7\xfc\x3f\x91\xd0\x4e\x22\x79\xd4\x9a\x32\x6e\x42\x5a\x0a\x7f\x3b\x47\xaa\xd3\xfa\xad\x7f\x68\x7b\x19\x0f\x07\xc1\xe0\x30\x83\x78\x78\x6f\x62\x3e\x50\xd2\xef\x36\x27\xf9\xab\xd7\x8d\xdf\x77\x5e\xbc\x96\xf1\xce\x2d\x33\x93\x31\xf9\xb2\x59\xf6\x20\x7f\x97\x4b\x98\x3d\xef\xac\x0c\x2d\xb3\xa8\x5c\x37\xdc\x2c\xb7\x14\x6a\x29\xf2\xed\xe8\x26\xe7\x22\x90\xa1\xdd\x89\x27\x3a\x60\xcc\xdb\x1e\xec\x7f\xbd\xee\x6d\xcb\x83\xcb\xe1\x8b\xe3\xdc\xd5\xad\x3b\x4e\x33\xa1\xf9\x3a\xd9\x1a\x8d\x95\xcb\x75\x99\xdc\x83\x1b\x8e\x54\x63\xf6\x23\x08\x8d\x28\x73\x2e\x23\x6b\x40\xae\xb1\x01\xcd\x1f\xf0\x80\x98\xd4\xef\x81\x04\x52\x8d\xa1\x3d\x54\x12\xaa\xec\x29\xb9\x67\xde\x17\xe6\xae\x80\xdf\x12\xee\x81\x8c\x13\x8e\x06\xc3\xde\xde\xa3\x2b\xff\x29\xa2\xea\x95\xfa\x6f\x1c\xce\xbc\x77\x04\x57\xe1\xa5\xf9\xb6\x30\x80\xd6\x3c\x77\x88\xf5\x3b\xe7\xf6\x65\x63\x1f\x56\xdf\xc3\xfc\xde\xf8\x77\x00\x00\x00\xff\xff\x63\xa2\x04\x54\x23\x13\x00\x00")

func windowsupgradescriptsWinbootstrapupgradescriptPs1Bytes() ([]byte, error) {
	return bindataRead(
		_windowsupgradescriptsWinbootstrapupgradescriptPs1,
		"windowsupgradescripts/winBootstrapUpgradeScript.ps1",
	)
}

func windowsupgradescriptsWinbootstrapupgradescriptPs1() (*asset, error) {
	bytes, err := windowsupgradescriptsWinbootstrapupgradescriptPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "windowsupgradescripts/winBootstrapUpgradeScript.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _windowsupgradescriptsWinnodeupgradescriptPs1 = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\xcd\x6e\xe3\x36\x10\xbe\xeb\x29\x06\x84\x50\xd8\xd8\x65\x90\xe4\xe8\x22\x07\xc3\xd2\x2e\x0c\x38\xb6\x21\x39\x71\x8b\xf5\x42\x60\xa5\xb1\x42\x94\x22\x55\x72\x14\xaf\xbb\xf5\xbb\x17\x94\x7f\x24\x67\xd3\x53\xd1\xea\x60\x98\x33\xdf\x7c\xf3\x3f\x5b\xa9\x08\x2d\xac\x64\x85\x8e\x44\x55\xc3\x77\xf6\x25\x1c\x7c\x46\xe2\x91\x20\x04\xfe\xc9\xd8\x4a\x10\x98\xe1\x57\x08\x33\x76\x08\x82\x6d\xa3\x73\x92\x46\xc3\xda\x4a\x42\x3e\x33\xe5\x20\xac\xd0\x39\x51\xe2\x30\xf8\x1e\x00\x00\x84\x95\x2b\xe1\x01\xce\x62\xf8\xab\xa3\x6f\xf5\x47\xcb\x45\x43\x75\x43\x2d\x38\xe8\xf3\x26\x48\x76\x3f\x69\xac\x1a\x84\x8d\x55\x1f\x21\xac\x05\xbd\x9c\xa9\xb7\xc6\x0e\x42\x09\x0f\x70\xf7\x33\x84\x12\xb8\x42\xb8\xbb\xf5\x7f\x3f\x7c\x18\xc2\x11\xe2\x3f\xb2\xfb\xde\xcb\x7f\x3f\x41\xde\x58\x75\x83\xdf\x10\x38\xff\x1d\xb1\x16\x4a\xbe\x22\x27\x59\x21\xdc\x03\xdf\xce\x5c\x0a\x9c\x5b\xef\x1b\xee\x6f\x81\x9b\xa3\x5f\xf0\x31\x5c\x11\xc9\x2d\x0c\xc2\xd9\x38\x5d\xc5\xbf\x4c\x57\x93\x45\x14\x03\xc7\x3f\xe0\x76\xf8\xc6\x5f\x97\xe7\xcc\x94\xc0\x22\xb3\xd3\xca\x88\x02\x8b\x96\x11\xa4\xf6\xe1\x0b\x22\xac\x6a\x72\xec\x07\x53\x8b\xd4\x58\x7d\x25\x3e\x5c\x5e\x07\xc8\x05\xe5\x2f\x3d\x8f\x9d\x2e\x55\x88\xf5\xe0\x7e\x18\x74\x62\x7a\xb1\x66\x07\xec\x93\x90\x0a\x0b\x20\x03\xc5\x29\x98\x36\x14\xe6\x8b\x1f\x36\x75\x69\x45\x81\x69\x6e\x65\x4d\x4f\xc9\x0c\x1e\x80\xad\xa7\xf3\xec\x69\xf9\x39\x19\x47\x71\x96\x4e\x92\xe9\x72\x95\x3d\x25\x33\x76\x01\x47\xd2\x7a\xd8\x64\xb4\x19\xff\xd9\x58\x8c\x04\x89\xcd\x49\xb5\x99\xc7\xeb\xec\x39\x4e\xd2\xe9\x62\xce\x82\x50\x99\xf2\x9f\x90\x59\x0f\x79\xa3\x4c\xc9\x82\x50\x14\x95\xd4\x4f\x0e\x5b\xf6\x71\xf4\xe8\xc3\x48\xe3\x84\x05\x61\x2d\x9c\xdb\x19\x5b\x74\x8a\xe5\x38\x4d\xd7\x8b\x24\x62\xc1\x75\xcb\x53\x12\x96\xf8\xca\x0a\xed\xda\x94\x80\x2f\xdb\x6e\xfa\x48\xb8\xa8\x6b\xd4\x45\xf0\x4e\x9b\x5a\x33\xa9\x4b\xd0\xa6\x40\x38\x85\xe8\x4b\x16\x4d\x16\x29\x5c\x25\x75\xb6\x4e\xb0\x32\xaf\xc8\xa7\x84\x15\xf0\x04\xf3\xc6\xba\xe3\xda\xe4\x08\x3c\xb6\xd6\xd8\xf1\x71\xae\x53\xa9\x50\x93\xda\x4f\x8c\x26\xa9\x1b\x84\x5e\x1d\x2f\x6c\x73\xdc\x9d\xa8\xfc\xef\x6a\x5f\x23\x44\xd2\x62\x4e\xc6\xee\x2f\xac\xc7\x54\xde\xb1\xce\x8b\x2b\xf1\x45\xfe\x25\xd6\xaf\xd2\x1a\x5d\xa1\xa6\xaf\xa3\x51\x8a\xd4\x13\x3c\x0b\x2b\xc5\x6f\x0a\x07\x2c\xfd\x35\x5d\xc5\x8f\x51\x96\xc6\xc9\xf3\x74\x12\xb7\x55\x9f\x8f\x1f\x63\xf6\x11\x58\x88\xfa\x75\x94\x9b\xaa\x6e\x08\xad\x16\x15\x6e\x36\x5d\xa3\x3c\xe0\x51\xe4\x2f\x52\x23\x1b\xfe\x5b\xaf\x97\x96\xb6\xab\x7f\x6c\xf8\x15\xff\xff\x9a\xd6\xd2\x9a\x1c\x9d\xfb\xcf\xd2\xea\xf8\x7b\x13\x75\x3a\x80\xf0\xe3\x5a\xb2\x22\x37\x2e\xf3\xd3\x99\x9d\x74\x37\xb5\xbb\x63\x9d\xf1\xcd\xe6\x5d\x44\x70\x78\x7b\x32\x7a\x53\xdf\x5d\x86\xf3\xc4\xaf\xa5\x2e\xcc\xce\x81\x28\x51\x53\xbb\x0c\x23\x7f\xfe\x7b\xeb\x65\xea\xde\x76\xb5\x72\xfc\x26\x09\xee\x82\x43\xd0\xdf\xa7\x26\xf7\xd9\x6d\x1b\xa5\xf6\x67\xf2\xe2\x1d\x76\x16\xbc\x65\x6c\xd9\x6e\x83\xbf\x03\x00\x00\xff\xff\x01\xe3\xb4\xbb\x9f\x06\x00\x00")

func windowsupgradescriptsWinnodeupgradescriptPs1Bytes() ([]byte, error) {
	return bindataRead(
		_windowsupgradescriptsWinnodeupgradescriptPs1,
		"windowsupgradescripts/winNodeUpgradeScript.ps1",
	)
}

func windowsupgradescriptsWinnodeupgradescriptPs1() (*asset, error) {
	bytes, err := windowsupgradescriptsWinnodeupgradescriptPs1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "windowsupgradescripts/winNodeUpgradeScript.ps1", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"windowsupgradescripts/winBootstrapUpgradeScript.ps1": windowsupgradescriptsWinbootstrapupgradescriptPs1,
	"windowsupgradescripts/winNodeUpgradeScript.ps1":      windowsupgradescriptsWinnodeupgradescriptPs1,
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
	"windowsupgradescripts": {nil, map[string]*bintree{
		"winBootstrapUpgradeScript.ps1": {windowsupgradescriptsWinbootstrapupgradescriptPs1, map[string]*bintree{}},
		"winNodeUpgradeScript.ps1":      {windowsupgradescriptsWinnodeupgradescriptPs1, map[string]*bintree{}},
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
