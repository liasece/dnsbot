package namesilo

import (
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/liasece/dnsbot/util"
)

func TestRun(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// mock to list out the articles
	httpmock.RegisterResponder("GET", "https://www.namesilo.com/api/dnsListRecords",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `<?xml version="1.0"?>
			<namesilo>
				<request>
					<operation>dnsListRecords</operation>
					<ip>127.0.0.1</ip>
				</request>
				<reply>
					<code>300</code>
					<detail>success</detail>
					<resource_record>
						<record_id>a</record_id>
						<type>A</type>
						<host>liasece.com</host>
						<value>192.168.1.1</value>
						<ttl>7207</ttl>
						<distance>0</distance>
					</resource_record>
					<resource_record>
						<record_id>b</record_id>
						<type>A</type>
						<host>www.liasece.com</host>
						<value>192.168.1.1</value>
						<ttl>7207</ttl>
						<distance>0</distance>
					</resource_record>
					<resource_record>
						<record_id>c</record_id>
						<type>A</type>
						<host>unknown.liasece.com</host>
						<value>192.168.1.1</value>
						<ttl>7207</ttl>
						<distance>0</distance>
					</resource_record>
				</reply>
			</namesilo>
			`), nil
		},
	)

	// mock to list out the articles
	httpmock.RegisterResponder("GET", "https://www.namesilo.com/api/dnsUpdateRecord",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `<?xml version="1.0"?>
			<namesilo>
				<request>
					<operation>dnsUpdateRecord</operation>
					<ip>127.0.0.1</ip>
				</request>
				<reply>
					<code>300</code>
					<detail>success</detail>
					<record_id>d</record_id>
				</reply>
			</namesilo>
			`), nil
		},
	)

	type args struct {
		interval time.Duration
		times    int
		key      string
		domains  map[string]*util.ListenDomain
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "OK",
			args: args{
				interval: time.Second,
				times:    1,
				key:      "****key****",
				domains: func() map[string]*util.ListenDomain {
					res, err := util.DecodeDomains([]string{"www.liasece.com", "liasece.com", "http://nothis.liasece.com"})
					if err != nil {
						t.Error(err)
					}
					return res
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run(tt.args.interval, tt.args.times, tt.args.key, tt.args.domains, "")
		})
	}
}
