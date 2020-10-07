// Code generated for package util by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../resources/events.yaml
package util

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

var _ResourcesEventsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5d\xe9\x72\x1c\x37\x92\xfe\xef\xa7\x40\xf0\xc7\x8c\x1d\x41\x6a\xc5\x3e\x24\x92\xff\x28\x1e\x5a\x4e\x88\xa2\x86\x2d\xc9\xbb\xbf\x1c\xe8\x2a\x74\x37\xcc\x2a\xa0\x0c\xa0\x48\xb6\x27\xe6\x5d\xf8\x2c\x7c\xb2\x0d\x1c\x75\x74\x01\x55\x9d\x8d\xd6\x78\x1d\xe1\x08\x91\x44\xe1\xcb\xc4\x91\x48\xe4\x85\xa3\xa3\xa3\x9f\x3e\x72\x9c\x9d\xa1\xab\x47\xc2\x94\x44\xdf\x18\x5d\xd0\x04\x2b\xca\xd9\x4f\xdf\x89\x90\x94\xb3\x33\x74\xf0\xf8\xe6\xf4\xf8\xe0\xa7\x9c\xa7\x65\x46\xe4\xd9\x4f\x08\x1d\x21\x86\x73\x72\x86\x0e\x2e\xee\x6e\x6f\xef\x3e\x1f\xfc\x84\x10\x42\x09\x2f\x99\x3a\x43\xa7\xa7\xa7\xe6\x47\x9a\xce\x14\x16\xea\x0c\xbd\x75\x3f\x5e\xb1\xf4\x0c\xa1\xe6\xef\x6c\xc1\xcf\xcc\xbf\x74\x7f\x09\x4f\x49\xd5\x54\xff\x97\x91\x47\x92\x9d\xa1\x83\x9b\xcf\xd7\x77\x07\xf5\x6f\x73\x22\x25\x5e\x6a\xe0\x59\x99\x24\x44\xca\xe6\x4f\x85\xe0\xf3\x8c\xe4\x67\xe8\xa0\xf9\x9d\xe4\x59\xa9\x2c\x0b\x07\x1b\x54\x7f\xba\xd9\x20\xf9\xf8\xed\xdb\x4d\x92\x8f\xdf\xbe\xed\x50\x7d\xfc\x76\x80\xec\xba\x79\x8b\x72\x4d\xb8\x4f\x37\x9d\xf3\x05\x97\x88\x4a\x24\x35\x12\x49\x3d\xfa\x7d\xe2\x7d\xb0\xe3\x5d\xc1\x14\x11\x39\x65\x38\x16\x6f\x04\xc3\x4b\x32\x8a\x24\x11\x8f\x44\x68\x4c\xca\xa8\xa2\x38\xa3\x7f\x46\x82\x8e\x61\xa0\x8c\x3c\x69\x60\xc2\x94\x06\x4d\x38\x63\x24\x89\xe5\x73\x02\xe6\xd3\xc1\xa5\x54\xee\x87\x38\x85\x21\x0a\xf2\x47\x49\xa4\x42\xb9\x5c\x6a\x58\x41\x12\x42\x1f\x23\x21\xdf\x41\x21\x65\xc1\x99\x24\x15\xa6\x24\x4c\xc5\xe0\x1d\x03\x77\x46\x6b\xf1\x14\x82\x14\x58\x50\xb6\x44\xe4\x99\xc6\x81\x02\x77\x48\x0b\x54\xad\x04\xc1\x29\xfa\x9d\x53\x16\x37\xb0\xc7\xc0\x5d\x42\x18\x9e\x67\x04\x09\x52\x4a\x72\x84\xd3\x54\x44\x81\x79\xbb\xe3\xd7\xf3\xfb\xcf\x00\x30\xb4\xc0\x34\x8b\x64\xd0\xdb\x1e\x57\xf7\xf7\x77\xf7\x3e\xa8\xe4\xc9\x03\x51\x28\x11\xc4\x1c\x24\xfb\x40\x7a\xfb\x63\x18\x72\x4e\x59\xaa\xd7\xcd\x1e\x88\xde\xf6\x18\x46\xcc\xa8\x54\x64\x2f\x16\xdf\x03\x01\x49\xc1\xb3\xec\x87\x0c\xea\xc9\x6e\x2c\xe2\x24\x21\x85\xda\x07\xf0\x14\x08\x98\xe3\xe7\x4a\xb8\x12\x21\x78\xd4\xce\x18\x79\xf2\xa6\x0f\xcc\xfe\x43\x0b\xb6\x3d\xd7\xcc\xc8\x93\x36\x5b\x20\xad\xfc\xde\x13\xd4\x13\x37\x61\x09\xf0\xf5\xe6\xf6\xea\x12\xdd\x7d\xfb\x1a\x05\xe2\x89\x99\x1e\xce\x6e\x3e\x7f\x3f\xff\x74\x73\x89\xbe\x9c\xdf\x9f\xdf\xc6\x20\x8d\x81\xc7\xc4\x97\xbb\x19\xba\x99\xa1\x0f\xdf\x66\xff\x0b\x83\xa9\x95\xbe\x9b\x0f\xfc\xfa\x6e\x86\x66\x0a\x2b\x82\x6e\x31\xc3\x4b\x22\x36\xb4\xc0\x91\xa7\x05\x1e\x7b\x5a\xe0\x68\x48\x0b\x3c\xf6\x15\xb3\xcb\xab\x0f\xdf\x3e\x06\x76\x96\x21\x22\xe1\x4c\x91\x67\x85\x70\x9a\x46\xad\x81\x63\xa8\x62\xe6\xe0\x56\x98\x2d\xa3\x80\x46\x63\x4f\xf4\xdf\x5f\x7d\xb9\xbb\xff\xfa\xdb\xd7\xfb\xf3\x8b\xab\x80\xe6\xe9\xc6\x7a\x2d\x15\xc9\xd1\x3d\x49\xf8\x23\x11\x6b\x74\xc3\x0a\xc1\x97\x82\x48\xb9\xe3\xdc\x7d\xe7\x59\x99\x93\xd0\xa4\x8d\xbb\x93\x36\xf2\x54\xf7\xd1\xd0\xa4\x8d\xc0\xaa\xbb\xa5\xc1\x0a\xe0\x98\x51\x1c\x81\xf5\x76\x87\x94\x92\x8c\xc4\x22\x01\x17\x86\x43\xca\xf5\x58\x46\x22\x01\xd5\x74\x87\x54\xb2\x7d\xb0\x80\xfa\xb9\xc3\x32\xdb\x0a\x29\x8e\xd4\x8a\x98\x73\x3a\x0a\x13\xa8\xa1\x5b\xcc\xd7\x17\x41\x72\xfe\x48\x52\xb4\x10\x3c\xd7\xc0\xaf\x2f\xb1\xc8\xbe\xe2\xdc\x23\xe0\x57\xa4\xba\x1f\x90\x14\x3d\xba\xb5\xc3\x89\x44\x8c\x2b\xad\x41\x07\xe0\xcd\x47\xae\xe9\x13\x55\x2b\x33\x44\x5e\x27\x7a\xf3\x21\x2e\xaa\x1f\x6f\x2e\xfb\xba\x6d\xee\xd9\x57\x4c\x19\x5d\x5a\x0b\x37\x21\x48\xa2\xfa\xfb\xc2\x0b\xdd\x34\x59\x91\xe4\x41\x9f\x84\xaa\xa1\xa8\x35\x62\xad\xe1\xf0\x36\x4f\xcf\x70\x70\x8e\x72\xcc\xd6\xae\x33\x5f\xd2\xd4\x1b\x19\x33\xc3\xca\xbc\xde\xd3\x68\x4e\x12\x5c\x4a\x62\x68\xc9\xf1\x33\xcd\xcb\x1c\xb1\x32\x9f\x13\x81\xf8\xa2\xea\x10\xa9\x15\x56\xe6\xeb\xd6\x97\x54\x22\xf2\x9c\x10\xd2\x16\xe4\xcd\xa8\xdc\x13\x25\xd6\x8e\x61\xb3\x40\x34\xc3\xa5\xbe\x37\x6a\xaa\x45\x45\x2b\xc2\x39\xb7\x97\x1e\xa9\x74\x8b\xaa\xf3\x4d\x4e\x5a\x43\x02\x54\x01\xae\x0c\x65\xb6\x3b\x73\x1d\x70\x70\x92\xfe\x49\xc2\x4b\xc3\x5b\x0a\xba\xa9\xe3\x50\xa2\x73\x21\xf0\x1a\x25\xb8\xc0\x09\x55\xeb\x00\xbf\x17\x7a\x52\xcd\x28\x4a\x7b\x02\x54\x6d\x11\x66\x29\x32\x63\xb1\xc4\x94\x79\x0c\xf9\xba\x5b\x98\xa1\xef\xad\x35\x45\x25\x52\x9c\x23\xb9\xe2\x62\x78\x9d\x9b\xd6\x44\xaf\x4f\x3b\x5f\xaa\xfb\x51\x77\x15\xe3\x8d\x2f\xe7\x44\x3d\x11\xc2\xd0\xc8\xf0\x30\x9a\x4e\xf5\x79\x2a\x70\xa2\x88\xf0\x67\xc6\xd7\x08\xb7\x32\xf2\xfa\x52\xb1\x92\x71\xb6\xec\x5d\xb5\x3e\x17\x9d\x0f\x86\xb9\x30\x6b\xb7\xb5\x8b\xcd\xc2\xd8\xc6\x0c\x70\x99\x75\x66\x25\x2d\x8b\x8c\x26\xc1\xb3\x12\x9d\x6f\x08\x1f\xdc\xb4\xb5\x5f\xe3\x4c\xdf\xc5\xd7\x76\x27\xc8\x01\xd6\x52\xba\x58\x10\xa1\xef\x0b\x2d\x26\x7d\x06\x80\x97\xe5\x6f\xcc\x5e\x75\xda\x7b\xa4\xd5\xe1\x96\xc9\xd0\xda\x1c\xa6\x4c\x22\x8c\xa4\x12\x56\xa4\x61\x63\x25\xd2\x43\x8d\xb3\x8c\x3f\x05\x85\x43\x23\x32\xbd\x89\xca\x09\x51\xd2\xfb\x93\x28\xb3\x80\x30\xf0\x15\xe8\x30\x93\xf7\xf5\xd6\x36\x7b\xda\xac\x1f\x2c\x96\x3b\x09\x82\x6a\xdd\x6d\x7c\x17\x16\x77\x1b\xf2\xbd\xe0\x52\xd2\xb0\x00\x6a\x31\x02\xdc\x3b\x01\x46\x64\x8e\xb3\x6c\x77\x46\x5e\x5f\x36\x3f\x0c\x09\xb2\x9c\x32\x73\x1c\xe8\x79\x4c\x7c\x31\x6a\x84\x82\xd0\x7c\xfb\xfc\x00\xb7\x4f\x87\x1f\xbb\x68\xe8\x32\x64\x8d\x0a\x72\xf4\xfa\x62\xbe\xab\x37\xb7\xf9\xd8\x6e\xb0\x79\xc6\x93\x87\x4d\x71\xdf\xf0\x78\xc3\x8a\x52\x6d\xf0\xa2\xb8\x3e\xda\xf2\x32\x53\xb4\xc8\x88\x3e\xfa\xbc\x0e\x1a\xf6\x26\xbb\xc9\xec\x6a\x6b\xf7\xa9\x81\xe8\x5c\x29\x92\x17\x4a\x13\x61\xda\x34\x02\xac\xda\x4e\x7d\x5d\x34\x2c\x7d\xe6\x6a\x65\xd6\x1c\x47\x29\xf7\x29\xde\x4d\x38\x57\x70\xfd\xaa\x6b\x9b\x66\xd7\xaa\x9f\x6a\xbf\x1b\x28\xdd\x53\xff\xa2\xc2\x16\xdc\xa7\xfb\xc2\x29\x36\xf6\xfe\x50\x0d\x93\xa3\xa7\x97\x78\xab\x2a\xbb\x6f\x70\xdf\x57\x0d\xad\x2d\x9e\xdd\x47\x95\x12\x63\xf6\xbd\x63\x73\x53\xb3\xf3\x58\x7a\xe7\x1b\x6b\x36\xec\x3e\x35\x4f\xd7\x98\x66\x1a\xca\xaa\x44\x15\x54\x4e\x14\x8e\x51\xad\xdf\xf9\x06\x9b\x61\x58\x5e\x10\xb6\x37\xa8\x27\x07\xb6\x80\x1a\x7b\xf4\xbe\xa0\xbe\x01\x67\x18\xf4\x49\xd0\x1f\x30\xbe\xbe\xa5\x38\x8c\xfa\xbd\xc1\x79\x7d\x31\xe6\x10\xa6\xd0\x5c\xf0\x07\xc2\x62\x70\xdf\x43\x97\xd3\x2d\xc9\xb9\x3e\xa2\xac\x30\xa7\x9c\xbd\xbe\x2c\x30\xcd\x4a\x11\xda\x1f\x28\xc1\xd2\xed\x63\xb9\xe2\x65\x96\x22\x46\x1e\xf5\x95\x20\x49\x4a\x81\x8e\xd0\x8a\xe0\xa2\xd5\x15\xea\xf6\xd4\xec\x99\xaf\xbd\x9a\xef\x7b\xe8\x8a\xbc\x61\x8f\x38\xa3\x29\xa2\x2c\x25\xcf\x3d\x56\xd2\xed\x24\x9b\xaf\x7f\x76\xb3\x4c\xd3\x5f\x10\xd5\x4a\x08\xc3\x59\xb6\x46\x4b\x81\x99\xbb\xd2\x50\x0b\x16\x3c\x34\x6c\x7b\x94\xf1\x25\x4d\x5e\x5f\xda\x84\xb4\xb8\xda\x75\xc9\x9b\x51\x7c\xf3\xfa\xc2\xc8\xd3\xeb\x4b\x7d\x55\x8c\x60\xf0\x9b\xf5\x7b\x28\x8e\x96\xf4\x91\x34\xb7\xce\x43\x94\x12\x59\xe8\x25\xde\xd2\xaa\x8c\x29\xa9\x52\xd4\x72\xfc\x1c\xcf\x2f\x74\xb7\xe9\xf3\x5b\x62\x7b\x0f\x76\x44\x74\x74\x5d\x38\xab\xe7\xde\x45\xbe\xea\x79\xab\x1a\x0d\xe3\xea\xc4\x3f\xe1\xb1\x60\xb4\x75\xe7\xa8\xf9\xba\x2b\x95\x56\x17\xfe\xe0\x12\x09\xcc\x42\x4a\xe5\x39\x7a\xc4\x59\x49\x50\x46\xa4\xb9\x49\xb3\x4d\xed\xaa\x30\xf7\x00\x3d\x75\xba\x0f\xdb\xf4\x09\xcb\x4a\xc9\x06\xa9\x68\xcd\x97\xed\x9b\x7a\xa5\xa6\x6f\xdc\x40\xdf\x78\xcc\x9e\xc2\x99\xb5\xd7\x60\xb3\x66\xfa\x14\x83\xc6\xd4\xd0\xb1\x34\x70\x81\x32\x8e\x53\x92\x9a\x59\xe3\xa5\xaa\x9c\xf5\xfd\xca\x41\x2d\x3c\xdc\x09\x6b\xf5\x0c\xfb\x99\xcf\x86\xaf\xe3\xf4\xb1\xe1\x8c\xb4\xd7\xb8\xcc\x02\x17\xe8\x8a\x03\x9e\xe7\x7a\xe4\x1a\x4e\x0a\x22\x16\x5c\xe4\x5a\x50\xd8\x39\x9c\x7d\xbd\xfb\x62\x2d\xcd\x00\x49\x7d\xea\x2b\xc5\x7d\xf4\x5d\x72\xe6\xd6\x76\x8f\xb4\x9b\x71\xbd\x7d\xf4\xdf\x24\xca\xf1\xda\x6d\x8c\xb4\x14\xf5\xb5\x43\xf0\x84\x48\xa9\x7f\xe4\x8b\xb6\xa9\xeb\xd0\xae\x06\xbd\x65\xca\xb9\xd4\xbf\x63\x4a\x9f\xf5\xc2\x0a\xf2\xdc\x4d\xef\x13\x17\x0f\xe8\x89\x64\xd9\x9b\xd0\xf5\x4d\x03\xa3\x05\x17\x96\x04\xb4\xc2\x2c\xcd\x34\x14\xce\xf4\xc4\x2e\x57\x88\xaa\x6a\xd8\x2c\x65\x86\x97\x52\x12\x81\x2c\x64\xe2\xe9\x45\xc7\xbe\x4d\xba\x77\x78\x34\xe5\x76\x06\x35\x42\x9f\x75\xda\x57\xab\xdb\x42\x83\xf1\x56\x37\xdd\x2e\x1a\x56\x67\x55\x1b\x89\xf2\x52\x76\x2c\x67\x0b\x2e\x9c\xae\xa9\x99\xef\x31\x59\x05\x3c\x24\xbd\x8c\x75\xd4\xef\x15\x96\x48\x12\x55\x13\xca\x10\xfb\x83\xfd\x87\x88\x84\x2f\xce\xfa\xe0\x52\x0a\x27\x2b\x73\x3f\x97\x05\x4e\xcc\xc9\x53\x0f\x69\xaf\xd4\xa2\x0b\xe3\xf2\xac\xbf\x92\x46\xd6\xc9\x82\x24\x74\x41\x49\x5a\xad\xe1\xce\xdc\xe8\xa5\xf9\x33\x79\x7e\x83\x8e\x72\x34\x9a\xbe\xfb\xa5\xed\x28\xb9\xfd\x70\x1f\xf4\x6d\x79\x11\x4e\xa3\xb1\xe7\x26\x19\x0f\xb9\x49\xc6\x50\x37\xc9\xb9\x55\x84\xb4\x52\x67\x74\x2c\x69\x03\xb5\x22\xd4\xb9\x31\xd4\x5f\x92\xcf\x05\x4a\xb1\xc2\xd5\xed\x40\x6f\x6a\xa3\xc7\x46\x81\x02\x5d\x27\x35\x28\x4e\xd3\x3d\x11\x81\x2e\x94\x02\x0b\xaa\xd6\xd6\xa2\xb2\xd7\xb0\x02\xdd\x28\x6e\xcd\x95\x25\x4d\xf7\x07\x85\x7a\x33\x52\xf2\x48\x13\x6b\xfd\x58\xf0\x92\xc5\xf8\x89\xc6\x50\x57\xc1\xc6\x80\x6a\xad\x3d\x0a\x0c\x68\xde\xf1\x47\x33\x1a\x11\x68\xce\xac\x57\xe8\x7e\x83\xe9\xad\x96\x7e\xb4\x07\xb2\xde\x0f\xcc\x37\x62\xf6\xc4\x1b\x74\xa5\x4c\xec\x58\xfa\xc6\xc6\xbe\xd0\x8d\x6a\x30\xad\x27\x24\x2e\x3a\x65\x34\xf6\x8d\x81\x03\x78\x5a\xbf\xc2\xd1\xa1\x30\xa3\xf1\x18\x1a\xbd\xb1\xd4\xe7\x6a\x6b\x7d\xc2\x47\xb3\x3e\x78\x2e\xcd\xbe\x85\x1d\x3d\x13\xef\xe8\x99\x0c\x1d\x3d\x13\xe8\xd1\x73\x63\x23\x4e\x91\x4c\x70\x8c\xf9\x60\x02\x3d\x6f\xee\x49\x34\x02\xf0\x70\xf9\x44\xa5\x72\xa2\x30\x0a\x06\x78\xa2\x18\x18\xac\xaf\x31\xaf\x2f\x7b\xa0\x05\xcf\x93\xff\x0a\x8b\x89\x6b\xca\x52\xcd\xd9\xcf\x25\x4d\x7f\x89\x42\x03\x87\x04\xea\x8b\x5f\xec\xde\x99\xf8\x61\xb2\x3d\x30\xee\xc0\x32\x16\xc2\x78\x34\x68\x10\xa0\x43\x4b\x32\x2e\xe3\xa5\xd0\xe4\xad\x17\x01\x38\x78\x18\xbf\xbe\xf4\x0b\x75\x64\x37\x7e\xcb\x1b\x61\xbd\x07\x55\x1c\x84\x56\x68\xfd\xaf\x07\x48\xf3\x62\x05\xc3\xa4\x7d\xe6\xcc\x40\x18\x5b\xc1\x65\x78\xed\xa2\x3b\x96\xad\x51\x4a\xb4\x6a\x4e\xd2\x7a\xec\xac\x3d\xc0\x11\x08\x20\x09\x1a\xc1\x9c\x62\x92\x73\x66\x03\xfb\x63\x66\x05\x1a\xb4\x5c\xe3\xf0\x22\x0a\x06\x28\x83\xdc\x70\xd9\x9b\x4d\xae\x6f\xc6\x29\x51\x91\x91\xee\xa3\x89\xaf\xb4\x0c\xa1\x6a\x71\xf4\x63\x60\x81\xba\xae\xbb\xbf\xc5\x0b\x41\x3f\x52\xb9\x8f\xbf\x7d\x81\x3c\xc1\x14\xde\x20\x1b\x1c\xc5\x8b\x0a\x3f\x3c\xb9\x4f\x54\xfc\x18\x38\xa0\x64\x2a\x6b\xfb\xae\xd9\x6e\x28\xe7\x8c\x2a\x2e\x68\x28\x72\x02\x67\x59\xeb\xef\x6e\xfb\x48\x84\x45\x6d\x5c\x78\x7d\x11\x25\xd3\x17\xfc\x43\xc4\x05\x62\xbc\x6a\x2e\xfb\xe2\x9c\x7c\xba\x81\x62\xab\x4d\x37\x2f\xb6\x91\xfd\xfa\xe2\xd3\xfd\xfa\xd2\x22\xdc\xf4\x52\x90\x34\x9a\x6e\x70\xb4\x74\x4b\xe8\x14\x24\x45\x25\x23\xcf\x85\xd9\x95\xd9\xda\x27\x1d\xd2\x78\x80\x26\x68\x14\xad\xcd\x13\x21\x7b\x6c\x26\x3f\xb4\xa5\x07\x4a\x90\x8c\x60\xb9\x1b\x54\xad\x1b\x1b\x43\xf1\x96\xa8\xd5\xa9\xa7\x13\xbf\x1f\xd2\x89\xa7\xbe\x4e\xdc\x77\x53\x32\x56\x6a\xe7\x80\x4e\x71\xed\x36\x7f\x13\x31\x5e\x53\x5f\x43\x86\xa0\x1a\xb7\xf7\x5e\xb0\xd0\x5b\x93\x85\xed\xf1\x7b\xec\x02\x08\xbd\x39\x59\xc0\xcd\xb0\xc8\x38\x44\x68\x72\xcd\xe5\xfe\x06\x92\x29\x58\x87\x56\x2b\xd2\x0a\x80\x94\x32\x6d\x8c\xaa\x55\x3c\xdc\xd8\x18\xbf\x5a\xf7\xda\x5d\xe8\x80\x2a\xd9\xec\x31\xaf\x8e\x96\xbd\xee\xdf\x53\xb0\xa2\x3d\x93\x2e\x0e\xa7\xe2\x77\x69\x6c\x8b\xc2\xfa\xac\xfe\xf5\x6f\x34\x5f\xab\x40\x70\x29\x84\x04\x68\xfa\xcd\xac\x3d\xdc\xb5\x5f\x4f\x13\x15\x87\x0b\xcd\xc2\xb9\xfd\x70\xff\xfa\x62\x82\x0f\xa2\x87\xd9\x57\x99\xfb\xb1\x5c\xc8\x41\x3c\x16\x54\x26\x19\x0f\xd7\x91\xe2\x19\x11\x98\x25\xc4\x08\x56\xb4\x27\x9f\x50\xc1\xf4\xab\xe0\x6c\xe9\x51\x90\x13\xb5\xe2\x29\x52\xeb\x22\xe6\xf4\x9a\xfa\x3a\x75\x0f\xfa\xc1\xbf\xfe\x8d\xbe\x60\xa1\xa8\xf1\x1d\xd4\x4e\x04\xc3\xb6\x9f\x41\x0d\x41\x86\x4a\xab\x06\x99\x33\xe3\xf1\xdc\x07\x14\x2a\xb5\xcc\x60\xbf\xbe\x58\xd9\x4c\x1e\x4d\xae\x6e\x94\x9c\x04\x27\x03\xba\x40\x09\xeb\x47\xc7\x19\xc2\x69\x2a\x88\x94\x7b\x2c\x2c\xa8\x90\xb0\x0b\x4b\xaf\x20\xe3\x83\xc4\xf6\x3c\xea\xb9\x0f\x7f\x5d\x11\xdb\xf4\x67\xdd\x76\x5e\x2e\x16\x5a\xb0\x5b\xef\x65\x8a\x15\x3e\x92\x8a\x0b\xbc\x24\xbf\xb4\x7c\x50\xf3\xb5\x91\x3d\xed\x8e\x6b\xaf\x29\x4e\x54\x89\xb3\xea\xb7\xa6\x67\xa3\x8f\x55\x51\xb9\x21\x7f\x69\xe3\xaf\xb7\xed\xfb\xe2\x26\xa7\xbe\x6a\x1a\xbe\x51\x19\x87\x6b\xa5\x99\xa1\xe6\xde\x18\x33\xec\xd0\xb8\x6d\xa6\x05\x47\x8e\x29\x33\xee\xb9\xfd\x94\xd0\x29\x34\xbe\x9a\x71\x37\x0b\x46\xcb\x30\xca\x65\x8a\x0e\x14\x16\x4b\xa2\x70\xa3\x62\xee\x8a\x0e\xf4\x26\x30\x42\x52\xe7\x3a\x5e\x50\x21\x4d\xd6\x9a\x5e\x3b\xd6\x8e\xb2\x07\xfb\xd0\xcc\x3f\x6f\x15\xae\xb0\x44\x73\xad\x82\x44\xe6\xcf\x8d\xa6\xbe\xf5\x7f\x57\xe8\xae\x15\x69\x17\x70\xe0\xb4\xdb\x14\x51\xe3\x64\x4e\x53\x8f\x14\x7f\x93\x7b\xc4\xc6\xc5\xd8\xfb\xf4\x02\x17\x4a\x43\xaf\x5b\x1a\xdb\x49\xd6\x72\xa9\x4d\xec\x66\x36\x52\x95\x50\x60\x82\x9a\xc3\xbb\x6d\x80\x6a\xa8\x01\xa8\x45\xa0\xb5\x0d\x24\x19\xc1\x22\x72\x6a\xa1\x27\xd5\xe6\xdc\x6e\xdd\x4d\x8e\xca\x50\xa0\xc7\x00\x31\xd0\x43\x6c\x57\x62\x4a\xf6\xc0\xf8\x13\xdb\x98\xb9\x27\x73\x1c\x6d\x24\x4d\x0c\x50\x06\x55\xc6\x37\x29\x1b\x5a\xf8\xcd\xaa\x71\x54\xb5\xae\xa0\x2e\x4e\xc9\x8c\x21\x80\xba\x40\xb8\xf5\xc0\xca\x91\x0a\xab\x52\xa2\xb2\x48\x23\xb3\x43\xa7\x53\xa0\x30\xb2\xe6\x04\x97\xc7\x7c\x86\xae\xcc\x61\xf0\x99\x8b\x1c\x67\x07\x7e\xa7\x40\x63\x6f\xb0\xd3\x4b\xb2\x14\x38\x25\x69\xa0\x5b\xa0\x35\x37\xd8\xed\x2d\x35\xb1\x58\x81\x5e\x81\x7b\x35\xd8\xeb\x07\x13\x35\x1c\xe8\x14\x68\x98\xed\x74\xda\x3b\xa0\xc0\xf2\x29\x9d\xee\x06\x86\xd2\xdb\x04\xa0\x0e\xef\xc9\xbc\xa4\x59\x1a\x1e\x47\x4f\x79\x04\x75\x39\x53\xbc\xd8\x28\x93\xe4\x30\xb6\x78\x73\x4f\x3c\xcb\xd5\xc9\x90\xe5\xea\x04\xea\xcd\x6d\x38\xbc\x24\xf3\x72\x99\xf1\xa5\xdf\x15\x70\x75\x37\x5d\x55\x36\x49\xbf\x2b\xe0\x8a\x6e\x75\xb5\x51\xbb\xa1\xd5\xd3\x60\xae\x3c\xf2\xbb\x14\xb6\x4b\x54\xe5\xc6\x23\x41\x8a\x26\x2f\xb1\xd5\x31\xb8\x58\xd0\x66\x7f\x11\xc2\xe8\x04\x5a\x25\xc8\xdd\x13\x12\x93\xf2\x56\x18\x71\x9b\xf1\xe4\x21\x4a\x02\x9e\xf8\xc6\x19\x10\xe8\xeb\x0b\x95\xa8\x64\xf1\xb8\x50\x07\x5f\x35\xb0\x09\xcf\x0b\x93\xf1\xe2\xe2\xa1\x16\x65\x16\xb0\x82\x43\x80\xa1\xa6\x8b\x83\x0a\x5a\x10\x59\x66\xaa\x21\xc5\xe5\x1a\xc4\x68\xfe\x27\xd0\xc8\xa5\x5e\x70\xb7\x97\xe2\xc0\xa1\xd7\x0e\x5e\xc3\xd9\x9b\x4e\xf8\xdc\xdf\xb8\xfd\xd9\xe4\x05\x1b\xa7\x9c\x71\x9b\x4f\x85\xd9\x1a\x15\x95\x09\x04\x42\x1f\xf0\x2c\x62\xdc\xde\x8c\x54\x4d\xa7\x4f\xa0\x2d\xbf\xe1\x48\x4a\xdd\x49\x00\xa1\x01\xaa\x46\x5e\xd7\xfa\x91\xf5\x86\x39\x4a\x0e\xab\xa0\x1e\x53\x6e\x8e\x17\x9d\xe8\x6c\x98\xf3\xe2\x82\xb3\x05\x5d\xc2\x02\x7b\x4e\xbd\xa3\x60\xa8\xd8\x5f\xdd\x7c\xeb\x20\x27\x86\x06\xb4\xa0\xa6\x94\x15\x4e\x51\xca\x59\xcc\x5d\xf3\x14\x1a\xe1\xb3\x24\xca\xa5\x10\x44\x23\x01\x8f\x26\x6b\x37\xb2\x58\xf1\xe1\x95\xa7\xd0\xe3\xab\x0d\x17\x19\x33\x77\x1a\x08\xe5\xec\x0b\x8f\xae\xea\xd6\xb9\x09\x34\x16\x9f\x58\x9b\xd8\x69\x20\xac\x73\x1b\xee\x03\x59\xc7\xe3\x05\x7c\xa3\xe1\x6c\x9e\xf6\xf2\x34\x71\x40\xb1\x23\x1b\x28\x24\xb5\x15\xd1\xa5\xfd\xc6\x73\x09\xcd\xd1\xda\xc0\xdc\xcb\x6c\x7e\x1a\x28\x2b\xb5\x7d\x64\xf7\x8c\x86\x3c\x1d\x41\x93\x12\x7f\x97\x9c\xa1\x94\x27\x95\xc4\xe6\xf3\xdf\x49\x12\x38\x77\xaa\x45\xd6\x08\x8b\xbf\xd9\xfd\x65\x32\x54\xcd\x2f\x6c\xa6\xce\xe6\xc2\x38\x6c\xf9\x1c\xd0\xdf\xc2\x6c\x05\x25\xf1\x2d\x2e\x8a\x6d\x25\xab\xc6\x5e\xf5\xa3\xf1\x71\x47\x04\x37\x6e\x69\x1b\x43\xcb\xb7\x75\x39\xf2\xba\xec\x66\x0a\xd4\x5d\x7e\xbc\x18\x3e\x21\xc6\x5e\xe8\xe7\x78\x28\xf4\x73\x0c\x0e\xfd\xfc\x78\x81\x94\xa0\xcb\x25\x89\xb2\xd8\x8c\xc1\xa1\x9f\x1f\x2f\xe2\x0b\xb7\x8e\xc1\xe1\x9f\x1f\x2f\x5e\x5f\xe2\x0e\x9e\x31\x38\xf6\xf3\xe3\x45\x5d\x01\x27\x32\x62\x6d\x0c\x8e\x8c\xfb\x4e\x13\x45\xf3\x5a\x5d\x4f\x38\x93\x4a\x94\x89\x8a\xd9\xc8\x63\x70\x9c\xdc\x27\x8e\xb5\xce\xfa\x48\x84\x24\x28\xc7\x11\xc1\x72\x63\x70\xb0\x9c\xc1\xb2\x87\xab\xa9\x95\xb0\x6b\x81\xb3\x5b\xa2\xf0\xf5\x6c\x63\xf3\x4c\xbb\x9b\x67\xe2\xed\xed\xc9\xd0\xe6\x99\x04\x2a\x9b\xf5\x44\xb7\xdc\xda\x7f\x34\x6e\xa4\x54\xdf\xb8\x97\xa1\xb0\xa4\xad\x63\x36\x09\x24\xaf\x85\xc7\xcc\x43\xd5\x5c\x68\x61\xd8\xd6\xd0\xe1\xb8\xa3\x40\xd2\x5c\xf0\x4a\x71\x53\x17\x26\x46\xb7\xd7\xb3\x8d\x84\xec\xdd\xf0\x80\xce\xa2\x5b\xe3\x34\xd9\x0f\x0a\x78\x55\xfb\xe6\xaa\x4f\xec\x07\x06\xbc\x9a\x35\xb7\x0e\xc2\xfe\x28\x49\x49\x4c\x2d\xe6\x5c\x06\x62\xd9\x7a\x9b\xea\x5f\xdd\x5e\xcf\x0e\x6d\xae\xa5\xbb\x48\x91\xe7\x02\xb3\x14\x2d\x04\xb1\xd5\x87\xff\x09\x21\x1a\x98\xa6\x72\xc1\xf3\x02\x27\x6a\xa0\x82\x6a\xbb\x49\xc2\xcb\x2c\x65\x7f\x37\x11\x18\x5a\x20\xa3\xb4\x34\xc1\x7b\x26\xd6\x87\x99\xec\x4e\x43\xa5\xc9\xde\xfb\xcf\x50\x19\x43\x42\x0d\x18\xc8\xda\xeb\x0b\xc4\xb8\x9e\x69\x7d\x27\xba\x3a\xdf\x24\x90\xad\x37\x0c\xb5\x47\x29\xc0\x49\x20\x4b\xaf\x0f\xac\x2b\x64\x22\xf5\xc7\x49\x20\x4d\xaf\xaf\x16\xaa\x0b\x06\x28\xb0\xc0\x39\x69\x95\xd4\xda\x0d\x0e\x1a\x62\xa1\x17\x82\x71\x70\x44\xa1\x40\x4d\x0c\x95\x13\x25\x1e\x69\x77\x37\x54\x53\x41\x3e\x0a\x70\x77\xef\xd2\xdc\x04\xda\x96\x11\x9a\xc2\x64\x0c\x0e\xeb\x62\x1c\xe5\x24\xa5\xd8\xac\xc6\xdb\xeb\x59\x14\x18\x34\x96\xcb\x15\x95\xda\x28\xdb\x35\xa7\x31\xc7\x7a\x20\x83\x73\xeb\x78\xea\x3b\x4f\x14\x14\x54\x94\xd0\xba\x1a\xca\x3e\x9a\x43\x20\x87\x73\x2b\x6b\x2e\xd5\x58\x5f\xed\xa2\x10\xa1\xc2\xa4\x41\x5c\x50\x17\x75\x13\x8d\x09\x95\x28\x15\x48\x7d\x01\x2e\x48\x54\x4d\xfc\xc9\x18\x6c\xc1\xfc\x91\x6c\x42\x05\x8d\x12\xb4\x59\xa6\x0d\xd3\xb1\x7a\xd3\x18\x5c\xc9\xbd\xb5\x8a\x4c\x1e\xd7\x5e\xdc\x42\xc5\x4e\x6b\x84\xad\xc5\x61\x2f\x54\xa8\xfc\x69\x50\x5d\x91\xb0\x7d\x50\xc1\xc9\x0a\xed\x05\x5c\x1d\x5c\xb1\xee\xa1\xc9\x18\x5c\xde\xbd\x1d\xf9\x82\xf7\x5b\xc4\x01\xab\xdc\x36\x4c\x1b\x46\xeb\xca\x1d\x47\xe3\x42\xe5\xd2\x93\xe0\xaa\x29\xa2\x68\x0b\x27\x29\xe3\x7a\x31\x71\x20\xa2\x75\xee\xf8\x1a\xb6\x3e\xf9\x04\xce\x51\x4a\xe5\x03\x84\x28\xa8\xe0\x6a\x06\x43\x12\xf2\xd0\x3f\x10\x3b\x13\x00\x15\x63\x66\xde\xff\xa2\x41\x81\x8a\xb9\xbf\x74\xa6\xa0\x22\xd0\x05\x12\xf6\xdc\xc0\x76\x45\x85\xca\xc0\xaa\x38\xe3\x8f\xc2\x85\x4a\xc1\x4a\xe5\xdb\x1b\x11\x5c\x1f\xc1\x53\x54\x72\xa2\x30\x72\xe1\xbd\x31\x82\x01\x5c\x28\xa1\x7b\xb8\xed\x0d\x0c\xae\x98\xd0\x02\xaa\x56\x7a\xf4\x6d\x05\x5c\x3a\x21\x3c\xd0\x3d\xb5\x36\x21\xc0\xbb\xcb\x3a\xa3\xbe\xec\x0b\xbb\xbb\xa2\xd6\x9a\xde\x3d\x70\xc1\x19\x40\x8d\xb9\xa3\x2e\xb0\x64\x47\x3c\x0a\x16\x2a\xa7\x70\x96\x73\xa9\xd0\xa2\x0c\xd4\x15\x86\xe0\x40\x25\x53\xa5\xa5\x98\x41\x8d\xd2\x52\xc6\x50\x61\x64\x8a\xf3\x63\x85\x33\xbe\x0c\xd5\x78\xdc\x01\xd2\x2f\xfd\xdb\xf7\x74\x91\x33\xfa\xc5\xbe\x76\x33\x19\xfb\x35\x7b\x87\xd7\xa7\x39\xda\x1e\xb1\xa0\xbc\x94\x5a\x0a\xc8\xb8\x8b\xe1\x64\x77\x35\x4c\xe2\x47\x82\x28\xe3\x69\xd4\xc2\x9c\xec\x2e\x77\x0a\x5e\xd8\xd7\x26\x30\x32\xa3\x1c\x05\xbb\xbb\xd4\x29\x4a\xb9\x32\x81\x35\xd1\xa8\x01\xe7\xde\xc5\xfd\xcd\xd7\x9b\x8b\xf3\x4f\x3e\xf0\x45\xb3\x62\x77\xa9\x0f\x5b\x3b\x58\x66\x6b\x79\x64\x52\x81\x36\x7c\x2c\x27\x5d\x17\x8b\x97\x86\x3b\x99\xbe\x1f\x70\xb1\x04\xd2\x70\x7b\x5c\x2c\x56\x44\xfe\x66\x48\x88\x18\xab\x40\x46\x5d\x0f\x90\x8d\x84\xfa\x0d\xa7\xf1\x58\xfe\x3d\x6b\xe8\x19\xa3\xdf\xdc\xbb\x42\xbb\xe0\xd5\xd3\xf2\x89\x2f\xfd\x67\x98\xba\x73\x72\xe2\xcd\xc9\x90\xdb\x6b\x1a\xa8\x57\x1a\x4e\x54\xa0\x99\x22\x02\xc9\x35\x53\xb8\xaf\x94\x2e\x60\xb4\x4e\xe0\x39\x43\x0e\xd1\x5e\x4d\xad\xc9\xa1\xc0\x6a\x15\x85\x0a\xf4\x08\x35\x01\xf1\x19\x5f\x1e\x99\xa6\xf6\x7e\xa4\x76\x16\xfe\xf5\xa4\xfd\x93\xcf\xc2\xef\x67\x75\x26\xee\x9d\xb7\x99\xde\x0d\x4d\xdc\x3b\xa8\xe3\x50\x13\x6f\xd3\xfa\x0a\x9e\xd1\x64\xbd\xf1\xca\xd7\xdd\x17\xac\x56\x47\xec\x31\x5f\x0c\x87\x22\x4c\x3d\x6f\xea\xb4\xfb\xc4\x6b\xb7\xd3\x85\x30\x15\xaa\xb7\x04\x44\x4f\xbd\x57\xc3\xa6\xbd\x21\x18\xae\xe3\x39\x4e\x1e\xb6\xf7\xeb\xc5\x61\x4c\xbb\x0f\x5b\xd5\xfd\xda\xb8\xbf\x6e\x7f\x9d\x0e\x4f\xbd\x11\x38\x1d\x0a\xd7\x3b\x05\xbf\x94\x35\xfb\x72\x7e\x71\xb5\xf9\x37\xf0\xca\x3e\x0d\xbc\x92\xb5\xf5\x2c\xaa\xee\x75\x1b\xc5\x70\x5b\xb7\x2c\x57\x5b\x37\x54\xaa\x77\x80\x8e\xdd\xcf\xfe\x6d\x74\xac\xe8\x72\x65\xde\x37\xa5\xdc\x54\x1e\xfc\x9d\xcf\x91\x2c\x93\x15\xc2\xb2\x8a\xd9\xa4\xcc\x9c\x71\x55\xec\x36\x65\xcb\x50\x12\xa5\xea\xd4\x02\xc6\xe8\x69\xd5\x18\x3a\x5a\x3c\x40\xd5\x89\xd6\x83\x12\xad\x1a\xbf\xe4\x99\x24\xa5\x7d\x91\xc2\x24\x4c\xd7\xef\x09\xf7\x16\x3c\x76\x4d\xda\x45\x5b\x4d\x9e\x5c\x7f\x59\xe3\x2e\x2b\xc1\x31\x6c\x71\x04\x8c\xcb\x35\xb9\xc6\x2e\xfc\xf5\xd1\x3e\x64\x7d\xee\x93\x3b\xd0\xaa\x05\x09\x8c\xbb\xaf\x2a\xe8\xd5\x2f\x1f\xf8\x78\x05\x97\xc1\x66\x4d\x19\xbd\xa3\xf3\x65\x57\x49\xd1\x3b\x37\xf0\x3e\x75\xf7\x69\xc2\xc1\xcd\xab\x3f\x80\xea\x78\xdf\x58\x4a\x16\x94\x91\x74\x23\x40\xae\xdd\x13\xf4\xb2\xf8\x0f\xc9\x19\xfa\xba\x2e\x48\x27\xd4\xae\x9b\x0d\x5c\x79\xa9\x3e\xf0\x74\x6d\xda\xfb\x98\x7e\x36\x56\xaf\xc7\xab\x7a\x5b\xd3\x46\xd8\x19\xd5\xf8\xcb\xdd\xcc\xeb\xf2\x18\xec\xd0\xfa\xc6\x70\xa9\x56\x5c\xd0\x3f\x49\x8a\xbe\x49\xd2\xcf\xc8\xb9\x6b\x67\xf3\xeb\xff\x9b\xe0\x94\xf8\xe3\x77\x0c\xb6\x63\x9b\xf1\x30\x83\x38\x3c\x7e\xa6\xdd\x17\xbc\x36\xcb\xd9\x9a\xf7\x7d\x54\xb0\xed\xe8\x7f\x8e\x9c\x24\x38\xba\x49\x1d\x0f\x5b\xf0\x37\xbe\xf8\xb9\x2c\x69\xfa\x0b\xfa\x8e\xb3\xd2\x9f\xc7\x40\x24\x4b\x5f\x9e\x8d\x5d\x12\xe7\xe6\x45\x26\x57\xb7\xbd\xc0\x52\x5a\x49\x1b\x98\xd0\xd0\x23\xec\x7d\x95\x35\xdc\x93\xaf\x3d\x05\xc5\xdd\x93\xde\x7a\x0a\xdd\xef\xea\x6a\x27\x5f\xee\x66\x86\x12\xab\xff\x98\xf2\x64\x33\x85\x93\x80\xed\x0e\x09\x62\x83\xf1\x83\xa4\x82\xd7\x5e\x60\xf6\x6b\x3a\x6f\x71\xe6\x4a\xb0\xff\x43\x86\x92\x1a\xf4\x5a\x75\x41\x69\xa6\x9b\xf0\xca\xd8\xe1\xc9\x5c\xfa\x81\x2f\xd0\x45\x33\x3a\x3d\x74\xfd\x05\xe3\x07\x2f\x14\xcb\xb8\x5a\x99\x67\xf7\xec\x7a\x6a\x25\x83\xf6\xe6\x47\x21\xdc\xf9\xc8\x68\xca\x75\xfb\x37\xe8\x57\x4c\x95\x8d\x58\x53\x7f\x97\x55\x3a\x50\xe3\x7b\x6e\xd1\x09\x36\x91\x9c\x4b\xc9\x13\x6a\xca\x94\xeb\x31\x4a\x70\x96\xf5\x9a\xa9\xab\x06\x5a\x69\x50\xa5\xd0\x92\xda\xca\x3a\x53\xd4\x41\x56\xc9\xa8\x81\x31\xed\xbe\x21\x20\x79\x4e\x90\xa2\xde\xcb\x3c\xc7\xc7\x70\x81\xfb\xbd\xfd\x88\x46\xf5\x44\x60\x46\xf3\xc0\xcb\xed\x76\xfa\xb3\x8c\x3f\x49\xc4\x59\xb6\x46\xa3\xe9\xbb\xe6\x59\xc5\xca\xa0\xf7\x06\x15\xb6\xd8\x96\x2d\x93\x1c\x7a\x32\xb0\xcb\x93\x4b\x71\x37\xcf\x78\x11\x51\xe5\x37\x70\x61\xa9\xea\x32\x17\x7a\x5b\xb3\x87\x39\xce\xcc\xfb\x57\xf4\x03\xbf\x36\x0f\xc5\xbb\x94\xe2\x94\xe0\x6e\xe2\xdf\x71\xe8\xf9\xc2\x9e\x5e\x13\xfb\x8c\x07\x57\xd6\xdb\x6d\xd3\x9b\x68\xf3\x98\x53\xdd\xe7\x18\x2e\xd3\xee\xab\xf7\xfa\xf5\x74\xf2\x52\x6d\x44\x97\xc2\xb4\x8a\x91\xaf\x55\x8c\x06\xb5\x8a\xd1\x5b\x70\x31\xea\xcb\xe1\x42\xd4\xed\xc4\x6d\xdd\xd2\xd9\xef\x0b\x41\x64\xdb\xa8\xd0\xcc\x7a\x52\x15\xd0\x37\xd9\x5f\xee\x93\xa6\x13\x53\x0b\x85\x35\x4f\x69\x12\xa6\x9a\xa7\x2a\x6b\xf2\x77\x38\x32\xfe\x59\x12\xb1\x8e\x3b\x33\x6e\xd8\x22\x2b\x9f\x2f\x3f\x98\x95\x6c\x26\x62\x40\xda\x55\x8d\x7d\x52\xc1\x6b\xf6\x2b\xcd\x49\x13\x7c\xd6\x47\x70\x15\xa5\xa6\x4c\x6b\x22\x28\x4f\x9b\xaa\x2f\x01\x02\x4b\xb7\xb4\xac\x42\x75\x9c\x1f\x4e\xf3\xc3\x63\xfd\xff\xea\xf0\xdd\xea\xf0\x78\xb4\x3a\x1c\x4d\x56\x87\xef\xd3\xc3\xf1\xdb\xb4\xbd\xf6\x3e\x7f\xbf\xbd\x42\x38\xcd\x29\xdb\x92\x1a\xe0\x2f\xbe\xf1\xdb\xee\x8d\xbc\x3d\x24\xcd\x07\xdb\x87\xc4\x10\x91\xa8\xac\x77\x34\x2e\x30\xfb\xbb\xcd\xe4\x60\x8f\x5a\x98\x29\x91\x75\x15\x37\x8d\x07\xbd\x9a\x7e\xe2\x4b\x54\xe0\x65\x5f\x5a\x4e\x0b\x2e\x73\x2d\x03\x60\xd0\xfb\xe7\x45\x7d\xfe\x0c\xc3\xb9\x67\x50\x5a\xe7\x95\xf4\x26\xea\xfc\xff\x79\xa2\x7e\xf0\x3c\xfc\xd8\x71\xee\x1d\xc5\xff\x0b\x00\x00\xff\xff\x41\x8b\xfd\x98\x47\x89\x00\x00")

func ResourcesEventsYamlBytes() ([]byte, error) {
	return bindataRead(
		_ResourcesEventsYaml,
		"../resources/events.yaml",
	)
}

func ResourcesEventsYaml() (*asset, error) {
	bytes, err := ResourcesEventsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../resources/events.yaml", size: 35143, mode: os.FileMode(420), modTime: time.Unix(1602095277, 0)}
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
	"../resources/events.yaml": ResourcesEventsYaml,
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
	"..": &bintree{nil, map[string]*bintree{
		"resources": &bintree{nil, map[string]*bintree{
			"events.yaml": &bintree{ResourcesEventsYaml, map[string]*bintree{}},
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
