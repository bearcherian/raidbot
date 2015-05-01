package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _www_app_jsx_swp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9a\x5d\x68\x1c\x55\x1b\xc7\x4f\xde\xb7\x7d\xdf\x6e\x3f\x5f\xfa\x42\x95\x56\xf0\x64\x62\xe9\x46\x9a\x9d\x24\xad\x56\x93\xdd\x25\x6d\x54\x88\x36\x52\xfa\x25\x12\x02\x3d\x3b\x73\xb2\x3b\xcd\xec\xcc\x30\x73\x26\x69\x88\x6b\x15\x6f\x8a\x88\xa0\x37\xe2\x45\x6f\xbc\x11\xbd\xf5\xeb\xca\x2a\xe8\x95\x82\x82\xc5\x5e\xaa\x08\x82\x37\xbd\xd1\x2a\x45\x85\xfa\x3f\x67\x66\x76\x27\x9b\x35\xb1\x0a\x16\xf1\x3c\xf0\xdb\x3d\x33\xe7\x39\xe7\x3c\xe7\xff\x9c\x8f\x0d\x6d\x6d\xf8\xf4\xd4\x34\x3d\x54\x3a\x48\x60\xff\x23\xe4\xc6\xf4\xc4\xa9\x0b\x3b\xce\xf5\x5d\xed\x27\x84\x05\xfe\x3c\x73\x97\x02\xe1\xcc\x93\xb5\xcc\xe3\x8b\xa5\x9c\x6f\xc9\xf2\x9b\x3d\xfd\x9e\xcc\x39\x99\x21\x73\xec\x9a\x2f\xcc\xc5\xc5\x45\x93\x05\x41\xe9\x6c\x74\x6e\xcd\x41\xb4\x69\xd3\xf6\x27\x6d\x1b\x39\x30\x3a\x32\x2c\x4b\x03\x46\x3f\xfd\xff\xce\x53\xb7\x3a\x20\x6d\xda\xb4\x69\xd3\xa6\x4d\x9b\x36\x6d\xda\xb4\xfd\x85\x26\x82\x3e\x72\x1e\xdf\xff\x4a\x9f\x5f\x48\xbf\xfb\xba\xbe\xb5\x69\xd3\xa6\x4d\x9b\x36\x6d\xda\xb4\x69\xd3\xa6\x4d\x9b\x36\x6d\x7f\x5f\x63\x36\x21\x0b\xf8\xe3\xff\xca\x06\xa2\xfe\xfd\x3f\xfb\xfb\xff\xd2\x0e\x42\xde\x05\x0b\x20\x06\x27\x40\x05\xec\x06\x57\xb7\x13\xf2\x19\xf8\x04\x7c\x08\x2e\x82\x87\xc1\x04\x18\x07\x87\xc0\x01\x30\x0a\xb6\x82\x1f\xb7\x11\xf2\x15\x78\x05\x5c\x04\x17\x80\x05\x1e\x01\xf7\x82\x11\x30\x04\xb6\x83\xcb\x5b\x09\x79\x16\x3c\x03\xce\x83\x18\x3c\x06\x0e\x03\x13\xec\x05\x14\xdc\x01\x6e\x07\xb7\x81\x6b\x5b\x08\xf9\x02\x7c\x0c\xde\x03\xcf\x83\xa7\xc1\x13\x20\x04\x27\xc1\x61\x30\x0a\x8a\x60\x00\xdc\x09\xf6\x80\xdd\xe0\x87\xcd\x84\x7c\x09\x3e\x02\x97\x40\x08\x38\x38\x03\x4e\x83\xfb\xc0\x00\xd8\x09\x36\x81\x7f\x83\x1b\x05\x42\x7e\x06\x3f\x81\x2b\xe0\x7d\xf0\x36\x78\x09\x3c\x07\x1e\x00\xf7\x83\x5d\x60\x33\x20\xd2\x7f\x13\x21\xd7\xc1\x35\xf0\x3d\xb8\x0c\xde\x02\xaf\x83\x97\xc1\x53\x40\x80\x00\xb8\xe0\x34\x38\x08\x76\x80\x02\xd8\x08\xfa\x00\x01\x5f\xff\x17\xda\x83\x77\xc0\x1b\xe0\x45\xb0\x04\x66\xc1\xe3\x60\x12\xec\x06\xd7\xff\x83\xf1\xc0\x37\xe0\x53\xf0\x26\x78\x15\x5c\x00\x75\x70\x06\x9c\x00\x47\xc0\x08\xb8\x1b\xdc\x05\xb6\x80\x5f\x36\x22\x77\xe0\x3b\xf0\x2d\xf8\x00\xbc\x06\x9a\xe0\x18\x18\x06\x7b\xc1\x1e\xb0\x0b\xec\x04\xdb\xc1\xb6\x8d\xc9\x7a\xfa\x7c\xc3\xad\x5b\xd7\xda\xfe\xc9\x76\x9c\x33\x4b\x94\x42\xee\xd9\x3c\x2c\x96\x0f\x07\x01\x35\xab\xfb\xa9\xed\x5b\x71\x93\x7b\xa2\x54\xe7\xe2\x41\x97\xcb\xe2\x91\xa5\x29\xbb\xb8\x8f\x05\xc1\xbe\xc1\xc1\x71\x42\x5a\xf8\x28\xb4\xf6\x93\x42\x41\x16\x0a\x85\xb2\x69\x3b\x0b\x55\x59\xca\x17\x0b\xe5\x69\xde\xac\xf1\xf0\xa8\x13\x09\x6a\x33\xc1\x2a\xcb\xf2\xb3\x65\x56\x0b\x69\xfd\x71\xe6\xd8\xab\x6b\xd3\xca\xc9\x06\xf3\x3c\xee\x76\xd7\xd3\x88\xbb\xdc\x12\x95\x65\xd1\x70\xa2\x52\xf2\x90\xba\x66\x6d\xcb\x08\x81\x5a\x2e\x8b\xa2\x47\x59\x93\x57\x8c\xd0\x5f\x34\x54\x4d\x77\x85\xe5\x7b\x82\x39\x1e\x0f\x87\xe6\xdc\xd8\xb1\x95\x53\xc8\x45\x1c\x7a\x45\x94\x5a\xb2\x89\x9a\x65\xa1\x60\x58\x96\x65\x8c\xd1\xe5\xe4\xd1\xb0\x79\x24\x1c\x6f\x69\x48\xfe\xef\xac\xa1\x05\xbf\x2e\xab\x72\xde\xe9\x57\xa1\xa9\xe6\x1f\x8d\xcd\xcc\xa6\x2f\x3c\x0c\x3b\x66\xcc\xf9\xbe\x91\x76\x5b\xab\xd5\xb2\xb6\xed\x56\x46\xda\x0c\x15\x33\x18\xa9\xe9\x88\xd0\xf1\xe3\x68\xde\xc8\x7a\x31\x64\x37\xa8\x35\xac\x06\xe7\x11\x1f\x12\x0c\xd1\x58\xfd\x59\x9f\x8c\xb1\xac\xcf\x95\x81\x5a\xa1\x2f\x3a\x55\x13\xc8\xbd\xb3\x20\xfb\x49\x0a\xb2\xf5\x02\x0b\x95\xd6\xb4\xa2\x9c\x94\x02\xa9\x20\x54\x8a\x57\x2d\x1f\xf5\xeb\x8e\x67\x56\x93\x34\x53\x95\x7e\x67\x8e\x16\xd3\x3e\x4a\x2c\x16\x0d\xac\x17\xc7\x62\x82\xdb\xb4\x52\xa1\x73\xcc\x8d\x38\x1d\x94\xdd\x25\xeb\x6c\x8c\xce\xc5\x1e\x9c\x7d\xaf\xa8\xde\xaa\x59\xa7\xb9\x14\x27\x04\xda\x15\x93\x27\x59\x1c\x44\x5d\xda\xb3\x95\xe4\x18\x91\xc9\xd9\x63\xdc\x15\xa9\xcf\xf5\x2a\xab\x3b\x3d\x27\xd1\x23\x71\x68\x81\xf5\x3c\x85\xd8\x98\xab\x86\xe9\x0a\x44\x4e\x5d\xee\x80\x0a\x4d\x76\x85\x15\x72\x38\x4d\xca\xc5\x52\x5c\xce\x2f\x79\xd3\xcc\xc9\x42\x8b\x39\x59\xa6\x3c\x47\x64\xd2\x28\x65\x4c\x53\x69\xd3\x99\x0f\x3e\x79\x20\x65\x19\x5e\x57\x12\x35\x88\x74\x1f\xa3\xc3\x2a\xb1\x2b\xa4\x1d\x4b\x84\x55\x15\x71\xc4\x43\xb5\xb0\xa8\x61\xe4\xa7\xfc\x3b\x26\xac\xc2\x5e\x7b\xca\xbd\x37\xf9\x9c\x1f\x36\xb3\x8d\xea\x78\x41\x2c\xa8\x58\x0a\xb0\xa3\x04\x3f\x27\x0c\xba\xc0\xdc\x18\x0f\x46\x7b\x2f\xd7\xc2\x76\x31\x12\xa1\xef\xd5\xab\xc7\x5c\xce\xb0\x30\x30\x21\x1e\xd2\x25\x3f\x0e\x69\xe4\x32\x6b\x9e\x66\x93\xa1\x7e\x48\x79\x93\x39\x2e\x65\xb6\x1d\xf2\x28\x2a\x9b\x69\xcb\x24\x82\x76\x00\xab\xf7\xb4\x3b\xd4\xb4\x87\x46\x46\x57\xee\xe6\x9e\x4a\xb7\x25\x90\x99\x5b\x3f\xf3\xab\x74\x58\x6e\xb6\x0f\xb8\x56\x12\x56\x63\xb4\x9a\x1c\x7a\x88\x17\xe5\x35\x02\x3c\xd8\x75\xda\x64\x9d\xe7\x67\x99\x8a\x94\xac\x74\xca\xa8\xdc\xc5\x54\xf8\x78\xc1\xb1\xa8\x38\x4d\x86\xa7\xae\x3c\x22\x99\x67\xd3\x1a\xa7\xac\xe6\x72\xe9\x72\xd6\x47\x62\xa1\x61\xc0\x42\x41\x72\xba\xa7\x9b\xbc\x13\x38\xa6\xbd\x96\x3c\xd3\x79\xc7\x9b\xd6\x27\x4c\x0f\xf8\x8e\x3a\xf2\xc8\x5f\x57\x9b\x7b\xfe\x80\x36\xd9\xf9\x90\x93\x47\xc9\x25\xc5\xe9\x21\x40\x16\xd9\x3a\xd3\x3f\xde\x71\xbb\xe9\xc9\x5b\x9d\x0b\xac\x33\xff\xf4\xbc\x5a\x57\x82\x03\x3d\x2f\xa3\x64\x90\xf6\xd9\x43\xb3\x4b\x52\x9d\x89\x95\x6c\x44\xd9\x51\xd7\x3d\x19\x84\x7e\x90\xdd\x96\xb8\x26\x93\x83\x5b\x09\x91\x8b\x12\x93\x4c\x9f\xa2\x52\x93\x05\xc5\x4c\x0d\x5a\xcc\x75\xac\x4e\xa8\xe4\xcc\x6f\x3b\xbb\xdc\xab\x8b\x06\xad\x66\xc7\x5a\x1a\xad\x8a\xb5\xed\x14\xc4\x51\xa3\x68\xc9\x13\x3d\x6d\x3d\x33\x3c\x4b\xfb\x2b\xd4\x98\x30\xd2\x56\xd8\xd3\x78\xaf\x82\xa2\x58\xbd\xb9\xc0\xd5\x9d\x94\x38\xe5\x62\x8e\x10\xf0\xcc\xec\x78\x47\xfc\x34\xc5\x27\x1b\x3c\xc4\x4e\x00\x56\x1c\x22\xb5\xc2\x5d\xa2\x9e\xaf\x52\x1e\x61\x8f\x38\x5e\x9d\x36\x7c\x1c\xad\x76\x7e\x45\xad\xd6\x62\xad\x75\x31\xb9\xc2\xf3\x66\x97\x46\x99\xd1\x46\xc8\xe7\x2a\xc6\x80\x41\x7d\x6f\xd2\x75\xac\xf9\x1e\x79\x2a\xe6\xde\xa8\x4b\xad\x55\x5d\xee\x7a\xd5\x2a\x9b\xac\xd7\x2a\x5a\x56\x45\x1e\x95\xe4\x31\x50\x34\xa8\x81\xb6\xab\x97\x53\xe6\xa4\x32\x63\xa4\x3f\x07\x3a\x57\x7a\xd7\x58\xf2\xda\xea\xba\x8b\x73\x19\x49\xba\x92\x09\xa1\x86\x14\x3a\x15\xc8\xa0\x32\x3f\xeb\xc9\xf8\x5b\x12\xaa\x8d\xaa\x62\x7e\x48\xfd\x8c\xa8\x74\x7a\xa0\xea\x47\x05\x95\x37\xaa\x1f\x1c\x43\x90\xac\xce\x92\xae\xc7\xd3\x26\xc9\x0d\x39\x4e\x5b\x49\x3f\x49\xe8\xf2\x87\x0d\xed\x79\x93\x66\xc9\x1f\xa3\x5e\xec\xba\xfb\xd5\x6a\x49\xca\x14\xbf\x1d\x7e\x0d\x00\x00\xff\xff\x07\xbb\x1c\xdb\x00\x30\x00\x00")

func www_app_jsx_swp_bytes() ([]byte, error) {
	return bindata_read(
		_www_app_jsx_swp,
		"www/.app.jsx.swp",
	)
}

func www_app_jsx_swp() (*asset, error) {
	bytes, err := www_app_jsx_swp_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "www/.app.jsx.swp", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1430280010, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _www_app_js = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x51\x6f\xdb\x36\x10\x7e\x76\x7e\xc5\x95\x7b\xa8\x0c\xd8\x4e\xd7\x6d\x2f\x0e\x3c\x74\x0b\x36\x20\x40\x37\x14\xdd\xde\x82\x3c\xd0\x12\x6d\x71\xa1\x28\x81\xa4\xb2\x19\x81\xfe\xfb\xee\x48\xca\x92\x3c\xd9\x51\x9b\x3d\xa4\xb5\xc8\xbb\x8f\x77\xdf\x77\x3c\xde\x13\x37\xc0\x53\x27\x9f\x04\x6c\xe0\x19\x78\xed\x72\xa1\x9d\x4c\xb9\x13\xd9\x1a\x76\x5c\x59\xb1\x80\x34\xe7\x5a\x0b\xb5\x06\x5d\x2b\xb5\x00\xc3\x65\x16\x7e\x43\x73\x73\x75\x45\x10\x46\xb8\xda\xe8\x5f\xc9\x1c\x71\x76\xb5\x46\xc8\x52\x27\x20\x60\x8e\xa8\x62\x65\x5d\x59\x7d\x32\x65\xc5\xf7\xdc\x6f\xcc\x6f\xa2\x4b\x38\xe2\x06\x9a\x80\x73\x1b\x4e\x42\x8c\xcf\x02\xc3\x5a\xa5\x46\x60\x24\xb7\x8a\x5b\x9b\x3c\x67\xd2\x56\x8a\x1f\x7e\xe7\x85\x58\x03\x8b\xa6\x6c\x71\x35\x33\x42\x67\xc2\xac\xbb\x73\xf1\xd0\xab\xd9\x8c\x00\x53\x72\x15\x16\x01\xef\x81\x51\xe0\xad\x1b\x3c\xdc\xa0\x89\xdc\x41\x02\x2e\x97\x76\x55\x61\x74\x76\xa5\x11\x1b\x36\x9b\x48\xc9\x2a\x26\x0e\x01\x6f\x16\xc1\x56\x55\x6d\xf3\x84\x05\x1b\x36\x27\x9c\x06\xff\x42\x42\x09\x19\xf6\x83\xff\x45\x89\x02\x29\x4d\x58\x26\x9f\xd8\x02\x9e\x3d\x48\xc8\xa1\xc5\xfb\xab\x94\x3a\x61\xc0\xe6\xcd\x02\xc8\x7d\xdc\x9f\x93\x77\x6e\xc4\x0e\x93\xff\x06\x7f\x97\xfa\x56\xc9\xf4\x71\xdd\x8f\xdf\x0a\x25\x52\x97\x9c\x64\x44\xb8\xa7\x4b\x74\x0e\xfd\x43\xe1\x37\x8b\xab\x66\x7e\x33\x50\xe0\xa3\xb4\xee\x0b\x54\x20\xf3\x97\x94\x18\x00\x9f\xa7\xc9\x3a\x53\xea\x3d\xe6\x17\x8a\x8d\xfd\x99\x0b\x23\x80\xe3\x5f\x5a\x1b\xc4\x77\xea\x00\xba\xf4\x45\x68\x61\x2b\xa4\xde\x43\x5e\x5a\x2c\x57\xd6\xa6\xd3\x3f\xce\x2b\xef\xa5\xde\x95\x06\xb5\xf6\x5b\x20\x75\x9f\x8f\x8c\x3b\xde\x4a\xec\x0b\x22\xbd\x7f\xf7\x00\x6f\x36\xc0\x3e\xb0\x76\x7d\xd6\xe2\x05\xf1\x53\xcf\x5f\x13\x95\x0f\x4e\xad\x81\x12\x7a\xef\x72\xf8\x11\xde\xb5\xce\xff\x4d\xff\x68\x5c\xf0\x2a\x69\xc9\x82\x24\x2e\x13\xb7\xed\xb9\xf1\x9a\x24\x30\x46\x56\x64\x1f\xeb\x42\x87\x82\xea\xfc\x17\x10\x6a\x61\xa4\x3c\x9a\x39\x78\x9a\x66\xcd\x2b\x6a\x97\xa5\xa5\x5a\x16\xd9\xf2\x3b\x76\xb1\x68\xf3\xf7\x9d\x92\x31\x5a\xcb\xe6\xd1\xa3\xc7\xc9\xb9\x72\xfc\x8c\x32\x4f\xad\xc5\xd6\xf6\x85\x42\x34\x1d\xe4\xe4\x2a\xf4\xe1\xb2\x4f\x4a\x70\x6c\x70\x81\x45\xe0\x2d\xdd\xe0\x4a\x5c\x13\x48\xb4\xf0\xe0\xa0\x28\x88\x7e\x46\xaf\xe0\xf7\x87\xe9\xfc\x52\xfe\x1d\xb9\x6d\x9a\xe7\x98\xfd\x4d\x14\x5b\x61\xa6\x72\xdb\x59\xbf\xc0\x6e\xd1\x87\x7d\x25\xbf\x9e\xcc\x1e\xb9\x01\xdb\xd3\x0b\x5c\x67\x78\xfb\x81\x6f\x95\x20\x13\x6a\xa1\x80\x37\xbc\xe2\xe6\x7f\xa3\xfe\xfb\xe9\xd4\x07\x7a\x3a\xf2\x3b\x16\xce\xd1\xff\xb1\xdc\x4b\x7d\xa7\xe5\x24\xf6\x8f\xc6\x17\xc8\xff\xfa\x3c\xbf\x7d\x7f\x39\x51\xec\x9c\xc5\x50\xa9\x69\x9d\x3b\xca\x89\x7b\x28\xda\xa1\xac\x0d\x58\xc5\xd3\x47\xa8\xad\x30\xfe\xa1\x45\xbd\x44\xc1\xa5\x02\x9e\x65\x46\xd8\x8e\xbe\x71\xfc\xad\x89\xd8\x97\xcd\xa4\xae\x6a\x47\xa9\xba\x43\x45\x59\x3a\xf1\x0f\x7d\x3e\x71\x55\xd3\x27\x6b\x7c\xef\x9e\x0d\x5f\xc0\x13\x61\x26\x8b\x42\x82\xec\x85\xbb\xc3\x99\x89\xab\x3f\x1c\xda\x9e\x11\x26\xb4\xf3\x36\x75\x8a\x63\x41\x0b\xa3\x03\x17\x6d\xe0\x8b\x56\xad\xe1\xdd\x22\x34\xe8\xe6\xbc\xee\xd7\xd7\xdd\x10\x63\x29\x80\x15\xb9\xd2\x10\x73\x7c\x7f\xda\x37\xe4\x42\x5d\x04\xc9\xc6\x0c\x8e\xb5\x17\xb9\x9f\xfb\x5b\x75\x7d\xdd\x0c\xeb\xf9\xa7\xaa\x9a\x42\x1a\x9a\x7d\x01\x65\x4d\xb8\x34\xb3\xd0\x11\x6e\xdb\x31\xf4\x68\xad\x8f\x0f\xe5\xc9\xc0\xb6\x01\xda\xa2\x40\x03\x2f\xc2\xf9\x83\x92\x8e\xa5\xf9\x45\x52\x3d\xa5\x11\x73\x20\x11\xf1\xea\x45\x1a\x72\x3b\xfe\x3c\x4f\xe3\xb6\xe5\x15\x8e\xaf\x31\xf1\xe9\x87\x92\x4d\x38\x83\x7d\x88\xb3\xe6\x3a\xc6\xe4\x2b\x84\x65\xc2\x3a\xa9\x0f\x4b\x6a\x93\xcb\xd4\x94\x8e\xa3\x41\x98\x1a\x18\xe7\xdd\xc7\x8c\x11\x19\x8c\x2e\x7c\x2e\x84\x15\x4b\xc7\xd1\x31\x7d\x13\x2a\x10\xb7\x8b\xd8\xbd\xd6\x70\x8f\xa0\x85\x74\x46\x96\xb5\x7d\x64\x0f\xc1\xa0\x09\xff\xb1\xed\x76\xdb\x61\xfa\x2a\xc6\xd6\x50\xb6\x28\x11\x64\x7d\x3f\xf0\x6a\x46\x42\x7d\x2a\xf7\x5d\xa0\x69\x9a\xd2\x47\xcf\xfa\xab\xe7\x11\xed\xb8\xd4\xc2\x2c\x77\xaa\x96\xd9\xe5\x96\x36\xe2\x6f\xca\xbf\x8f\x3e\xa3\x4e\xbd\x51\x17\x5d\x49\x9f\xb5\x57\xe9\x64\xd2\x1a\xd4\x6a\x73\xb1\x55\xb5\x03\xcb\x00\xee\xb2\x4b\xf7\x0e\x9f\x38\xcd\x46\x9a\x5a\xbc\x9e\x01\x27\x54\xfa\x68\x0f\xc0\x4b\x79\x6c\xac\x59\x99\xd6\xb4\xb8\xc2\x2b\x1a\xf7\x7f\x3e\xdc\x65\xc9\x5b\x5e\x55\x6f\xe9\xee\xff\x1b\x00\x00\xff\xff\xf5\x0d\xc3\x38\x32\x0e\x00\x00")

func www_app_js_bytes() ([]byte, error) {
	return bindata_read(
		_www_app_js,
		"www/app.js",
	)
}

func www_app_js() (*asset, error) {
	bytes, err := www_app_js_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "www/app.js", size: 3634, mode: os.FileMode(420), modTime: time.Unix(1430277631, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _www_app_jsx = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x56\x4b\x6f\xe3\x36\x10\x3e\x2b\xbf\x62\x96\x3d\xac\x0c\xd8\x56\xba\x6d\x2f\xb6\x6c\x6c\x1b\xb4\x40\x80\x6d\xb1\xd8\xf6\x16\xe4\x40\x4b\xb4\xc5\x46\x26\x05\x92\x4a\x6b\x18\xfa\xef\x1d\xbe\xf4\x70\xf3\x68\xb3\x07\x27\x12\x39\xf3\xf1\x9b\x6f\x1e\xd4\x23\x55\x40\x0b\xc3\x1f\x19\x6c\xe0\x0c\xb4\x35\x15\x13\x86\x17\xd4\xb0\x72\x05\x7b\x5a\x6b\x36\x87\xa2\xa2\x42\xb0\x7a\x05\xa2\xad\xeb\x39\x28\xca\x4b\xff\x0c\xdd\xfa\xea\xca\x42\x28\x66\x5a\x25\x7e\xb1\xe6\x88\xb3\x6f\x05\x42\x4a\x91\x02\x83\x19\xa2\xb2\xa5\x36\xb2\xf9\xac\x64\x43\x0f\xd4\x6d\xcc\xd6\xc1\xc5\x1f\xb1\x86\xce\xe3\xdc\xf8\x93\x10\xe3\x0b\x43\x5a\xcb\x42\x31\x64\x72\x53\x53\xad\xd3\xf3\x55\xa2\x98\x28\x99\x5a\x0d\x07\x20\xfa\x55\x92\x58\xcf\xc2\xda\x30\x8d\x9e\x77\x40\x2c\xc3\x00\x45\xe0\x7e\x8d\x26\x7c\x0f\x29\x98\x8a\xeb\x65\x83\x34\xf4\x52\xd0\x23\x32\xdd\x84\xd8\x97\x21\x42\xf0\x78\x49\x00\x5b\x36\xad\xae\x52\xe2\x6d\xc8\xcc\xe2\x74\xf8\xf3\xcc\x53\x6b\x98\x97\xfc\xd1\x1f\xfd\x1b\x02\x6e\xce\xd1\xf1\x4f\xc9\x45\x4a\x80\xcc\xba\xad\x35\x4b\x72\x0a\x95\x62\xfb\x0d\xf9\x86\x80\x14\x37\x35\x2f\x1e\x36\xe7\x11\x1d\xcd\x6a\x56\x98\xf4\x82\x20\x7a\x9f\x2f\x96\xba\x3c\xa3\x0e\x32\xcf\xf0\x68\xfb\x64\x69\x75\xf3\xab\x6e\xb6\x9e\x48\xf8\x89\x6b\xf3\x16\x19\x27\xde\x3e\x46\x6d\x94\x14\x87\xed\x1f\x15\x53\x0c\x28\xfe\x8a\x56\x21\x84\xa9\x4f\x20\xa4\xab\x06\x0d\x3b\xc6\xc5\x01\x2a\xa9\xb1\x6e\xf2\x2c\x78\x04\x76\x63\x60\x97\x20\x97\x91\xbd\x54\x98\x12\xb7\x05\x5c\x8c\x73\x53\x52\x43\x63\x26\x5c\xde\x8a\xbb\xeb\x7b\x78\xb7\x01\xf2\x91\xc4\xf5\x24\xe2\xf9\x1c\x15\x33\xbb\xd8\x85\x04\x79\xa7\x68\x50\x33\x71\x30\x15\x6c\xe1\x3a\x3a\xff\x3b\xd0\xde\xf8\x48\x9b\x34\xca\x02\x69\x58\xb6\xb9\x8d\xe7\x86\xb2\x4d\x21\x8f\xb5\x2a\x7c\xe6\x07\xd3\x0e\x7c\x3a\x9f\xc8\x70\x97\x6d\xc1\x49\x92\x74\xff\xa1\x9c\x48\x21\xeb\xc5\xb1\x5c\x7c\x47\x42\x15\x55\x1f\xb6\xe1\x54\x9d\x67\xf8\xe2\x56\xcf\xa3\x50\xba\x17\x4b\xe3\x0b\xa6\xea\x8d\x75\xa1\x06\xd7\x49\x51\x38\x06\x9f\x6b\x46\xb1\xf1\x7d\x88\x40\xa3\x9a\x60\x24\xae\x31\x4c\x2d\x73\xfe\x50\x23\x80\x27\x78\x51\x20\xaf\x4b\xf0\xc3\x48\x02\x1b\xc6\x38\xfe\xc8\xed\xe5\xe0\x7f\x65\xc7\x1d\x53\x6f\x0c\xff\x38\x76\x7e\x5d\x00\x17\xed\x28\x7a\xef\xee\xe2\x07\x2a\x4a\xec\x16\xa0\xbb\x9a\x59\x13\x3b\x29\x00\x5b\xa1\xa1\xea\xad\xda\x7c\x3f\xd2\xc6\x47\x39\x56\x67\xa0\xfe\xb2\x3e\x9f\xe4\x81\x8b\x5b\xc1\xff\xaf\x3c\xaf\xf3\xfb\xf6\x43\x24\x88\x3d\x7f\xf4\x8f\xbd\x7e\x41\x3a\x1c\x27\x28\xd0\x49\xb6\x0a\x74\x4d\x8b\x07\x68\x35\x53\x6e\x48\xa3\x36\xec\x48\x79\x0d\xb4\x2c\x15\xd3\x7a\x2c\x90\xc5\xd9\xa9\x2c\x3e\x72\xd1\xb4\x06\xcc\xa9\xc1\xb3\x0d\xfb\xdb\x10\x1c\x2f\x75\x8b\x2f\x24\x98\xe4\x59\xcf\x60\xaa\xc3\x85\x0c\xcf\x49\x70\x60\xe6\x16\x6f\x47\x5a\xff\x6e\x70\xf9\x19\x29\xfc\xa0\x88\xfc\x57\x40\xc8\xdc\x2e\x3c\x79\xb5\xda\x0d\x1c\x99\xcd\x0a\xae\xe7\x7e\x1c\x60\x4a\x9e\x53\x3a\xcb\x86\x5b\x4c\x5b\x02\x4b\xeb\x6a\x6f\xb1\x7e\xb2\xc5\xe9\x64\x33\xb1\xcd\xfb\x94\x66\x5b\x1f\xae\x2b\xa8\x2c\xeb\xa6\x99\xff\xb1\x69\xbe\x32\xe0\xce\xd7\x52\xe2\x1b\xe0\x26\x7e\x2e\xf4\xd6\xa2\x1f\xa0\x17\xf7\xed\xc6\x4d\x4f\xcb\xca\x47\xc5\x8c\x3b\x28\x1d\x62\x9c\xbd\x28\x89\x13\x24\x60\x4e\x04\xb6\xaa\x38\x89\xa7\xca\xe0\xd8\x1e\x94\x89\xaa\x40\x3f\x8a\xad\x1a\xee\xf6\xd9\x78\x27\xf2\x31\xdc\xfd\xab\x70\x88\x4b\x18\x29\x99\x36\x5c\x9c\x16\xb6\xcd\x17\x85\x92\x86\xa2\x81\xbf\x1e\x08\xa5\xc3\x4b\x42\x6c\x74\xf8\x4a\x8a\x8a\x31\xcd\x16\x86\xa2\x63\xf1\xce\x17\x04\x6e\xfb\xde\xd4\x68\x71\x87\xa0\x47\x6e\x14\x97\xad\x7e\x20\xf7\xde\xa0\xf3\xff\xc8\x6e\xb7\x1b\x30\x5d\x51\x91\xbd\x94\x11\x25\x80\xac\xee\x26\x5e\xdd\x13\x54\x1f\xe5\x61\x20\x5a\x14\x85\x7d\x19\x59\xbf\x7a\x19\x09\x43\xb9\x60\x6a\xb1\xaf\x5b\x5e\xc6\x9e\xbe\xb0\x52\xf2\x2f\x12\xfb\x71\xfc\x29\x62\x65\xdd\x9c\xed\xdf\x8b\xeb\x71\x52\x33\x5d\xdf\xcb\xfd\x5d\x35\x72\xec\x37\x47\xb3\x7c\xb2\x9d\x84\x2e\x0f\xad\xfd\xf4\xb0\xf3\x85\xee\x0b\x2a\xcd\x6d\xf1\x67\xdb\x39\x94\xb2\x68\x8f\x58\x3e\x4b\x2c\xf9\x9f\x6b\x66\x1f\x7f\x3a\xdd\x96\xe9\x7b\xda\x34\xef\x67\xe8\xf7\x4f\x00\x00\x00\xff\xff\x66\xf5\xbf\x8d\x2a\x0b\x00\x00")

func www_app_jsx_bytes() ([]byte, error) {
	return bindata_read(
		_www_app_jsx,
		"www/app.jsx",
	)
}

func www_app_jsx() (*asset, error) {
	bytes, err := www_app_jsx_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "www/app.jsx", size: 2858, mode: os.FileMode(420), modTime: time.Unix(1430277631, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _www_filename_js = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func www_filename_js_bytes() ([]byte, error) {
	return bindata_read(
		_www_filename_js,
		"www/filename.js",
	)
}

func www_filename_js() (*asset, error) {
	bytes, err := www_filename_js_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "www/filename.js", size: 0, mode: os.FileMode(420), modTime: time.Unix(1430277349, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _www_index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x54\x4f\x6f\x1b\x2f\x14\x3c\x3b\x9f\x02\x73\xb2\xa5\x00\xb1\xf3\x8b\x7e\x6a\xcb\xfa\xd2\xe6\xd0\x5e\x5a\xa9\x39\xb4\x8a\x72\xc0\xf0\x36\x8b\xcb\xc2\x16\xf0\xae\xad\x24\xdf\xbd\xb0\xbb\x4d\xdc\x24\xaa\x7a\x48\xa4\x95\xde\x9b\x59\x98\x19\xfe\x08\x3e\xfd\xf0\xf9\xfd\xc5\xf7\x2f\xe7\xa8\x8a\xb5\x59\x1d\xf1\x5c\x90\x11\xf6\xba\xc0\x60\xf1\xea\x68\xc2\x2b\x10\x2a\xd5\x09\xaf\x21\x0a\x64\x45\x0d\x05\x6e\x35\x74\x8d\xf3\x11\x23\xe9\x6c\x04\x1b\x0b\xdc\x69\x15\xab\x42\x41\xab\x25\x90\x1e\x1c\x23\x6d\x75\xd4\xc2\x90\x20\x85\x81\x62\x81\x7b\x99\x29\x21\x97\xba\x44\x26\xa2\x8f\xe7\xe8\xcd\x55\xe6\x26\x3c\x48\xaf\x9b\xd8\xf7\x93\x59\xb9\xb5\x32\x6a\x67\x67\xf3\x9b\x9e\x98\xb4\xc2\x23\x28\x51\x81\x0e\xfe\xdc\xbd\x1b\xfe\x75\xda\x2a\xd7\xd1\x94\x23\x38\x03\x69\xcc\x23\xe2\xf6\x16\xdd\x18\x77\xfd\x16\xca\xe3\x4e\x78\x9b\x2b\x78\xef\x7c\x6e\x94\xce\x65\x54\xba\x9b\xcd\xe7\x7d\xc7\xd9\x41\x9a\x31\x19\x0a\x5e\x16\x98\x31\xa9\xec\x26\x50\x69\xdc\x56\x95\x46\x78\x48\x36\x35\x13\x1b\xb1\x63\x46\xaf\x03\xcb\xbb\x77\x16\x2a\xdd\xb2\x53\xfa\x3f\x5d\x3e\x60\x5a\x6b\x4b\x37\x01\xaf\x5e\x43\x9c\x34\x5e\xdb\xf8\x52\x36\x9b\x9f\x5b\xf0\x7b\xb6\xa0\x8b\xf4\x8d\xe8\x25\x74\x21\x9c\x91\x94\xb1\x4e\xe9\xff\xa3\x27\xf7\xf0\x55\x64\xc5\x53\x59\x3e\xbd\x04\xab\x74\x79\x45\x48\x0f\x8d\xb6\x3f\x90\x07\x53\xe0\x10\xf7\x06\x42\x05\x90\xae\x73\xe5\xa1\xcc\xa6\xb5\xd8\x25\x5f\xba\x76\x2e\x86\xe8\x45\x93\x41\xf6\xbd\x27\x92\x5d\x32\x64\x32\x84\x07\xae\xdf\xa5\xc4\x0c\xf7\xfc\xcf\x55\x94\x6b\x5a\x03\xf3\x20\x64\x24\x27\x74\x71\x4a\x97\x4f\x13\x3e\x37\xe3\xd3\xd7\x6f\x17\x5e\xd8\x50\x3a\x5f\x83\xff\xc7\xa9\xd2\x29\xa0\xe3\xc9\xe5\xd4\x43\x4b\x96\x74\x91\x42\x3f\x77\x94\x7f\xdf\x8d\x9e\xfa\xbd\x30\xce\xc6\x07\x81\xaf\x9d\xda\xf7\x93\x95\x6e\x91\x56\x05\x16\x4d\x93\x65\x13\x3c\xcc\x14\xf7\x4d\x7a\x31\x22\xec\x22\xdb\x84\x1d\x1e\x33\xa6\xb1\x8f\x53\x70\x36\x28\x72\x36\xbc\x45\xbf\x02\x00\x00\xff\xff\xaf\x4d\xa8\xde\x9c\x04\x00\x00")

func www_index_html_bytes() ([]byte, error) {
	return bindata_read(
		_www_index_html,
		"www/index.html",
	)
}

func www_index_html() (*asset, error) {
	bytes, err := www_index_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "www/index.html", size: 1180, mode: os.FileMode(420), modTime: time.Unix(1430277349, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _www_style_css = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x90\xcf\x0a\xc2\x30\x0c\x87\xef\x7d\x8a\xc0\xce\x95\xb6\x30\x84\xed\x69\xda\xa6\x6e\x65\x6e\x91\xcc\x31\xff\xe0\xbb\xdb\xe0\x45\x14\xc1\x91\x43\x0e\xdf\x97\x84\xfc\x02\xe1\x15\xee\x0a\x20\xf8\x38\x74\x4c\xcb\x84\x3a\xd2\x91\xb8\x81\xea\x60\xa5\x5a\x78\x28\xd5\xbb\x97\x44\x8c\x89\x35\x7b\xcc\xcb\xdc\x80\xd9\xb9\x3a\x8d\x6d\x01\x7a\xa4\x9b\xfe\x4d\xd7\x14\x86\x7c\xfe\x47\xb8\xe8\xb9\xf7\x48\xab\x50\x53\xe8\x47\xab\xac\xd9\x47\x6b\xe0\xed\xe6\x86\x89\x6d\xf2\x77\x20\x35\x46\x67\x93\xc0\x93\x47\xcc\x53\x27\x6b\xac\xfc\x50\x32\x7a\x06\x00\x00\xff\xff\x6e\x34\x71\xfe\x4a\x01\x00\x00")

func www_style_css_bytes() ([]byte, error) {
	return bindata_read(
		_www_style_css,
		"www/style.css",
	)
}

func www_style_css() (*asset, error) {
	bytes, err := www_style_css_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "www/style.css", size: 330, mode: os.FileMode(420), modTime: time.Unix(1430280014, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"www/.app.jsx.swp": www_app_jsx_swp,
	"www/app.js": www_app_js,
	"www/app.jsx": www_app_jsx,
	"www/filename.js": www_filename_js,
	"www/index.html": www_index_html,
	"www/style.css": www_style_css,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"www": &_bintree_t{nil, map[string]*_bintree_t{
		".app.jsx.swp": &_bintree_t{www_app_jsx_swp, map[string]*_bintree_t{
		}},
		"app.js": &_bintree_t{www_app_js, map[string]*_bintree_t{
		}},
		"app.jsx": &_bintree_t{www_app_jsx, map[string]*_bintree_t{
		}},
		"filename.js": &_bintree_t{www_filename_js, map[string]*_bintree_t{
		}},
		"index.html": &_bintree_t{www_index_html, map[string]*_bintree_t{
		}},
		"style.css": &_bintree_t{www_style_css, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

