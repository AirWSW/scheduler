package config

import (
	"testing"

	"github.com/BurntSushi/toml"
)

func TestLoadConfigFile(t *testing.T) {
	configStr :=`
[node]
id    = "node_id1"
name  = "master1"
role  = "master-slave"
debug = true

[node.master]
hostname = "master"
address  = "0.0.0.0"
port     = 8080

[node.server]
hostname = "master"
address  = "0.0.0.0"
port     = 8080

[node.token]
access_token   = "access_token"
refresh_token  = "refresh_token"
register_token = "register_token"

[[schedule]]
id      = "plan_id1"
name    = "plan1"
node_id = "node_id1"
status  = "enabled"
	[schedule.task]
	id     = "task_id1"
	name   = "task1"
	status = "enabled"

[[schedule]]
id      = "plan_id2"
name    = "plan2"
node_id = "node_id1"
status  = "enabled"
	[schedule.task]
	id     = "task_id2"
	name   = "task2"
	status = "enabled"
	
[[tasks]]
id     = "task_id1"
name   = "task1"
status = "enabled"

[[tasks]]
id     = "task_id2"
name   = "task2"
status = "enabled"
`

	config := new(Config)
	if _, err := toml.Decode(configStr, config); err != nil {
		// logger.Errorf(err.Error())
		t.Error(err)
	}
	t.Log(config)
}
