package influxDB

import (
	"fmt"
	"strings"
)

type Query struct {
	bucket   string
	qRange   *QRange
	qFilter  []string
	influxQL string
}

type QRange struct {
	Start string
	Stop  string
}

func NewQuery(bucket string) *Query {
	return &Query{
		bucket: bucket,
	}
}

func (q *Query) AddFilter(filter ...string) *Query {
	q.qFilter = append(q.qFilter, filter...)
	return q
}

func (q *Query) AddRange(start, stop string) *Query {
	q.qRange = &QRange{
		Start: start,
		Stop:  stop,
	}
	return q
}

func (q *Query) gen() {
	q.influxQL = fmt.Sprintf(`from(bucket:"%s")`, q.bucket)
	if q.qRange != nil {
		var rangQL []string
		if len(q.qRange.Start) != 0 {
			rangQL = append(rangQL, fmt.Sprintf("start: %s", q.qRange.Start))
		}
		if len(q.qRange.Stop) != 0 {
			rangQL = append(rangQL, fmt.Sprintf("stop: %s", q.qRange.Stop))
		}
		q.influxQL = fmt.Sprintf(`%s|> range(%s)`, q.influxQL, strings.Join(rangQL, ","))
	}
	for i := range q.qFilter {
		q.influxQL = fmt.Sprintf("%s|> filter(%s)", q.influxQL, q.qFilter[i])
	}
	return
}

func (q *Query) QL() string {
	q.gen()
	return q.influxQL
}
