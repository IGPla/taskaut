# Taskaut

This tool aims to be a quick way to automate common tasks on your computer. Even more, it will provide you an easy format to paralelize or serialize all involved tasks in a given workflow.

## Format

The following is the valid format

```
[
	{
		"commands": [
			{
				"command": <here your command>,
				"params": [<param 1>, <param n>],
				"retries": <number of retries for the given command on fail>
			},
			{
				...
			}
		],
		"dir": <directory where you want to run your command list>
		"abort_on_fail": <true if you want to abort the rest of commands if one fail; false otherwise>
	},
	{
		...
	},
	...
]
```

```
---
- commands:
  - command: <here your command>
    params:
    - <param 1>
	- <param n>
    retries: <number of retries for the given command on fail>
  dir: <directory where you want to run your command list>
  abort_on_fail: <true if you want to abort the rest of commands if one fail; false otherwise>
  ...
...
```

## Example

You can find here a config file example

```json
[
    {
	"commands": [
	    {
		"command": "ls -ltrh",
		"retries": 1
	    },
	    {
		"command": "mkdir test",
		"retries": 0
	    }
	],
	"dir": "/tmp",
	"abort_on_fail": true
    },
    {
	"commands": [
	    {
		"command": "ls -ltrh",
		"retries": 5
	    }
	],
	"dir": "/etc",
	"abort_on_fail": false
    }
]
```

If you store the previous content in /tmp/myconfig.json, you can run it with the following command

```bash
taskaut -actions-file /tmp/myconfig.json
```

And it will run ls and mkdir in /tmp (serialized), and at the same time (paralel) will execute ls in /etc

Here you have the same config file in yaml

```yaml
---
- commands:
  - command: ls -ltrh
    retries: 1
  - command: mkdir test
    retries: 0
  dir: "/tmp"
  abort_on_fail: true
- commands:
  - command: ls -ltrh
    retries: 5
  dir: "/etc"
  abort_on_fail: false
```

