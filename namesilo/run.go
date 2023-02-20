package namesilo

import (
	"time"

	"github.com/liasece/dnsbot/util"
	"github.com/liasece/micserver/log"
)

// Run namesilo function. if times == 0, infinite loop
func Run(interval time.Duration, times int, key string, domains map[string]*util.ListenDomain, ipGetterServer string) {
	if len(domains) < 1 {
		log.Error("no domain in list", log.Reflect("domains", domains))
		return
	}
	updateDNSLoop(interval, times, key, domains, ipGetterServer)
}

func updateDNSLoop(interval time.Duration, times int, key string, domains map[string]*util.ListenDomain, ipGetterServer string) {
	tick := time.NewTicker(interval)
	defer tick.Stop()
	runTimes := 0
	do := func() {
		for _, domain := range domains {
			err := checkDomain(key, domain, ipGetterServer)
			if err != nil {
				log.Error("Update DNS record failed", log.Reflect("domain", domain), log.ErrorField(err))
			}
		}
	}

	{
		do()
		runTimes++
	}

	for {
		if times > 0 && runTimes >= times {
			log.Info("run finish", log.Int("runTimes", runTimes), log.Int("times", times))
			break
		}
		select {
		case <-tick.C:
			do()
		}
		runTimes++
	}
}

func checkDomain(key string, domain *util.ListenDomain, ipGetterServer string) error {
	externalIP, err := util.GetHostExternalIP(ipGetterServer)
	if err != nil {
		return err
	}
	listResp, err := dnsList(domain.Domain, key)
	if err != nil {
		return err
	}

	for _, host := range domain.Hosts {
		fullHost := host + "." + domain.Domain
		if host == "" {
			fullHost = domain.Domain
		}
		find := false
		// find the one need to be updated
		for _, item := range listResp.ListReply.DNSRecords {
			if item.Host == fullHost {
				if item.Value != externalIP {
					// update record
					log.Info("Update dns record IP", log.String("host", item.Host), log.String("oldIP", item.Value), log.String("newIP", externalIP))
					err = dnsUpdate(key, domain.Domain, item.RecordID, host, externalIP)
					if err != nil {
						return err
					}
				} else {
					log.Info("Nothing to do", log.String("host", item.Host), log.String("recordedIP", item.Value), log.String("localIP", externalIP))
				}
				find = true
			}
		}
		if !find {
			log.Error("No target host record to update", log.String("fullHost", fullHost))
		}
	}
	return nil
}
