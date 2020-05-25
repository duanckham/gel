package gel

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/duanckham/gel/pb"
	"github.com/duanckham/hands"

	"github.com/duanckham/gel/utils"
)

// LogVariablePlaceholder ...
const LogVariablePlaceholder = "??"

type record struct {
	Ts       time.Time
	Numbers  map[string]int64
	Instants map[string]float64
	Logs     map[string]*pb.Logs
}

type records [3]*record
type recordsTriggerFunc func(rec *pb.Record)

// Gel ...
type Gel interface {
	Increment(name string, value int64)
	Decrement(name string, value int64)
	Gauge(name string, value float64)
	Log(message string, parameters ...interface{})
	SetTrigger(recordsTriggerFunc)
}

// Gel implement.
type gi struct {
	round      int32
	rec        [3]*record
	recordSP   *sync.Pool
	numbersMu  sync.Mutex
	instantsMu sync.Mutex
	logsMu     sync.Mutex
	callback   recordsTriggerFunc
}

// New ...
func New(period time.Duration) Gel {
	g := &gi{}

	g.recordSP = &sync.Pool{
		New: func() interface{} {
			return record{
				Ts:       time.Now(),
				Numbers:  map[string]int64{},
				Instants: map[string]float64{},
				Logs:     map[string]*pb.Logs{},
			}
		},
	}

	g.rec = records{}

	for i := 0; i < len(g.rec); i++ {
		t := g.recordSP.Get().(record)
		g.rec[i] = &t
	}

	return g.interval(period)
}

// Increment ..
func (g *gi) Increment(name string, value int64) {
	defer g.numbersMu.Unlock()
	g.numbersMu.Lock()

	if v, ok := g.rec[g.round].Numbers[name]; ok {
		g.rec[g.round].Numbers[name] = v + value
	} else {
		g.rec[g.round].Numbers[name] = value
	}
}

// Decrement ..
func (g *gi) Decrement(name string, value int64) {
	g.Increment(name, -value)
}

// Gauge ...
func (g *gi) Gauge(name string, value float64) {
	defer g.instantsMu.Unlock()
	g.instantsMu.Lock()
	g.rec[g.round].Instants[name] = value
}

// Log ...
func (g *gi) Log(template string, parameters ...interface{}) {
	defer g.logsMu.Unlock()
	g.logsMu.Lock()

	tsb := g.rec[g.round].Ts
	now := time.Now()

	m := pb.Message{
		Parameters: utils.InterfacesToStrings(parameters),
		Offset:     int64(now.Sub(tsb)),
	}

	if v, ok := g.rec[g.round].Logs[template]; ok {
		v.Logs = append(v.Logs, &m)

		g.rec[g.round].Logs[template] = v
	} else {
		g.rec[g.round].Logs[template] = &pb.Logs{
			Logs: []*pb.Message{&m},
		}
	}
}

// SetTrigger ...
func (g *gi) SetTrigger(f recordsTriggerFunc) {
	g.callback = f
}

// Dump ...
func (g *gi) dump() (*pb.Record, error) {
	// Landing round.
	l := (g.round + 1) % 3

	ts, err := ptypes.TimestampProto(g.rec[l].Ts)
	if err != nil {
		// TODO
	}

	p := pb.Record{
		Ts:       ts,
		Numbers:  g.rec[l].Numbers,
		Instants: g.rec[l].Instants,
		Logs:     g.rec[l].Logs,
	}

	t := g.recordSP.Get().(record)
	g.rec[l] = &t

	return &p, nil
}

func (g *gi) reap() {
	r, err := g.dump()
	if err != nil {
		// TODO
	}

	if len(r.Numbers) == 0 && len(r.Instants) == 0 && len(r.Logs) == 0 {
		return
	}

	g.callback(r)
}

func (g *gi) interval(period time.Duration) Gel {
	go func() {
		t := time.NewTicker(period)

		for {
			<-t.C

			if g.round == 2 {
				g.round = 0
			} else {
				g.round++
			}

			g.reap()
		}
	}()

	return g
}

// RecordUnit is the smallest unit of a record
type RecordUnit struct {
	T string
	K string
	V interface{}
	D time.Time
}

// Read ...
func Read(in *pb.Record) (chan RecordUnit, chan struct{}) {
	controller := hands.New()

	ch := make(chan RecordUnit)
	done := make(chan struct{})

	date, err := ptypes.Timestamp(in.Ts)
	if err != nil {
		// TODO
	}

	controller.Do(func(ctx context.Context) error {
		for k, v := range in.Numbers {
			ch <- RecordUnit{
				T: "number",
				K: k,
				V: v,
				D: date,
			}
		}

		return nil
	})

	controller.Do(func(ctx context.Context) error {
		for k, v := range in.Instants {
			ch <- RecordUnit{
				T: "instant",
				K: k,
				V: v,
				D: date,
			}
		}

		return nil
	})

	controller.Do(func(ctx context.Context) error {
		for template, logs := range in.Logs {
			segment := strings.Split(template, LogVariablePlaceholder)

			for _, message := range logs.Logs {
				t := make([]string, len(segment)+len(message.Parameters))

				for j, parameter := range message.Parameters {
					t = append(t, segment[j])
					t = append(t, parameter)

					if j == len(message.Parameters)-1 {
						t = append(t, segment[j+1:]...)
					}
				}

				ch <- RecordUnit{
					T: "log",
					K: "",
					V: strings.Join(t, ""),
					D: date.Add(time.Duration(message.Offset)),
				}
			}
		}

		return nil
	})

	controller.Done(func() {
		close(done)
		close(ch)
	})

	go controller.RunAll()

	return ch, done
}
