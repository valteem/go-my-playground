// TODO: add йгукhandling of URL query variables

package env

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

const (
	filename = "envdata.json"
)

var (
	envs  = make([]EnvData, 0)
	pairs = make([]EnvPair, 0)
)

type ConnData struct {
	Proto string   `json:"proto,omitempty"`
	Host  string   `json:"host,omitempty"`
	Port  string   `json:"port,omitempty"`
	Path  []string `json:"path,omitempty"`
}

type EnvData struct {
	VarName string    `json:"var"`
	Conn    *ConnData `json:"conn"`
}

type EnvPair struct {
	Name  string
	Value string
}

func init() {

	buf, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read env data: %v", err)
	}

	err = json.Unmarshal(buf, &envs)
	if err != nil {
		log.Fatalf("failed to decode connection data: %v", err)
	}

	if len(envs) != 0 {

		for _, env := range envs {
			p := EnvPair{env.VarName, urlFromConnData(env.Conn)}
			pairs = append(pairs, p)
		}

		for _, p := range pairs {
			os.Setenv(p.Name, p.Value)
		}

	}

}

func urlFromConnData(c *ConnData) string {

	sb := strings.Builder{}
	sb.WriteString(c.Proto)
	sb.WriteString("://")
	sb.WriteString(c.Host)
	if c.Port != "" {
		sb.WriteString(":")
		sb.WriteString(c.Port)
	}
	if len(c.Path) != 0 {
		for _, p := range c.Path {
			sb.WriteString("/")
			sb.WriteString(p)
		}
	}

	return sb.String()

}
