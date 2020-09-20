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

var _ResourcesEventsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5d\xdd\x72\xdb\x38\xb2\xbe\xcf\x53\xa0\x7c\xb1\x3b\x53\x65\xcf\x89\xf5\x93\xd8\xb9\x73\xfc\x93\xe3\xad\x38\xf6\x5a\x76\xf6\xec\xd5\x14\x44\x42\x12\xc6\x24\xc0\x01\x40\xd9\x9a\xa9\x79\x17\x3f\x8b\x9f\xec\x14\x7e\xf8\x23\x01\xa4\x5a\x50\xce\x9e\xa9\x9a\xaa\xd8\x26\xf1\x35\x1a\x8d\x46\x77\xa3\xbb\x79\x74\x74\xf4\xee\x0b\xc7\xd9\x27\x74\xb9\x24\x4c\x49\xf4\xc8\xe8\x8c\x26\x58\x51\xce\xde\x7d\x27\x42\x52\xce\x3e\xa1\x83\xe5\x2f\x27\x1f\x0f\xde\xe5\x3c\x2d\x33\x22\x3f\xbd\x43\xe8\x08\x31\x9c\x93\x4f\xe8\xe0\xfc\xf6\xe6\xe6\xf6\xdb\xc1\x3b\x84\x10\x4a\x78\xc9\xd4\x27\x74\x7a\x7a\x6a\x7e\xa4\xe9\x44\x61\xa1\x3e\xa1\xf7\xee\xc7\x4b\x96\x7e\x42\xa8\xf9\x3b\x9b\xf1\x4f\xe6\x5f\x7a\xbc\x84\xa7\xa4\x7a\x54\xff\x97\x91\x25\xc9\x3e\xa1\x83\xeb\x6f\x57\xb7\x07\xf5\x6f\x73\x22\x25\x9e\x6b\xe0\x49\x99\x24\x44\xca\xe6\x4f\x85\xe0\xd3\x8c\xe4\x9f\xd0\x41\xf3\x3b\xc9\xb3\x52\xd9\x29\x1c\xac\x51\xfd\xf5\x7a\x8d\xe4\xe3\xf7\xef\xd7\x49\x3e\x7e\xff\x7e\x83\xea\xe3\xf7\x3d\x64\xd7\x8f\xb7\x28\xd7\x84\xfb\x74\xd3\x29\x9f\x71\x89\xa8\x44\x52\x23\x91\xd4\xa3\xdf\x27\xde\x07\x3b\xde\x15\x4c\x11\x91\x53\x86\x63\xf1\x06\x30\xbc\x24\xa3\x48\x12\xb1\x24\x42\x63\x52\x46\x15\xc5\x19\xfd\x23\x12\x74\x08\x03\x65\xe4\x59\x03\x13\xa6\x34\x68\xc2\x19\x23\x49\xec\x3c\x47\xe0\x79\x3a\xb8\x94\xca\xfd\x10\xc7\x30\x44\x41\x7e\x2f\x89\x54\x28\x97\x73\x0d\x2b\x48\x42\xe8\x32\x12\xf2\x03\x14\x52\x16\x9c\x49\x52\x61\x4a\xc2\x54\x0c\xde\x31\x70\x67\xb4\x84\xa7\x10\xa4\xc0\x82\xb2\x39\x22\x2f\x34\x0e\x14\xb8\x43\x5a\xa0\x6a\x21\x08\x4e\xd1\x6f\x9c\xb2\x38\xc6\x1e\x03\x77\x09\x61\x78\x9a\x11\x24\x48\x29\xc9\x11\x4e\x53\x11\x05\xe6\xed\x8e\x7f\x9d\xdd\x7f\x03\x80\xa1\x19\xa6\x59\xe4\x04\xbd\xed\x71\x79\x7f\x7f\x7b\xef\x83\x4a\x9e\x3c\x11\x85\x12\x41\xcc\x41\xb2\x0f\xa4\xb7\x3f\xfa\x21\xa7\x94\xa5\x5a\x6e\xf6\x40\xf4\xb6\x47\x3f\x62\x46\xa5\x22\x7b\x4d\xf1\x23\x10\x90\x14\x3c\xcb\x7e\x08\x53\x4f\x76\x9b\x22\x4e\x12\x52\xa8\x7d\x00\x4f\x81\x80\x39\x7e\xa9\x94\x2b\x11\x82\x47\xed\x8c\x81\xa7\x6f\xba\xc0\xec\x3f\xb4\x62\xdb\x53\x66\x06\x9e\xb6\xd9\x02\x69\xf5\xf7\x9e\xa0\x9e\xba\x09\x6b\x80\x87\xeb\x9b\xcb\x0b\x74\xfb\xf8\x10\x05\xe2\xa9\x99\x8e\x99\x5d\x7f\xfb\x7e\xf6\xf5\xfa\x02\xdd\x9d\xdd\x9f\xdd\xc4\x20\x0d\x81\xc7\xc4\xdd\xed\x04\x5d\x4f\xd0\xe7\xc7\xc9\xbf\x61\x30\xb5\xd1\x77\xfd\x99\x5f\xdd\x4e\xd0\x44\x61\x45\xd0\x0d\x66\x78\x4e\xc4\x9a\x15\x38\xf0\xac\xc0\x63\xcf\x0a\x1c\xf4\x59\x81\xc7\xbe\x61\x76\x71\xf9\xf9\xf1\x4b\x60\x67\x19\x22\x12\xce\x14\x79\x51\x08\xa7\x69\x94\x0c\x1c\x43\x0d\x33\x07\xb7\xc0\x6c\x1e\x05\x34\x18\x7a\xaa\xff\xfe\xf2\xee\xf6\xfe\xe1\xd7\x87\xfb\xb3\xf3\xcb\x80\xe5\xe9\x78\xbd\x92\x8a\xe4\xe8\x9e\x24\x7c\x49\xc4\x0a\x5d\xb3\x42\xf0\xb9\x20\x52\xee\xb8\x76\xdf\x79\x56\xe6\x24\xb4\x68\xc3\xcd\x45\x1b\x78\xa6\xfb\xa0\x6f\xd1\x06\x60\xd3\xdd\xd2\x60\x15\x70\x0c\x17\x07\x60\xbb\xdd\x21\xa5\x24\x23\xb1\x48\x40\xc1\x70\x48\xb9\xe6\x65\x24\x12\xd0\x4c\x77\x48\x25\xdb\x07\x0b\x68\x9f\x3b\x2c\xb3\xad\x90\xe2\x48\x2d\x88\x39\xa7\xa3\x30\x81\x16\xba\xc5\x7c\x7b\x15\x24\xe7\x4b\x92\xa2\x99\xe0\xb9\x06\x7e\x7b\x8d\x45\xf6\x0d\xe7\x0e\x05\xbf\x20\x95\x7f\x40\x52\xb4\x74\xb2\xc3\x89\x44\x8c\x2b\x6d\x41\x07\xe0\xcd\x4b\xee\xd1\x67\xaa\x16\x86\x45\xde\x20\x7a\xf3\x21\x2e\xaa\x1f\xaf\x2f\xba\x86\x6d\xfc\xec\x4b\xa6\x8c\x2d\xad\x95\x9b\x10\x24\x51\xdd\x63\xe1\x99\x7e\x34\x59\x90\xe4\x49\x9f\x84\xaa\xa1\xa8\xc5\xb1\x16\x3b\xbc\xcd\x13\x66\x87\x5b\x7c\x49\xff\x20\xc6\xf3\xe5\x1c\x65\x58\xcc\x89\xcf\x83\x6a\x43\x63\x66\xa6\x34\xad\xf7\x36\x9a\x92\x04\x97\x92\x18\x9a\x72\xfc\x42\xf3\x32\x47\xac\xcc\xa7\x44\x20\x3e\x73\x54\x4a\xa4\x16\x58\x99\xb7\x5b\x6f\x52\x89\xc8\x4b\x42\x48\x5b\xa1\x37\xdc\xb9\x27\x4a\xac\xdc\xc4\x8d\xa0\xe8\x89\x97\xda\x7f\xd4\xd4\x8b\x55\xc5\x00\x9c\x73\xeb\xfc\x48\xa5\x9f\xa8\x06\x77\xc0\x3e\x6b\x80\xa6\xc0\xa5\xa1\xcc\x0e\x67\xdc\x82\x65\xc3\xab\xb0\x88\x78\x22\x61\xd8\x6a\x67\x28\xd1\x99\x10\x78\x85\x12\x5c\xe0\x84\xaa\x55\x60\xbe\xe7\x7a\x71\x0d\x17\xa5\x3d\x09\xaa\x67\x11\x66\x29\x32\xbc\x98\x63\xca\xbc\x09\xf9\x36\x5c\xef\x5a\x57\x2b\xb8\xc0\x4b\x82\x30\x9a\x66\x98\x3d\x19\x81\xf3\x27\xf5\xd9\xfc\x8d\x5a\x21\xc6\x59\xc6\x9f\xf5\x5e\x6d\xe4\x72\xed\x2d\xa8\x58\xb7\xde\xaf\xcc\x73\x7f\x4a\xbb\x89\xaf\x1e\xf7\xed\xb5\x16\x60\xce\xe6\x9d\xf2\x6b\x48\x20\x9a\x40\x2b\x7f\xde\x0b\x9b\xd3\xc0\x6b\xc4\x1b\x29\x6e\xed\x6b\x23\x22\x83\xf1\x58\x9b\x0b\x02\x27\x8a\x08\x5f\xe0\xa0\xb6\x67\x9b\x42\x2a\x51\x5a\x16\x19\x4d\x82\xa7\x27\x3a\x5b\x53\x47\xb8\x79\xd6\xbe\x8d\x33\xed\x9d\xaf\xec\x9e\x90\x3d\x53\x4b\xe9\x6c\x46\x84\xf6\x20\xfc\x15\x6d\x4d\x00\xe8\x3e\x3f\x32\xeb\xfc\xb4\x77\x4b\x58\xb0\x42\x8b\xa1\xed\x3b\x4c\x99\x44\x18\x49\x25\xac\x92\xc3\x6a\x43\xfa\x7a\xa5\xcd\x5b\xa8\x9c\x10\x25\xbd\x3f\x89\x32\x0b\xa8\x05\xdf\xa4\x0e\x4f\xf2\xbe\xde\xe4\x66\x77\xf7\x68\xcc\x6e\x95\x10\xd4\xb4\x61\xc5\xb7\xa6\xf1\x0b\x2e\x25\x0d\xab\xa2\xd6\x44\x80\x7b\x27\x30\x11\x99\xe3\x2c\xdb\x7d\x22\x6f\xaf\xeb\x2f\x86\x54\x5a\x4e\x99\x39\x18\xf4\x3a\x26\xbe\x42\x35\x2a\x4e\xe8\x79\xfb\xf3\x01\x6e\x9f\x8d\xf9\x58\xa1\xa1\xf3\x50\x7c\x2a\x38\xa3\xb7\x57\xf3\x5e\xbd\xb9\xcd\xcb\x76\x83\x4d\x33\x9e\x3c\xad\x2b\xfe\x66\x8e\xd7\xac\x28\xd5\xda\x5c\x14\xd7\x87\x5c\x5e\x66\x8a\x16\x19\xd1\x87\xa0\x37\x40\x33\xbd\xd1\x6e\xda\xbb\xda\xda\x5d\x86\x21\x3a\x53\x8a\xe4\x85\xd2\x44\x98\x67\x1a\x05\x56\x6d\xa7\xae\x21\x9a\x29\x7d\xe3\x6a\x61\x64\x8e\xa3\x94\xfb\x14\xef\xa6\x9c\x2b\xb8\x6e\x63\xb6\x4d\xb3\x7b\xaa\x9b\x6a\x7f\x18\x28\xdd\x63\xdf\x75\x61\x33\xee\xd3\x7d\xee\x0e\x48\xeb\x51\x54\x6c\x72\xf4\x74\x12\x6f\x8d\x67\xf7\x0e\xee\x7a\xab\xa1\xb5\x35\x67\xf7\x52\x65\xce\x98\x7d\xef\xa6\xb9\x6e\xeb\x79\x53\xfa\xe0\x87\x6f\xd6\x22\x41\xf5\x9c\xae\x30\xcd\x34\x94\x35\x8e\x2a\xa8\x9c\x28\x1c\x63\x6c\x7f\xf0\x43\x38\xfd\xb0\xbc\x20\x6c\x6f\x50\x4f\x0f\x6c\x01\x35\x11\xea\x7d\x41\xfd\x90\x4e\x3f\xe8\xb3\xa0\x3f\x80\xbf\x7e\xec\x38\x8c\xfa\xbd\xc1\x79\x7b\x35\x01\x12\xa6\xd0\x54\xf0\x27\xc2\x62\x70\x3f\x42\xc5\xe9\x86\xe4\x5c\x1f\x51\x56\x99\x53\xce\xde\x5e\x67\x98\x66\xa5\x08\xed\x0f\x94\x60\xe9\xf6\xb1\x5c\xf0\x32\x4b\x11\x23\x4b\xed\x1c\x24\x49\x29\xd0\x11\x5a\x10\x5c\xb4\x86\x42\x9b\x23\x35\x7b\xe6\xa1\xd3\x06\xfe\x08\x95\xc8\x6b\xb6\xc4\x19\x4d\x11\x65\x29\x79\xe9\x88\x9b\x6e\x27\xd9\xbc\xfd\x93\x5b\x65\x9a\xfe\x8c\xa8\x36\x42\x18\xce\xb2\x15\x9a\x0b\xcc\x9c\x73\x43\x2d\x58\xf0\xd0\xb0\xcf\xa3\x8c\xcf\x69\xf2\xf6\xda\x26\xa4\x35\xab\x5d\x45\xde\x70\xf1\x97\xb7\x57\x46\x9e\xdf\x5e\x6b\xe7\x31\x62\x82\x8f\xf6\x26\x44\x71\x34\xa7\x4b\xd2\xf8\xa1\x87\x28\x25\xb2\xd0\x22\xde\xb2\xaa\x4c\x70\xa9\x32\xd4\x72\xfc\x12\x3f\x5f\xe8\x6e\xd3\xe7\xb7\xc4\xd6\x33\x76\x44\x6c\xd8\xba\xf0\xa9\x9e\x79\xae\x7d\x35\xf2\x56\x33\x1a\x36\xab\x13\xff\x84\xc7\x82\xd1\x96\xcf\x51\xcf\xeb\xb6\x54\xda\x5c\xf8\x9d\x4b\x24\x30\x0b\x19\x95\x67\x68\x89\xb3\x92\xa0\x8c\x48\xe3\x53\xb3\x75\xeb\xaa\x30\x7e\x80\x5e\x3a\x3d\x86\x7d\xf4\x19\xcb\xca\xc8\x06\x99\x68\xcd\x9b\x6d\x9f\xbd\x32\xd3\xd7\x7c\xd1\x5f\xbc\xc9\x9e\xc2\x27\x6b\x1d\x62\x23\x33\x5d\x86\x41\xe3\xb2\x6e\xc4\x1c\xb8\x40\x19\xc7\x29\x49\xcd\xaa\xf1\x52\x55\xd7\xf7\xdd\xc6\x41\xad\x3c\xdc\x09\x6b\xed\x0c\xfb\x9a\x3f\x0d\xdf\xc6\xe9\x9a\x86\x0b\xdb\x5e\xe1\x32\x0b\x84\x8e\xaa\x19\xf0\x3c\xd7\x9c\x6b\x66\x52\x10\x31\xe3\x22\xd7\x8a\xc2\xae\xe1\xe4\xe1\xf6\xce\xc6\x9e\x01\x9a\xfa\xd4\x37\x8a\xbb\xe8\xbb\xe0\xcc\xc9\x76\x87\xb6\x9b\x70\xbd\x7d\xf4\xdf\x24\xca\xf1\xca\x6d\x8c\xb4\x14\xb5\xdb\x21\x78\x42\xa4\xd4\x3f\xf2\x59\x3b\xf8\x75\x68\xa5\x41\x6f\x99\x72\x2a\xf5\xef\x98\xd2\x67\xbd\xb0\x8a\x3c\x77\xcb\xfb\xcc\xc5\x13\x7a\x26\x59\xf6\x4b\xc8\x7d\xd3\xc0\x26\x28\x60\x48\x40\x0b\xcc\xd2\x4c\x43\xe1\x4c\x2f\xec\x7c\x81\xa8\xaa\xd8\x66\x29\x33\x73\x29\x25\x11\xc8\x42\x26\x9e\x5d\x74\xec\x47\xa9\x3b\xd9\xa3\x29\xb7\x2b\xa8\x11\xba\xe2\xd5\xbe\x59\xdd\x56\x1a\x8c\xb7\x86\xd9\x1c\xa2\x99\xea\xa4\x7a\x46\xa2\xbc\x94\x1b\x31\xb4\x19\x17\xce\xd6\xd4\x93\x6f\x82\x57\x75\x54\xff\xe6\xf3\x7d\xf0\x22\xc6\x4b\xc7\x19\x0c\xbd\x98\xfe\xb0\x2f\xa6\x3f\x84\xc6\xf4\xcf\xec\x19\xad\xed\x0d\x73\xfc\x4b\x9b\x55\x14\x61\x69\x0c\xa1\xc1\xfd\x7c\x2a\x50\x8a\x15\xae\x0c\x57\x2d\x6f\xc6\xc4\x8a\x02\x05\xc6\xf9\x6b\x50\x9c\xa6\x7b\x22\x02\xe3\xfd\x05\x16\x54\xad\xac\xb3\xbf\x17\x5b\x81\x31\x7f\x27\xaa\x65\x49\xd3\xfd\x41\xa1\xa1\xf7\x94\x2c\x69\x62\x1d\xf3\x19\x2f\x59\xcc\xa5\xc6\x10\x1a\xd7\x5e\x63\xa8\x36\x28\xa3\xc0\x80\x91\x07\x9f\x9b\xd1\x88\xc0\x48\x5b\x2d\xa1\xfb\x31\xd3\x93\x96\x6e\xb4\x27\xb2\xda\x0f\xcc\x8f\xaf\x75\x5c\x8e\x6f\x6a\x99\x58\x5e\xfa\x71\xb0\xae\x3c\x83\x8a\x99\x36\x5c\x1f\x97\x4a\x31\x18\xfa\x71\xaa\x1e\x3c\x7d\xf4\xe3\xe8\xbc\x8d\xc1\x70\x08\x4d\x35\x98\x13\xb5\x26\x9f\x70\x6e\xd6\x07\xcf\x85\xd9\xb7\xb0\xa3\x67\xe4\x1d\x3d\xa3\xbe\xa3\x67\x04\x3d\x7a\xae\x6d\x7a\x24\x92\x09\x8e\xf1\x6c\x47\xd0\xf3\xe6\x9e\x44\x23\x00\x0f\x97\xaf\x54\x2a\xa7\x0a\xa3\x60\x80\x27\x8a\x81\xc1\xda\xc2\x7e\x7b\xdd\x03\x2d\x78\x9e\xfc\x57\x58\x4d\x5c\x51\x96\xea\x99\xfd\x54\xd2\xf4\xe7\x28\x34\x70\xfe\x9a\xf6\x49\x62\xf7\xce\xc8\xcf\xe9\xec\x80\x71\x07\x96\x09\x5e\xc5\xa3\x41\x33\xd6\x1c\x5a\x92\x71\x19\xaf\x85\x46\xef\xbd\x74\xb5\xde\xc3\xf8\xed\xb5\x5b\xa9\x23\xbb\xf1\x5b\x81\x72\x1b\xd8\xae\x2e\xed\xb5\x5f\xe9\xbf\xdd\x43\x9a\x97\xd8\x16\x26\xed\x1b\x67\x06\xc2\xb8\xb1\x17\x61\xd9\x45\xb7\x2c\x5b\xa1\x94\x28\x9c\x2c\x48\x5a\xf3\xce\xba\xaa\x8e\x40\x00\x49\xd0\x74\xdb\x14\x93\x9c\x33\x9b\x85\x1e\xb3\x2a\xd0\x0c\xdb\x1a\x87\x17\x51\x30\x40\x1d\xe4\xd8\x85\x95\x66\x5f\xae\x9d\xb6\x94\xa8\xc8\xb4\xec\xc1\xc8\x37\x5a\xfa\x50\xb5\x3a\xfa\x31\xb0\x40\x5b\xd7\xce\x72\x0f\x95\xeb\xa7\xd5\x76\xcd\x6f\x5f\x20\x4f\x31\x85\x37\xc8\xda\x8c\xe2\x55\x85\x9f\x4b\xdb\xa5\x2a\x7e\x0c\x1c\x50\x33\x95\x75\xe8\xd1\x6c\x37\x94\x73\x46\x15\x17\x34\x74\xa9\x8f\xb3\xac\xf5\x77\xb7\x7d\x24\xc2\xa2\xbe\x76\x7a\x7b\x15\x25\xd3\x9e\xff\x21\xe2\x42\x3b\xea\xee\x71\xd9\x95\x94\xe3\xd3\x0d\x54\x5b\x6d\xba\x79\xb1\x8d\xec\xb7\x57\x9f\xee\xb7\xd7\x16\xe1\x66\x94\x82\xa4\xd1\x74\x83\x53\x7b\x5b\x4a\xa7\x20\x29\x2a\x19\x79\x29\xcc\xae\xcc\x56\x3e\xe9\x90\x87\x7b\x68\x82\xa6\x7c\xda\xa2\x06\xb2\xc7\x66\xf2\xb3\x2e\x3a\xa0\x04\xc9\x08\x96\xbb\x41\xd5\xb6\xb1\x89\x61\x6e\x49\xb1\x1c\x7b\x36\xf1\xc7\x3e\x9b\x78\xec\xdb\xc4\x5d\x9e\x92\x09\xa0\xba\xbb\xd1\x14\xd7\x37\xba\xbf\x44\xf0\x6b\xec\x5b\xc8\x10\x54\x73\x23\xbb\x17\x2c\xd4\x6b\xb2\xb0\x1d\x21\xf9\x5d\x00\xa1\x9e\x93\x05\x5c\xcf\xe1\x8b\x43\x84\x56\x82\x5c\xec\x1f\x20\x19\x83\x6d\x68\xb5\x20\xad\x2c\x3d\x29\xd3\x26\x28\x39\x25\xea\x99\x10\x86\xfe\xfc\xcb\x44\xbf\xfe\xfc\x2b\x8a\x10\xa8\x95\xcd\x96\x79\x75\xb6\xec\xe5\x80\x8f\xc1\x96\xf6\x44\xba\x1c\x91\x6a\xc2\x73\x13\x5c\x14\xf6\x3e\xe5\xcf\xbf\xd0\x74\xa5\x48\xdc\x52\x43\x8b\x45\x26\x6d\x7e\xd7\x77\x4e\x9a\xa8\x38\x5c\x68\xcd\xc8\xcd\xe7\xfb\xb7\x57\x73\x31\x1e\xcd\x66\xdf\x66\xee\xc6\x72\xd7\xe1\xf1\x58\x50\xa5\x64\x6e\x5f\x8e\x14\xcf\x88\xc0\x2c\x21\x46\xb3\xa2\x3d\xe7\x09\xd5\x4c\xff\x12\x9c\xcd\x3d\x0a\x72\xa2\x16\x3c\x45\x6a\x55\xc4\x1c\x5f\x63\xdf\xa8\xee\x40\x3f\xf8\xf3\x2f\x74\x87\x85\xa2\xe6\xca\xa5\x2e\x78\x32\xd3\xf6\xeb\x7d\x21\xc8\x50\x75\xd5\x20\x73\x66\x6e\xe3\xf6\x01\x85\xaa\x2d\xc3\xec\xb7\x57\xab\x9c\xc9\xd2\x54\x96\x46\x29\x4a\x70\xe9\x9a\xbb\xc4\xb7\x77\xbc\x38\x43\x38\x4d\x05\x91\x72\x0f\xc1\x82\x2a\x09\x2b\x58\x5a\x82\xcc\xfd\x18\xb6\x07\x52\x87\x43\xfc\xb0\x20\xf6\xd1\x9f\xf4\xb3\xd3\x72\x36\xd3\x9a\xdd\xde\xac\xa5\x58\xe1\x23\xa9\xb8\xc0\x73\xf2\x33\x92\x05\x49\xe8\x8c\x92\x14\x4d\x57\x46\xf7\xb4\x07\xae\x6f\xf4\x70\xa2\x4a\x9c\x55\xbf\x35\x23\x1b\x83\xac\xca\x18\x0d\xdd\xe5\x35\x77\xc9\xf6\xf9\xae\x9c\xbe\xb1\x6f\x9b\x86\x5d\x2a\x73\x19\x58\x99\x66\xa8\x71\x1c\x63\xd8\x0e\xcd\x29\x66\x5a\x71\xe4\x98\x6a\xa7\x01\xc9\xfd\xac\xd0\x31\x34\xf7\x97\x71\xb7\x0a\xc6\xcc\x30\xd6\x65\x8a\x0e\x14\x16\x73\xa2\xcc\x1f\xa2\xc0\x81\xb7\x09\x8c\x90\xd4\xdd\x6a\xce\xa8\x90\xa6\xc4\x4a\x8b\x8e\x8d\xa3\xec\x31\x7b\x68\x99\x9a\x27\x84\x0b\x2c\xd1\x54\x9b\x20\x91\xc5\x5e\x83\xb1\x1f\xfd\xdf\x15\x7a\x33\x8a\xb4\x0b\x38\x70\xd5\x6d\x3d\xa3\x49\xdc\x49\x53\x8f\x14\x7f\x8f\x7b\xc4\xc6\xa5\x7f\xfb\xf4\x02\x05\xa5\xa1\xd7\x89\xc6\x76\x92\xb5\x5a\x6a\x13\xbb\x5e\x3a\x53\xe5\xba\x9b\x7c\xdb\xf0\x66\xeb\xa1\x1a\x1a\x00\x6a\x11\x68\x63\x03\x49\x46\xb0\x88\x5c\x5a\xe8\x41\xb5\xbe\xb6\x5b\x77\x93\xa3\x32\x94\x83\xd0\x43\x0c\xf4\x0c\xdb\x95\x98\x92\x3d\x31\xfe\xcc\xd6\x56\xee\xd9\x9c\x46\xe1\x92\x0f\x9f\x32\xa8\x2d\xbe\x4e\x59\x9f\xe0\x37\x52\xe3\xa8\x6a\xb9\xa0\x2e\x85\x66\x5d\x53\x76\x53\x17\xc8\x04\xee\x91\x1c\xa9\xb0\x2a\x25\x2a\x8b\x34\xb2\x94\xf1\x83\xef\xd1\x9f\xdf\x5f\x3f\x5c\x9f\x9f\x7d\x0d\x78\x06\xff\x9e\x3c\x5c\xde\x98\x6c\xa0\xbb\xcb\x40\xea\xde\x4c\xdb\x99\xa8\xb1\x33\xab\x02\x2a\x94\x96\xc4\x26\x13\x1b\xe6\x74\x27\x51\x0a\x52\x64\xd8\x3c\xa1\x07\xaa\xb7\x85\x39\xab\x8f\xaa\x34\x0b\xc7\xc8\x3a\xce\x71\x4f\xa6\x25\xcd\xd2\x2d\xb7\x7f\x27\x5e\xa4\xe3\xa4\x2f\xd2\x71\x02\xbd\xfd\x73\xe8\x94\xcd\x2f\xc8\xb4\x9c\x67\x7c\xee\x0f\x05\x0c\x7d\x37\x43\x55\x31\x2c\x7f\x28\x60\x3c\xbb\x35\xd4\x5a\x61\x7a\x6b\xa4\xde\x42\x60\xe4\x0f\x29\xec\x90\xa8\x2a\xfc\xd5\x8b\xc5\x85\x57\xde\x77\x02\xef\x84\xb2\x3e\x5e\x84\xf0\x9e\x40\x5b\xa0\x38\xb3\x32\x31\xd5\x3b\x85\xd9\x9e\x19\x4f\x9e\xa2\x76\xcc\x89\xef\xcb\x83\x40\xdf\x5e\xa9\x44\x25\x8b\xc7\x85\x5e\x08\x55\x8c\x4d\x78\x5e\x98\xe4\x7d\x97\x3f\x33\x2b\xb3\x40\xd4\x14\x02\x0c\xf5\x74\x0f\x2a\x68\x41\x64\x99\xa9\x86\x14\xb7\xe3\x63\xdc\xaf\x13\x68\xa6\x4b\x27\xb8\xdb\x4b\x71\xe0\x50\x33\x95\xd7\x70\xd6\x30\x0e\x9f\x13\x6b\xce\x82\xcd\xc3\xb6\x29\x97\x19\xb7\xa5\x21\x98\xad\x50\x51\x79\xcc\x10\xfa\x80\x76\x06\xe3\xd6\x92\x56\x35\x9d\x3e\x81\xb6\xb7\x80\x23\x29\x25\x73\x81\x83\xb5\xaf\x3e\x0d\x50\xb3\xe3\xaa\x3e\x4f\xed\xed\x89\xa3\xe4\xb0\x4a\x02\x31\xbd\xb4\x78\xb1\x91\x68\x0a\x0b\x76\x9f\x73\x36\xa3\x73\x58\x22\xc8\xa9\x77\x14\xf4\x75\x32\xab\x1f\xdf\xca\xe4\xc4\xd0\x80\x66\xd4\xf4\xe9\xc1\x29\x4a\x39\x8b\xf1\x4d\x4e\xa1\x19\x21\x73\xa2\x5c\x36\x74\x34\x12\xf0\x68\xb2\x61\x06\x8b\x15\x9f\x8e\x77\x0a\x3d\xbe\xda\x70\x91\x39\x56\xa7\x81\xd4\xbf\xae\x3c\xdb\xaa\x29\x97\x5b\x40\x13\x20\x88\x0d\xa1\x9c\x06\xd2\x00\xb7\xe1\x3e\x91\x55\x3c\x5e\xe0\x2e\x2d\x5c\x98\xd0\x16\x4f\x93\x37\x12\xcb\xd9\x40\x97\x9c\xad\x88\xae\x82\x31\x7e\x96\xd0\x72\x93\x35\xcc\xbd\xa2\xac\xa7\x81\x9e\x39\xdb\x39\xbb\x67\xf6\xdc\xe9\x00\x5a\x5f\xf5\x9b\xe4\x0c\xa5\x3c\xa9\x34\x36\x9f\xfe\x46\x92\xc0\xb9\x53\x09\x59\xa3\x2c\xfe\x66\xf7\x97\x29\xb6\x33\xbf\xb0\x45\x07\xeb\x82\x71\xd8\x0a\x51\xa3\xbf\x85\xa7\x15\xd4\xc4\x37\xb8\x28\xb6\xf5\xe3\x19\x7a\xad\x5d\x86\xc7\x1b\x2a\xb8\xb9\xc6\xb4\x39\x97\x7c\xdb\x90\x03\x6f\xc8\xcd\xcc\xf2\x7a\xc8\x2f\xe7\xfd\x27\xc4\xd0\x4b\x15\x1c\xf6\xa5\x0a\x0e\xc1\xa9\x82\x5f\xce\x91\x12\x74\x3e\x27\x51\x1e\xfe\x10\x9c\x2a\xf8\xe5\x3c\xbe\x2b\xe5\x10\x9c\x2e\xf8\xe5\xfc\xed\x35\xee\xe0\x19\x82\x73\x05\xbf\x9c\xd7\x4d\x21\x22\x33\x9c\x86\xe0\x4c\xaa\xef\x34\x51\x34\xaf\xcd\xf5\x84\x33\xa9\x44\x99\xa8\x98\x8d\x3c\x04\xe7\x55\x7d\xe5\x58\xdb\xac\x4b\x22\x24\x41\x39\x8e\x48\xae\x1a\x82\x93\xab\x0c\x96\x3d\x5c\x4d\xd9\xf7\xae\xdd\x9b\x6e\x88\xc2\x57\x93\xb5\xcd\x33\xde\xdc\x3c\x23\x6f\x6f\x8f\xfa\x36\xcf\x28\xd0\xb6\xa9\x23\x1b\xe2\xc6\xfe\xa3\xb9\x75\x48\xb5\xc7\x3d\x0f\xa5\xb1\x6c\xe5\xd9\x28\x50\x87\x13\xe6\x99\x87\xaa\x67\xa1\x95\x61\xdb\x42\x87\xe3\x0e\x02\xf5\x3f\x41\x97\xe2\xba\xee\xba\x8a\x6e\xae\x26\x6b\xb5\xa5\xbb\xe1\x01\xef\x16\x6e\x4c\x90\x7d\x3f\x28\xa0\xab\xf6\xe8\x0a\xe9\xf7\x03\x03\xba\x66\x8d\xd7\x41\xd8\xef\x25\x29\x89\x69\x34\x9b\xcb\x40\xee\x53\xe7\xa3\xfa\x57\x37\x57\x93\x43\x5b\x36\xe6\x1c\x29\xf2\x52\x60\x96\xa2\x99\x20\xb6\xb5\xea\x3f\x21\x44\x03\xcb\x1a\xce\x79\x5e\xe0\x44\xf5\xb4\x87\x6c\x3f\x92\xf0\x32\x4b\xd9\xdf\xcd\x85\xbd\x56\xc8\x55\xbc\xcd\xe4\x86\x30\x53\xa8\x66\xa8\x94\x05\x06\x44\xaf\xa3\xa8\x8c\x21\xa1\x06\x0c\x54\x79\x75\xdd\xdb\x5f\x4d\xb4\xbd\x13\xdd\x7a\x6c\x14\xa8\xee\xea\x87\xda\xa3\xcf\xd9\x28\x50\xd5\xd5\x05\xb6\xa9\x64\x22\xed\xc7\x51\xa0\xac\xab\xab\xd1\xa3\xbb\x3b\x2e\xb0\xc0\x39\x69\x75\x07\xda\x0d\x0e\x7a\x23\xaf\x05\xc1\x04\xc4\xa3\x50\xa0\x21\x86\x2a\xe8\x1e\x8f\xb4\xfb\xb5\x45\xd3\x1e\x3b\x0a\x70\xf7\xdb\x88\xa9\x49\xcc\x2c\x23\x2c\x85\xd1\x10\x9c\x05\xc4\x38\xca\x49\x4a\xb1\x91\xc6\x9b\xab\x49\x14\x18\x34\xf5\xc7\xf5\xc7\x59\xeb\x40\x34\xa5\x31\xc7\x7a\xa0\xe2\x6f\x2b\x3f\xb5\xcf\x13\x05\x05\x55\x25\xb4\x6e\xec\xb0\x8f\xe5\x10\xa8\xf9\xdb\x3a\x35\x77\x67\xa2\x5d\xbb\x28\x44\xa8\x32\x69\x10\x67\xd4\x25\x69\x44\x63\x42\x35\x4a\x05\x52\x3b\xc0\x05\x89\x6a\xf8\x3d\x1a\x82\x23\x98\x3f\x72\x9a\x50\x45\xa3\x04\x6d\xc4\xb4\x99\x74\xac\xdd\x34\x04\xb7\xa9\x6e\x49\x91\xa9\xfb\xd9\x6b\xb6\x50\xb5\xd3\xe2\xb0\x8d\x38\xec\x85\x0a\xd5\x3f\x0d\xaa\xeb\x77\xb4\x0f\x2a\x38\xb9\xbd\x2d\xc0\xd5\xc1\x15\x7b\x3d\x34\x1a\x82\x7b\x57\xb7\x33\x25\xf0\x7e\x42\x1c\x88\xca\x6d\xc3\xb4\x59\x97\xae\x97\x6b\x34\x2e\x54\x2f\x3d\x0b\xae\x9a\x7e\x70\xb6\x07\x8c\x32\x57\x2f\x26\x6f\x40\xb4\xce\x1d\xdf\xc2\xd6\x27\x9f\xc0\x39\x4a\xa9\x7c\x82\x10\x05\x55\x5c\x0d\x33\x24\x21\x4f\xdd\x8c\xd8\x99\x00\xa8\x1a\x33\xeb\xfe\x1f\x62\x0a\x54\xcd\xfd\x47\x57\x0a\xaa\x02\x5d\xe2\x59\x87\x07\xb6\x2b\x2a\x54\x07\x56\x7d\xe6\x7e\x14\x2e\x54\x0b\x56\x26\xdf\xde\x88\xe0\x7a\x7a\xcf\x50\xc9\x89\xc2\xc8\x65\x83\xc6\x28\x06\x70\x61\xfd\xe6\xe1\xb6\x37\x30\xb8\xc2\xbe\x05\x54\x49\x7a\xb4\xb7\x02\x2e\xb5\x0f\x33\xba\xa3\x6d\x20\x04\x78\x77\x5d\x67\xcc\x97\x7d\x61\x77\x37\xd4\x5a\xcb\xbb\x07\x2e\xb8\x60\xa4\x09\x77\x98\x48\x47\xc3\xf1\x28\x58\xa8\x9e\xc2\x59\xce\xa5\x42\xb3\x32\xd0\x22\x15\x82\x03\xd5\x4c\x95\x95\x62\x98\x1a\x65\xa5\x0c\xa1\xca\xc8\xb4\x68\xc6\x0a\x67\x7c\x1e\x6a\x57\xb7\x03\xa4\xdf\xc5\xb4\xeb\xbb\x2c\x2e\xe8\x17\xfb\x29\x8f\xd1\xd0\x6f\x3f\xda\x2f\x9f\xe6\x68\x5b\x62\x41\x79\x29\xb5\x16\x90\x71\x8e\xe1\x68\x77\x33\x4c\xe2\x25\x41\x94\xf1\x34\x4a\x30\x47\xbb\xeb\x9d\x82\x17\xb6\x95\x3e\x46\x86\xcb\x51\xb0\xbb\x6b\x9d\xa2\x94\x0b\x93\x58\x13\x8d\x1a\xb8\xdc\xeb\xce\x90\x3c\x6f\x24\x76\x97\x56\x97\xf5\x05\xcb\x64\x25\x8f\x4c\xe5\xc8\xda\x1d\xcb\xc9\xe6\x15\x8b\x57\xb6\x39\x1a\x7f\xec\xb9\x62\x09\x94\x6d\x76\x5c\xb1\x58\x15\xf9\xab\x21\x21\x82\x57\x81\x02\xac\x0e\x20\x9b\x09\xf5\x2b\x4e\xe3\xb1\x7c\x3f\xab\xef\x1b\x2d\xbf\xba\x8f\xa6\xec\x82\x57\x2f\xcb\x57\x3e\xf7\xbf\x31\xb3\xb9\x26\x27\xde\x9a\xf4\x5d\x7b\x8d\x03\xad\x17\xc3\x89\xed\x34\x53\x44\x20\xb9\x62\x0a\x77\x75\x05\x05\x70\xeb\x04\x5e\x62\xe2\x10\xad\x6b\x6a\x43\x0e\x05\x56\x8b\x28\x54\xe0\x8d\x50\x93\x40\x9d\xf1\xf9\x91\x79\xd4\xfa\x47\x6a\x67\xe5\x5f\x2f\xda\x3f\xf9\x24\xfc\x71\xa0\x8d\x85\xfb\xe0\x6d\xa6\x0f\x7d\x0b\x17\xc8\x98\xee\xc8\xff\x24\xca\x55\x81\x15\x3c\xa3\xc9\x5a\xce\xf2\xf5\xed\x1d\x56\x8b\x23\xb6\xcc\x67\xfd\xa9\x08\x63\xef\x36\x75\xbc\xf9\xfd\xca\xcd\x41\x67\xc2\x34\xdb\xdd\x92\x10\x3d\xf6\x3e\x89\x34\xee\x4c\xc1\x70\x03\x4f\x71\xf2\xb4\x7d\x5c\x2f\x0f\x63\xbc\xf9\xd5\x9e\x7a\x5c\x9b\xf7\xb7\x39\xde\xc6\x80\xa7\x1e\x07\x4e\xfb\xd2\xf5\x4e\xc1\x9f\x01\x9a\xdc\x9d\x9d\x5f\xae\xff\x0d\x2c\xd9\xa7\x81\x4f\x00\x6d\x3d\x8b\x2a\xbf\x6e\xad\xaf\x67\xcb\xcb\x72\x6d\x42\x43\x5d\x47\x7b\xe8\xd8\xfd\xec\xdf\x46\xc7\x82\xce\x17\xe6\xe3\x8d\x94\x9b\x4e\x75\xbf\xf1\x29\x92\x65\xb2\x40\x58\x56\x39\x9b\x94\x99\x33\xae\xca\xdd\xa6\x6c\x1e\xaa\xb9\x53\x1b\x6d\x4d\x31\x7a\x5e\x34\x81\x8e\xd6\x1c\xa0\xe6\x44\xab\x37\x7e\xab\x5d\x29\x79\x21\x49\x69\x9b\xeb\x9b\xfa\xda\xfa\x63\xa9\x9d\xbd\x5b\xdd\x23\x95\x39\x5b\xd7\x55\x75\x77\x68\xdd\x9c\x8a\xc7\xc3\xa6\x25\xd9\xd1\xd9\x7c\xf3\x00\xd7\x52\x1d\xf8\x30\xed\xe6\x37\xc9\x7a\x05\x5b\xbf\x00\xb5\x7f\x1e\x59\x4a\x66\x94\x91\x74\x2d\x79\xac\x3d\x12\xd4\x91\xfa\x87\xe4\x0c\x3d\xac\x0a\xb2\x91\x86\xb6\x59\x58\x59\xdd\xe0\x7c\xe6\xe9\xca\x3c\xef\x63\xfa\x95\x2d\x9d\xb7\x41\xd5\x47\xf5\x6c\xf6\x99\x31\x1b\xef\x6e\x27\xde\x90\xc7\xe0\xcb\x9e\x47\x86\x4b\xb5\xe0\x82\xfe\x41\x52\xf4\x28\x49\xf7\x44\xce\xdc\x73\xb6\x54\xf9\xbf\x09\x4e\x89\xcf\xbf\x63\x70\x8c\xd7\xf0\xc3\x30\xb1\x9f\x7f\xe6\xb9\x3b\xbc\x32\xb5\xca\x36\xf4\xed\xa3\x82\xe3\x2a\xff\x73\xe4\x76\xc9\xd1\x75\xea\xe6\xb0\x05\x7f\xed\x8d\x9f\xca\x92\xa6\x3f\xa3\xef\x38\x2b\xbd\x75\x0c\x7d\x21\xb9\xab\x91\x80\xfb\x1e\x63\x47\x6f\x5f\xf7\xbd\x5d\xcd\x66\xf7\xbb\xba\xbb\xc3\xdd\xed\xc4\x54\xfd\xd8\xf3\xdb\xb4\x63\x9a\x28\x9c\x04\x62\x4f\x48\x10\x9b\x4c\x1e\x90\x8f\xd0\x27\x6b\xfb\xc4\xfc\x32\x4c\xe7\x0d\xce\x5c\x37\xe4\x7f\xc8\x50\x52\xbe\x96\x27\x97\x54\x65\x86\x09\xaf\xde\x0e\xdf\xb3\xa4\x9f\xf9\x0c\x9d\x37\xdc\xe9\xa0\xeb\xff\x9c\x7f\xa1\x2f\x97\x75\x90\xcc\x99\xf9\x96\x08\xfd\xcc\xaf\xcc\x67\x78\x5d\x0d\x5c\x4a\xf0\x66\xe5\xd1\x71\xe8\xa3\x50\x1d\xa3\x26\xb6\x25\x3a\x57\xf6\xba\xcd\xd6\x57\xd0\xe6\xc3\x18\xf5\x98\x43\xb8\x50\xde\x57\x5f\x43\x56\x34\x27\xbc\x54\x6b\xe9\x6d\x30\xd5\x3d\xf0\x55\xf7\xa0\x57\x75\x0f\xde\x83\xbb\xa7\x5e\xf4\x77\x4e\x6d\x57\x1a\xea\x27\x5d\x00\xb1\x10\x64\xed\xab\xce\xcd\xf2\xda\x9e\xaf\xd4\xb6\xc6\xae\x5e\x69\x06\x31\xb5\xfb\xac\xf9\x40\x19\x61\xaa\xf9\x00\x58\x4d\xfe\x0e\x7b\xfe\x9f\x25\x11\xab\xb8\x4d\x7f\xcd\x66\x59\xf9\x72\xf1\xd9\x48\xae\x59\x88\x1e\x71\xad\x1e\xf6\x49\x05\xcb\xec\x03\xcd\x49\x93\xfd\xd2\x45\x70\x95\x26\xa3\xcc\xd3\x44\x50\x9e\x36\x5d\x0a\x02\x04\x96\x4e\xb4\xec\xa9\x75\x9c\x1f\x8e\xf3\xc3\x63\xfd\xff\xe2\xf0\xc3\xe2\xf0\x78\xb0\x38\x1c\x8c\x16\x87\x1f\xd3\xc3\xe1\xfb\xb4\x2d\x7b\xdf\xbe\xdf\x5c\x22\x9c\xe6\x94\x6d\xc9\x4d\xf6\x85\x6f\xf8\xbe\xef\x93\xf6\xcd\x0b\xdb\x59\x62\x88\x48\x54\xd6\xc9\x8d\x73\xcc\xfe\x6e\x53\xc9\xd9\x32\x27\x28\x51\x22\xdb\x3c\x1d\x35\x1e\xd4\x36\xfe\xca\xe7\xa8\xc0\xf3\xae\xba\x80\x16\x5c\xe6\x9e\x0c\x80\x41\x0d\xe0\x73\x5b\x0f\x57\x77\x20\xe9\x82\x73\x2d\xe5\xab\xfa\x39\xca\x99\xf4\x16\xea\xec\xff\x79\xa1\x7e\xf0\x3a\xfc\x58\x3e\x77\x72\xf1\xdd\xff\x06\x00\x00\xff\xff\xf2\x39\xb5\x86\xa6\x82\x00\x00")

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

	info := bindataFileInfo{name: "../resources/events.yaml", size: 33446, mode: os.FileMode(436), modTime: time.Unix(1600414621, 0)}
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
