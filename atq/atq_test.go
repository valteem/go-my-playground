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

func TestParseRedisURI(t *testing.T) {
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

func TestParseRedisSentinelURI(t *testing.T) {
	addr := "redis-sentinel://:pwd@host1:44567,host2:44568,host3:44569?master=mName"
	u, _ := url.Parse(addr)
	r, e := parseRedisSentinelURI(u)
	if e != nil {
		t.Errorf("error parsing Redis sentinel URI")
	}
	c, ok := r.(RedisFailoverClientOpt)
	if !ok {
		t.Errorf("wrong 'parseRedisSentinelURI' output format")
	}
	sentinelAddrsExpected := []string{"host1:44567", "host2:44568", "host3:44569"}
	for i, v := range c.SentinelAddrs {
		if sentinelAddrsExpected[i] != v {
			t.Errorf("wrong host address: expect %s, get %s", sentinelAddrsExpected[i], v)
		}
	}
	if c.MasterName != "mName" {
		t.Errorf("wrong master name: %s", c.MasterName)
	}
	if c.SentinelPassword != "pwd" {
		t.Errorf("wrong sentinel password: %s", c.SentinelPassword)
	}
}
