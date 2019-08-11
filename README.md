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
				"params": [<param1>, <paramn>]
			},
			{
				...
			}
		],
		"dir": <directory where you want to run your command list
	},
	{
		...
	},
	...
]
```

## Example

You can find here a config file example

```
[
    {
	"commands": [
	    {
		"command": "ls",
		"params": ["-ltrh"]
	    },
	    {
		"command": "mkdir",
		"params": ["test"]
	    }
	],
	"dir": "/tmp"
    },
    {
	"commands": [
	    {
		"command": "ls",
		"params": ["ltrh"]
	    }
	],
	"dir": "/etc"
    }
]
```

If you store the previous content in /tmp/myconfig.json, you can run it with the following command

```
taskaut /tmp/myconfig.json
```

And it will run ls and mkdir in /tmp (serialized), and at the same time (paralel) will execute ls in /etc
