package metric

import (
	"log"
	"strconv"
	"strings"

	"github.com/toolkits/sys"
)

func DuMetrics() (L []*MetricValue) {
	for _, path := range []string{
		"/tmp",
	} {
		out, err := sys.CmdOutNoLn("du", "-bs", path)
		if err != nil {
			log.Println("du -bs", path, "fail", err)
			continue
		}

		arr := strings.Fields(out)
		if len(arr) == 1 {
			continue
		}

		size, err := strconv.ParseUint(arr[0], 10, 64)
		if err != nil {
			log.Println("cannot parse du -bs", path, "output")
			continue
		}

		L = append(L, GaugeValue("du.bs", size, "path="+path))
	}

	return
}
