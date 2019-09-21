package main

const (
	version   = "0.1.0"
	bigBanner = `
...............................................................................
.. _____________________  _______________________  ________________________  ..
.. __  ___/_  ____/__  / / /__  ____/__  __ \_  / / /__  /___  ____/__  __ \ ..
.. _____ \_  /    __  /_/ /__  __/  __  / / /  / / /__  / __  __/  __  /_/ / ..
.. ____/ // /___  _  __  / _  /___  _  /_/ // /_/ / _  /___  /___  _  _, _/  ..
.. /____/ \____/  /_/ /_/  /_____/  /_____/ \____/  /_____/_____/  /_/ |_|   ..
.......................................................... version: %s .....
...............................................................................
`
	configStr = `
[node]
id    = "node_id1"
name  = "master1"
role  = "master-slave"
debug = false

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
id       = "plan_id1"
name     = "plan1"
node_id  = "node_id1"
interval = 10
status   = "enabled"
	[schedule.task]
	id     = "task_id1"
	name   = "task1"
	status = "enabled"

[[schedule]]
id       = "plan_id2"
name     = "plan2"
node_id  = "node_id1"
interval = 10
status   = "enabled"
	[schedule.task]
	id     = "task_id2"
	name   = "task2"
	status = "enabled"
`
// [[tasks]]
// id     = "task_id1"
// name   = "task1"
// status = "enabled"

// [[tasks]]
// id     = "task_id2"
// name   = "task2"
// status = "enabled"
)

