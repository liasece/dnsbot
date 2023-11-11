package util

import (
	"errors"
	"io"
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
	req, err := http.NewRequest(http.MethodGet, server, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", "curl/7.64.1")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	ipReg := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	// find all ip
	ips := ipReg.FindAllString(string(body), -1)
	if len(ips) == 0 {
		return "", errors.New("no ip found: (server: " + server + ")" + string(body))
	}
	log.Info("get external ip finish", log.String("body", string(body)), log.String("ip", ips[0]))
	return ips[0], nil
}
