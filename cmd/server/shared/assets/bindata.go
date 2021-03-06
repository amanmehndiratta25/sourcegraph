// Code generated by go-bindata. DO NOT EDIT.
// sources:
// nginx.conf (1.396kB)
// redis.conf.tmpl (300B)
// nginx/sourcegraph_backend.conf (192B)
// nginx/sourcegraph_http.conf (212B)
// nginx/sourcegraph_main.conf (174B)
// nginx/sourcegraph_server.conf (176B)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _nginxConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\x41\x6b\xdc\x30\x10\x85\xef\xfa\x15\x0f\x36\xd7\xda\xb9\x35\x64\xaf\xa5\xe4\xd0\x43\x68\xf6\xd0\x9e\xcc\x44\x9a\xb5\xd5\xd8\x1a\x33\x92\x9d\x2c\x61\xff\x7b\x91\xed\x24\xa6\x4d\x5a\xb6\x39\xd4\x60\x90\xc7\x33\x6f\xde\x37\x03\xda\x60\xd7\xf8\x08\x2b\x61\xef\x6b\xdc\x53\x44\xcd\x81\x95\x12\x3b\xdc\x1e\x70\x23\x83\x5a\xae\x95\xfa\xa6\x30\x1b\x7c\x97\x01\x96\x02\xc8\xfd\x18\x62\x42\x6a\x78\xa9\x1c\x94\x92\x97\x80\x24\x20\xe7\xf2\xeb\xf3\x37\xb5\xd8\x7d\xb9\x81\x28\xae\x76\xbb\x6b\xec\x99\xd2\xa0\x1c\xb3\xd4\x57\x26\x87\x4e\x94\x41\x09\x4d\x4a\x7d\xbc\x2c\x4b\x27\x36\x16\x71\xd5\xd3\x4a\x57\x92\xeb\x7c\x28\x43\xed\xc3\x83\x31\xac\x2a\x5a\xb5\x52\x23\x26\xc7\xaa\x5b\xd3\x7b\x87\x72\x24\x2d\x75\x58\xb2\x8a\xde\xbb\xad\x31\x1b\x7c\x12\x04\x49\x50\xee\x64\xe4\x02\xbb\xd9\x6e\xe2\x90\x22\x64\x8f\x55\xa3\xaa\x23\x1f\x8a\xcc\x32\xf1\xd9\x86\x42\xcd\xb8\xe5\x74\xcf\x1c\xcc\x06\x23\x6b\xf4\x12\x22\x28\x38\x74\x74\x80\x0f\xb6\x1d\x1c\xc3\x77\xbd\xca\xc8\xdd\xa4\x99\xe4\xf7\x91\x14\xe6\x29\x75\xf2\x56\xbe\xda\x74\x6b\x0c\x8f\x93\xc4\xa3\x39\x1a\x93\xc7\x81\x47\x03\x00\x91\x75\x64\xad\x92\xdc\x71\xc8\xa6\x73\x6a\x8e\x9f\x04\x97\xf5\xde\x84\x9b\xd5\xde\x03\x98\x15\xde\x86\x7c\x6e\xbe\x38\x27\x6b\x39\xc6\x69\x85\x13\x4e\x8e\x0d\x7d\x4c\xca\xd4\xe1\x96\xec\x1d\x07\xb7\xc0\x9f\x0c\xba\x94\xff\xca\xba\x12\x5b\xa8\xdf\xcd\xfb\x67\xe6\xb5\x8f\x05\xfb\x65\x99\x68\xc5\x52\xdb\x48\x4c\x97\x17\xe7\x17\xe7\xf3\x04\x8e\x66\xb5\xee\x7f\xc5\x9f\xab\xff\x3b\xfd\xca\xc6\x0a\xbe\xf5\x31\x71\xc0\xc7\x67\xe4\x29\x28\x76\xbe\x38\xca\x15\x73\x7e\x7a\x95\x87\x43\xd5\x53\x8c\xd3\xe5\x70\x59\x96\xcb\x4c\xb7\xaf\xa4\x45\x4e\x55\xc3\xe4\x58\x71\x25\x31\xe1\x2c\x97\x54\x79\xc2\x7f\xc9\xfe\xf6\xe1\xb3\xe8\x3d\xa9\x63\x97\x4f\x38\x9b\x33\xc8\xb9\xea\xa1\xda\x3f\xfd\xca\xa7\x13\x84\xae\x55\x92\xe0\x2c\xda\x86\x3b\x7e\xa9\x3b\x2e\x7b\x3e\x9a\x9f\x01\x00\x00\xff\xff\xc6\x5e\x09\x5c\x74\x05\x00\x00")

func nginxConfBytes() ([]byte, error) {
	return bindataRead(
		_nginxConf,
		"nginx.conf",
	)
}

func nginxConf() (*asset, error) {
	bytes, err := nginxConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nginx.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3d, 0x23, 0x63, 0xc0, 0xfa, 0x45, 0x34, 0x79, 0x1a, 0x5, 0xc6, 0xad, 0x2d, 0xc0, 0x1, 0xfb, 0x7, 0xd4, 0x6e, 0xd, 0xec, 0xc3, 0xab, 0xef, 0xab, 0x2b, 0xc, 0x3d, 0xd7, 0x69, 0xfd, 0x9f}}
	return a, nil
}

var _redisConfTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcd\x41\x6e\x02\x31\x0c\x05\xd0\x7d\x4e\xf1\x25\xb6\x50\xc1\xa6\x27\xe8\x45\x42\xe2\x06\xab\x89\x3d\xb2\xcd\x4c\x47\x88\xbb\x57\x29\x95\xba\xf4\xd7\xf7\x7f\x07\xe4\xde\x75\x43\x2e\x85\xdc\xf1\x69\x3a\x66\x02\x16\x8f\x2c\x85\x3c\x2d\xa6\x41\x25\xa8\x9e\x86\x56\x82\x68\x4a\x07\x74\x1e\x1c\x18\x34\xd4\x76\xdc\x3d\x37\x3a\xc2\x28\xee\x26\x20\x33\x35\x6c\x37\x12\xdc\x38\x82\xa5\xbd\xda\x69\xe4\xef\xbf\x87\x4b\xbb\xfe\x5f\xa7\x45\x3b\x97\x1d\xa2\xb4\x72\x09\x56\x79\x01\x2b\xa1\xe8\x98\x4c\xd7\x86\x50\x54\xf6\xaf\x23\x72\xad\x3c\x4b\xb9\xf7\x1d\x2e\x79\xf1\x9b\x86\x83\x56\x9a\xc3\x67\x0c\x96\x7b\x90\xa7\xca\x86\xc7\x03\x6f\x1f\x6c\x78\x3e\x53\x5e\x16\x92\xaa\xd2\x77\xec\xe4\xc9\xf3\x4a\x78\x3f\x9f\x71\xf9\xc5\x28\x7b\x60\x25\xbb\xaa\xd3\xe4\x1a\x4b\x4b\x5d\x5b\xa7\x95\x3a\xb6\x6c\x32\x83\x9f\x00\x00\x00\xff\xff\x3d\xe3\xfc\x7d\x2c\x01\x00\x00")

func redisConfTmplBytes() ([]byte, error) {
	return bindataRead(
		_redisConfTmpl,
		"redis.conf.tmpl",
	)
}

func redisConfTmpl() (*asset, error) {
	bytes, err := redisConfTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "redis.conf.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x39, 0xdc, 0xa1, 0x4e, 0x7, 0x3e, 0xee, 0x45, 0xc6, 0x63, 0x5d, 0xf, 0x9c, 0xfc, 0xc1, 0x41, 0xd0, 0xc1, 0xc5, 0x78, 0xf2, 0x89, 0xe0, 0xe1, 0x8f, 0x38, 0xbc, 0x9b, 0x8d, 0xff, 0xcf, 0x0}}
	return a, nil
}

var _nginxSourcegraph_backendConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcd\x3f\x6e\x87\x30\x0c\x47\xf1\x3d\xa7\xf8\x4a\x59\x2b\x4e\xd1\xa5\x73\x7b\x81\xfc\x31\xc1\x22\xd8\xc8\x71\x04\xdc\xbe\x62\xfd\xed\x4f\x9f\x17\xf1\xb7\xf1\x40\x51\x59\xb9\x81\x07\x1a\x09\x59\x72\xaa\xc8\x0f\x7e\x75\x5a\xa1\x66\xe9\xdc\x16\x7c\x2b\x44\x1d\x87\x56\x5e\x9f\x2f\xb0\xe3\xe2\xde\x91\x09\x46\x97\xb1\x3b\x49\x88\x50\xc1\xf0\x64\x3e\xcf\xe5\x93\x66\x29\x7d\x56\xaa\x60\x81\x6f\x84\x79\x0e\x37\x4a\xc7\x9b\x38\xdd\x8e\x55\x0d\x39\x95\x9d\xe4\xbd\x87\x08\x69\x2c\xf7\xf2\x0a\x21\x44\xfc\x88\x93\x38\xab\xa4\xde\x1f\xe4\x9e\x64\x0f\xff\x01\x00\x00\xff\xff\xfa\xe3\xdf\xcd\xc0\x00\x00\x00")

func nginxSourcegraph_backendConfBytes() ([]byte, error) {
	return bindataRead(
		_nginxSourcegraph_backendConf,
		"nginx/sourcegraph_backend.conf",
	)
}

func nginxSourcegraph_backendConf() (*asset, error) {
	bytes, err := nginxSourcegraph_backendConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nginx/sourcegraph_backend.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcf, 0x33, 0xf6, 0x61, 0xa1, 0x6e, 0xb8, 0x59, 0xc2, 0xa0, 0x57, 0x56, 0xfa, 0x3d, 0xb3, 0x7e, 0xf7, 0xef, 0x1b, 0x71, 0x2e, 0x54, 0x57, 0x34, 0xdf, 0x8b, 0x6c, 0xbd, 0x68, 0x36, 0xde, 0x47}}
	return a, nil
}

var _nginxSourcegraph_httpConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8d\xb1\x6a\xc4\x30\x10\x05\x7b\x7d\xc5\x03\xb7\xc1\x24\x45\xaa\xb4\x69\x53\x25\x90\xd2\xc8\xd2\x3b\x79\x41\xb7\x6b\xa4\x35\x27\xe7\xeb\x83\xdb\xab\x67\x98\x99\xf0\xb3\x49\x47\x32\xbd\x49\x81\x74\x14\x2a\x5b\x74\x66\xac\x27\xbe\xed\x68\x89\xa5\xc5\x7d\x9b\xf1\x69\x50\x73\xdc\x2d\xcb\xed\x7c\x81\x38\x1e\x52\x2b\x56\xa2\xf1\xd1\xc4\x9d\x1a\x26\x98\xa2\x7b\x6c\x7e\xec\xf3\x73\x5a\x34\xd5\x23\x33\x43\x14\xbe\x11\x9b\xfb\x7e\x61\xe7\xf0\xeb\xa6\x45\x74\xcc\x97\x1f\xc2\x84\x5f\x22\x45\xc5\xb1\x57\x8b\x19\x35\xb6\x42\x70\x38\xb5\x8b\x69\x0f\xa9\x0a\xd5\x97\x7b\x1c\xcb\x6a\xf9\x5c\xba\xfc\x11\x6f\xef\xaf\x5f\x1f\xe1\x3f\x00\x00\xff\xff\xfc\xc0\xaa\xd0\xd4\x00\x00\x00")

func nginxSourcegraph_httpConfBytes() ([]byte, error) {
	return bindataRead(
		_nginxSourcegraph_httpConf,
		"nginx/sourcegraph_http.conf",
	)
}

func nginxSourcegraph_httpConf() (*asset, error) {
	bytes, err := nginxSourcegraph_httpConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nginx/sourcegraph_http.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf, 0xd5, 0xb, 0x20, 0x74, 0xdb, 0xe1, 0xb2, 0x81, 0x38, 0x2c, 0xcc, 0x13, 0x1f, 0x9c, 0x78, 0xa4, 0xe3, 0x3d, 0xce, 0x37, 0xe4, 0x7e, 0x1b, 0x9a, 0x7e, 0x75, 0x8f, 0x92, 0x34, 0x9, 0x82}}
	return a, nil
}

var _nginxSourcegraph_mainConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8d\x31\x0e\x83\x30\x10\x04\x7b\xbf\x62\x25\xda\x88\x57\xa4\x49\x9d\x7c\xc0\xe0\xc3\x9c\x62\xf6\xd0\x71\x08\xf8\x7d\x44\x9b\x7a\x46\x33\x1d\x3e\xb3\x6e\x18\x8d\x93\x56\xe8\x86\x2a\x14\xcf\x21\x05\xc3\x85\xb7\xed\x3e\x4a\xf5\xbc\xce\x3d\x9e\x06\x5a\x60\xb1\xa2\xd3\xf5\x80\x06\x0e\x6d\x0d\x83\xc0\xe5\x70\x8d\x10\xa6\x0e\x46\x6c\x91\x3d\xf6\xb5\xff\x4f\x2b\xc7\xb6\x17\x29\x50\x22\x66\xc1\x92\x95\x37\x0e\x39\xe3\xbe\xb1\x2a\xcf\xfe\xf6\x53\xea\xf0\x62\x08\x43\x8d\xb9\xb5\x0b\x43\xcb\xfc\xa6\x5f\x00\x00\x00\xff\xff\x7d\xc8\x2e\x34\xae\x00\x00\x00")

func nginxSourcegraph_mainConfBytes() ([]byte, error) {
	return bindataRead(
		_nginxSourcegraph_mainConf,
		"nginx/sourcegraph_main.conf",
	)
}

func nginxSourcegraph_mainConf() (*asset, error) {
	bytes, err := nginxSourcegraph_mainConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nginx/sourcegraph_main.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2b, 0xe0, 0x54, 0x8, 0x8b, 0xc7, 0xfa, 0x50, 0x47, 0x74, 0x8f, 0x86, 0xcf, 0xef, 0xde, 0xc5, 0x8b, 0x99, 0xb0, 0xc2, 0xf9, 0x95, 0x17, 0x94, 0xc0, 0x22, 0x48, 0x47, 0x6b, 0x1b, 0x8e, 0x3c}}
	return a, nil
}

var _nginxSourcegraph_serverConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8d\x3d\x8e\x83\x30\x10\x85\x7b\x9f\xe2\x49\xb4\x2b\x4e\xb1\xcd\xd6\x9b\x0b\x18\x3c\x98\x51\x9c\x37\x68\x3c\x04\xb8\x7d\x44\x9b\xfe\xfb\x19\xf0\x58\xb5\x63\x36\x2e\x5a\xa1\x1d\x55\x28\x9e\x43\x0a\xa6\x0b\xff\xb6\xfb\x2c\xd5\xf3\xb6\x8e\xf8\x35\xd0\x02\x2f\x2b\xba\x5c\x3f\xd0\xc0\xa1\xad\x61\x12\xb8\x1c\xae\x11\xc2\x34\xc0\x88\x1e\xd9\x63\xdf\xc6\xef\xb4\x72\x6e\x7b\x91\x02\x25\x62\x15\x74\xf1\xb7\xf8\x0d\x84\x9c\x71\xff\x58\x95\xe7\x78\x1b\x29\x0d\xf8\x63\x08\x43\x8d\xb9\xb5\x0b\x53\xcb\x7c\xa6\x4f\x00\x00\x00\xff\xff\x9c\x8c\x8b\x56\xb0\x00\x00\x00")

func nginxSourcegraph_serverConfBytes() ([]byte, error) {
	return bindataRead(
		_nginxSourcegraph_serverConf,
		"nginx/sourcegraph_server.conf",
	)
}

func nginxSourcegraph_serverConf() (*asset, error) {
	bytes, err := nginxSourcegraph_serverConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nginx/sourcegraph_server.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5, 0xe0, 0x16, 0xbe, 0xd, 0x61, 0x2d, 0xed, 0xc, 0x94, 0x64, 0xfa, 0x64, 0x3c, 0x78, 0x30, 0xb, 0x84, 0xdf, 0x75, 0x70, 0xfe, 0x5a, 0x9b, 0xd8, 0x70, 0xad, 0x81, 0xd9, 0x30, 0x73, 0x67}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"nginx.conf": nginxConf,

	"redis.conf.tmpl": redisConfTmpl,

	"nginx/sourcegraph_backend.conf": nginxSourcegraph_backendConf,

	"nginx/sourcegraph_http.conf": nginxSourcegraph_httpConf,

	"nginx/sourcegraph_main.conf": nginxSourcegraph_mainConf,

	"nginx/sourcegraph_server.conf": nginxSourcegraph_serverConf,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"nginx": {nil, map[string]*bintree{
		"sourcegraph_backend.conf": {nginxSourcegraph_backendConf, map[string]*bintree{}},
		"sourcegraph_http.conf":    {nginxSourcegraph_httpConf, map[string]*bintree{}},
		"sourcegraph_main.conf":    {nginxSourcegraph_mainConf, map[string]*bintree{}},
		"sourcegraph_server.conf":  {nginxSourcegraph_serverConf, map[string]*bintree{}},
	}},
	"nginx.conf":      {nginxConf, map[string]*bintree{}},
	"redis.conf.tmpl": {redisConfTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
