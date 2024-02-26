package atq

import (
	"net/url"
	"testing"
)

func TestTaskStateString(t *testing.T) {
	tst := map[TaskState]string{
		TaskStateActive:      "active",
		TaskStatePending:     "pending",
		TaskStateScheduled:   "scheduled",
		TaskStateRetry:       "retry",
		TaskStateCompleted:   "completed",
		TaskStateArchived:    "archived",
		TaskStateAggregating: "aggregating",
	}
	for k, v := range tst {
		if taskStateString[k] != v {
			t.Errorf("expected %+v, returned %+v", v, taskStateString[k])
		}
	}
}

func TestParseRedisUri(t *testing.T) {
	addr := "rediss://u0:p0@server:44567/1"
	u, _ := url.Parse(addr)
	r, e := parseRedisURI(u)
	if e != nil {
		t.Errorf("error parsing Redis URI")
	}
	c, ok := r.(RedisClientOpt)
	if !ok {
		t.Errorf("wrong 'parseRedisURI' return format")
	}
	if c.Addr != "server:44567" {
		t.Errorf("wrong host address %+v:", c.Addr)
	}
	if c.DB != 1 {
		t.Errorf("wrong Redis DB number %+v:", c.DB)
	}
}

func TestParseRedisSocketURI(t *testing.T) {
	addr := "redis-socket://u0:p0@localhost:44567/some_path?db=1"
	u, _ := url.Parse(addr)
	r, e := parseRedisSocketUri(u)
	if e != nil {
		t.Errorf("error parsing Redis socket URI")
	}
	c, ok := r.(RedisClientOpt)
	if !ok {
		t.Errorf("wrong 'parseRedisSocketURI' return format")
	}
	if c.Addr != "/some_path" {
		t.Errorf("wrong host address %+v:", c.Addr)
	}
	if c.DB != 1 {
		t.Errorf("wrong Redis DB number %+v:", c.DB)
	}
}
