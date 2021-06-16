package prommetrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type PromMetrics interface {
	IncCreate()
	IncUpdate()
	IncRemove()
}

type promMetrics struct {
	createCounter prometheus.Counter
	updateCounter prometheus.Counter
	removeCounter prometheus.Counter
}

func New(serviceName string) *promMetrics {
	return &promMetrics{
		createCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_ops_total", serviceName),
			Help: "The total number of CREATE events",
		}),
		updateCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_ops_total", serviceName),
			Help: "The total number of UPDATE events",
		}),
		removeCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_remove_ops_total", serviceName),
			Help: "The total number of REMOVE events",
		}),
	}
}

func (m *promMetrics) IncCreate() {
	m.createCounter.Inc()
}

func (m *promMetrics) IncUpdate() {
	m.updateCounter.Inc()
}

func (m *promMetrics) IncRemove() {
	m.removeCounter.Inc()
}
