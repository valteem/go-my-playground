package atq

import (
	"crypto/tls"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

func TestParseSimpleRedisURI(t *testing.T) {
	uri := "rediss://u0:p0@server:44567/1"
	u, _ := url.Parse(uri)
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
	uri := "redis-socket://u0:p0@localhost:44567/some_path?db=1"
	u, _ := url.Parse(uri)
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
	uri := "redis-sentinel://:pwd@host1:44567,host2:44568,host3:44569?master=mName"
	u, _ := url.Parse(uri)
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

func TestParseRedisURI(t *testing.T) {

	tests := []struct {
		uri      string
		expected RedisConnOpt
	}{
		{
			"redis://localhost:44567",
			RedisClientOpt{Addr: "localhost:44567"},
		},
		{
			"rediss://localhost:44567",
			RedisClientOpt{Addr: "localhost:44567", TLSConfig: &tls.Config{ServerName: "localhost"}},
		},
		{
			"redis://localhost:44567/1",
			RedisClientOpt{Addr: "localhost:44567", DB: 1},
		},
		{
			"redis://:somepassword@localhost:44567",
			RedisClientOpt{Addr: "localhost:44567", Password: "somepassword"},
		},
		{
			"redis://:somepassword@192.168.1.100/2",
			RedisClientOpt{Addr: "192.168.1.100", Password: "somepassword", DB: 2},
		},
		{
			"redis-socket:///var/run/redis/redis.sock",
			RedisClientOpt{Network: "unix", Addr: "/var/run/redis/redis.sock"},
		},
		{
			"redis-socket://:somepassword@/var/run/redis/redis.sock",
			RedisClientOpt{Network: "unix", Addr: "/var/run/redis/redis.sock", Password: "somepassword"},
		},
		{
			"redis-socket:///var/run/redis/redis.sock?db=3",
			RedisClientOpt{Network: "unix", Addr: "/var/run/redis/redis.sock", DB: 3},
		},
		{
			"redis-socket://:somepassword@/var/run/redis/redis.sock?db=4",
			RedisClientOpt{Network: "unix", Addr: "/var/run/redis/redis.sock", Password: "somepassword", DB: 4},
		},
		{
			"redis-sentinel://localhost:44567,localhost:44568,localhost:44569?master=mname",
			RedisFailoverClientOpt{
				MasterName:    "mname",
				SentinelAddrs: []string{"localhost:44567", "localhost:44568", "localhost:44569"},
			},
		},
		{
			"redis-sentinel://:somepassword@localhost:44567,localhost:44568,localhost:44569?master=mname",
			RedisFailoverClientOpt{
				MasterName:       "mname",
				SentinelAddrs:    []string{"localhost:44567", "localhost:44568", "localhost:44569"},
				SentinelPassword: "somepassword",
			},
		},
	}

	for _, test := range tests {
		r, e := ParseRedisURI(test.uri)
		if e != nil {
			t.Errorf("ParseRedisURI(%q) returned an error: %+v", test.uri, e)
			continue // jump to next test in 'tests'
		}
		diff := cmp.Diff(test.expected, r, cmpopts.IgnoreUnexported(tls.Config{}))
		if diff != "" {
			t.Errorf("ParseRedisURI(%q) = %+v, expected %+v, (-expected, +result) %+v", test.uri, r, test.expected, diff)
		}
	}

}

func TestParseRedisURIErrors(t *testing.T) {

	tests := []struct {
		description string
		uri         string
	}{
		{
			"unsupported scheme",
			"rdb://localhost:44567",
		},
		{
			"missing scheme",
			"localhost:44567",
		},
		{
			"multiple db numbers",
			"redis://localhost:44567/1,2,3",
		},
		{
			"missing path for socket connection",
			"redis-socket://?db=1",
		},
		{
			"non integer db number",
			"redis-socket:///path/to/redis?db=two",
		},
	}

	for _, test := range tests {
		_, e := ParseRedisURI(test.uri)
		if e == nil {
			t.Errorf("%s: ParseRedisURI(%q) succeded for malformed input, should return error", test.description, test.uri)
		}
	}

}
