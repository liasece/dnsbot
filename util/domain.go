package util

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/liasece/micserver/log"
)

// ListenDomain type
type ListenDomain struct {
	Domain string
	Hosts  []string
}

// DecodeDomains func
func DecodeDomains(hostURLs []string) (map[string]*ListenDomain, error) {
	domainMap := make(map[string]*ListenDomain)
	for _, urlStr := range hostURLs {
		var urlObj *url.URL
		var err error
		if strings.HasPrefix(urlStr, "http://") || strings.HasPrefix(urlStr, "https://") {
			urlObj, err = url.Parse(urlStr)
		} else {
			urlObj, err = url.Parse("http://" + urlStr)
		}
		if err != nil {
			log.Error("url.Parse error", log.String("url", urlStr), log.ErrorField(err))
			return nil, err
		}
		dots := strings.Split(urlObj.Host, ".")
		if len(dots) < 2 {
			log.Error("unknown host", log.String("host", urlObj.Hostname()))
			return nil, fmt.Errorf("unknown host:%s", urlObj.Hostname())
		}
		domain := strings.Join(dots[len(dots)-2:], ".")
		host := strings.Join(dots[:len(dots)-2], ".")
		log.Info("info", log.String("host", host), log.String("domain", domain), log.String("url", urlStr))
		if domainObj, ok := domainMap[domain]; ok {
			domainObj.Hosts = append(domainObj.Hosts, host)
		} else {
			domainMap[domain] = &ListenDomain{
				Domain: domain,
				Hosts:  []string{host},
			}
		}
	}
	return domainMap, nil
}
