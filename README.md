# json2env
Simple CLI tool to export JSON data as environment variables

## Installation
### Linux/macOS
```
curl -sSL https://github.com/yamaszone/json2env/releases/download/v0.1.0/json2env-v0.1.0-$(
    bash -c '[[ $OSTYPE == darwin* ]] && echo darwin || echo linux'
  )-amd64 -o json2env && chmod a+x json2env && sudo mv json2env /usr/local/bin/
```
### Windows
Download executable from [releases page](https://github.com/yamaszone/json2env/releases/tag/v0.1.0)


## Use Case
`json2env -h`
```
Usage:
  json2env [flags]

Aliases:
  json2env, j2e

Flags:
  -e, --exportable     Print output as 'export ENV_VAR=value'
  -h, --help           help for json2env
  -i, --input string   Input JSON file name
```

### Examples
Given a JSON file `input.json`:
```
{
	"ENV_VAR1": "value1",
	"ENV_VAR2": "value2"
}
```

#### Example #1: Set as variables in your shell

	$ eval $(json2env -i input.json)
	$ echo $ENV_VAR1
	value1
	$ echo $ENV_VAR2
	value2

#### Example #2: Export as environment variables in your shell

	$ eval $(json2env -i input.json -e)
	$ echo $ENV_VAR1
	value1
	$ echo $ENV_VAR2
	value2

#### Example #3: Get output in .env format

	$ json2env -i input.json
	ENV_VAR1=value1
	ENV_VAR2=value2

#### Example #4: Get output to source in your shell

	$ json2env -i input.json -e
	export ENV_VAR1=value1
	export ENV_VAR2=value2

