package main

import (
	"flag"
	"path/filepath"
	"time"

	"github.com/liasece/dnsbot/namesilo"
	"github.com/liasece/dnsbot/util"
	"github.com/liasece/micserver/log"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	fServer := flag.String("s", "namesilo", "The dns service name.")
	ipGetterServer := flag.String("p", "http://ifconfig.me", "The get host external ip service host.")
	fKey := flag.String("k", "***your key***", " To be replaced by your unique API key. Visit the API Manager page within your account for details.")
	fDomainURLs := &arrayFlags{}
	flag.Var(fDomainURLs, "d", "The domain associated list with the DNS resource record to modify")
	fInterval := flag.Duration("i", 300*time.Second, "The time of updating interval, like: 30s or 5m")
	flag.Parse()

	log.SetDefaultLogger(log.NewLogger(nil,
		log.Options().
			FilePaths(filepath.Join("./logs/", "log.log")).
			RecordTimeLayout("2006/01/02-15:04:05").
			AsyncWrite(false)),
	)
	log.Info("begin main", log.String("fServer", *fServer), log.String("fKey", *fKey), log.Reflect("fDomainURLs", *fDomainURLs), log.Duration("fInterval", *fInterval))

	domains, err := util.DecodeDomains(*fDomainURLs)
	if err != nil {
		log.Error("util.DecodeDomains error", log.ErrorField(err))
	}

	switch *fServer {
	case "namesilo":
		namesilo.Run(*fInterval, 0, *fKey, domains, *ipGetterServer)
	default:
		log.Error("unknown server name", log.String("server", *fServer))
	}
}
