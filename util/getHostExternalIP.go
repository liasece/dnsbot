package util

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/liasece/micserver/log"
)

// call http request and get the external IP
func GetHostExternalIP(server string) (string, error) {
	if !strings.HasPrefix(server, "http") {
		server = "http://" + server
	}
	resp, err := http.Get(server)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	ipReg := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	// find all ip
	ips := ipReg.FindAllString(string(body), -1)
	if len(ips) == 0 {
		return "", errors.New("no ip found")
	}
	log.Info("get external ip finish", log.String("body", string(body)), log.String("ip", ips[0]))
	return ips[0], nil
}
