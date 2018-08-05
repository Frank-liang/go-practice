package metric

import (
	"log"

	"github.com/glorydeath/nux"
)

func UdpMetrics() []*MetricValue {
	udp, err := nux.Snmp("Udp")
	if err != nil {
		log.Println("read snmp fail", err)
		return []*MetricValue{}
	}

	count := len(udp)
	ret := make([]*MetricValue, count)
	i := 0
	for key, val := range udp {
		ret[i] = CounterValue("snmp.Udp."+key, val)
		i++
	}

	return ret
}
