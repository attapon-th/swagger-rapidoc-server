package model

import (
	"errors"
	"os"
	"strings"

	"github.com/attapon-th/go-pkgs/zlog/log"
	"github.com/goccy/go-json"
)

type Credential struct {
	URL       string `json:"url"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	API       string `json:"api"`
	Path      string `json:"path"`
}

func LoadCredential(filename string) (c Credential, err error) {
	if !strings.HasSuffix(filename, ".json") {
		err = errors.New("File extension not supported")
		return
	}
	var f *os.File
	if f, err = os.OpenFile(filename, os.O_RDONLY, os.ModePerm); err != nil {
		return
	}
	err = json.NewDecoder(f).Decode(&c)
	log.Debug().Interface("credential", c).Str("endpoint", c.GetEndpoint()).Bool("ssl", c.UseSSL()).Send()
	return
}

func (c Credential) GetEndpoint() string {
	ss := strings.Split(c.URL, "://")
	// log.Debug().Strs("url", ss).Send()
	if len(ss) == 2 {
		return ss[1]
	}
	return ""
}

func (c Credential) UseSSL() bool {
	return strings.HasPrefix(c.URL, "https://")
}
