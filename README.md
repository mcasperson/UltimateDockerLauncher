# Ultimate Docker Launcher
![Coverage](https://img.shields.io/badge/Coverage-85.2%25-brightgreen)
[![Github All Releases](https://img.shields.io/github/downloads/mcasperson/UltimateDockerLauncher/total.svg)]()

Customizing Docker configuration files for an environment can be a pain with legacy applications that do not have the ability
to be fully configured via environment variables. Some container orchestration platforms, like Kubernetes, provide
the limited ability to mount files at runtime without the need to expose a complete networked file system. Platforms
like ECS only allow files to be mounted from file systems like EFS. Other more basic container platforms like
AppRunner don't allow any kind of file mounting.

So, depending on your platform, you may be forced to set up expensive and complicated network file shares for the sake
of mounting a few kilobytes of config files, or you may be out of luck.

A common workaround is to use something like `envsubst` as part of the container execution. This works, to a point.
As soon as you need to inject values into structured configuration files like YAML or JSON, you must take care to
properly escape the injected values. Using `envsubst` is also a all-or-nothing proposition, meaning you can not have
default configuration files that are optionally customized.

Ultimate Docker Launcher provides the ability to write or modify configuration files during container initialization.
It works by scanning environment variables for known patterns indicating files that need to be (over)written or
modified, and then executes a wrapped executable.

## Project Goals

The aim of this project is to allow Linux based containers to add or modify configuration files when a container is
launched. All configuration is done via environment variables, which are supported by every container orchestration
platform.

## Project Non-goals

* This project does not support Windows containers.
* This project does not aim to create a standalone tool for use outside of containers.

## Quick Env Var Reference

There are two styles of environment variables. The first style embeds the file ane key information in the environment
variable name:

* `UDL_WRITEFILE[FILENAME]`: Writes a file e.g. `UDL_WRITEFILE[/etc/myapp/config.json]` with a value of `{"whatever": ["hello"]}`.
* `UDL_WRITEB64FILE[FILENAME]`: Writes a base64 encoded value to a file e.g. `UDL_WRITEB64FILE[/etc/myapp/config.json]` with a value of `e3doYXRldmVyOiBbaGVsbG9dfQo=`
* `UDL_SETVALUE[FILENAME][KEY]`: Sets a value in a config file e.g. `UDL_SETVALUE[/etc/myapp/config.json][entry2:entry3]` or `UDL_SETVALUE[/etc/myapp/config.yaml][entry2:entry3:0]` with a value of `newvalue`

The second style is useful for Kubernetes, which only supports alphanumberic characters, the dot, the dash, and the 
underscore in environment variable names. The filename and key is located in the environment variable value:

* `UDL_WRITEFILE_IDENTIFIER`: Writes a file e.g. `UDL_WRITEFILE_blah` with a value of `[/etc/myapp/config.json]{"whatever": ["hello"]}`.
* `UDL_WRITEB64FILE_IDENTIFIER`: Writes a base64 encoded value to a file e.g. `UDL_WRITEB64FILE_blah` with a value of `[/etc/myapp/config.json]e3doYXRldmVyOiBbaGVsbG9dfQo=`.
* `UDL_SETVALUE_IDENTIFIER`: The file name and accessor are defined in the env var value e.g. `UDL_SETVALUE_whatever` with a value of `[/etc/myapp/config.json][entry2:entry3]newvalue` sets the value of the property under `entry2.entry3` to `newvalue`.

`IDENTIFIER` in the examples above is any string with alphanumeric characters, underscores, dashes, or periods. 
The `INDENTIFIER` has no meaning, and is simply used to allow unique env vars to be defined.
   

## Quick Key Reference

### JSON, YAML, and TOML 
* Keys are colon separated path accessors e.g. `value` in the JSON blob `{"top": {"second": {"third": "value"}}}` is accessed via `top:second:third`.
* Array items are accessed with a zero based index e.g. `value` in the JSON blob `{"top": {"second": ["value"]}}` is accessed via `top:second:0`

### INI
* Keys reference the top level INI property, or are colon separated group and property e.g. `property` or `group:property`

## Docker CMD Example

Save the following to `Dockerfile`:

```dockerfile
FROM python:3

RUN apt-get update; apt-get install -y jq curl

# Download the latest version of udl
RUN curl -s https://api.github.com/repos/mcasperson/UltimateDockerLauncher/releases/latest | \
    jq '.assets[] | select(.name|match("udl$")) | .browser_download_url' | \
    xargs -I {} curl -L -o /opt/udl {}
RUN chmod +x /opt/udl

# UDL_WRITEFILE[filename] environment variables are used to save files
ENV UDL_WRITEFILE[/app/config.json]='{"whatever": ["hello"]}'

# UDL_SETVALUE[file][key] environment variables are used to set values inside configuration files like JSON, YAML, INI etc
ENV UDL_SETVALUE[/app/config.json][whatever:0]="world"

RUN mkdir /app
RUN printf 'f = open("/app/config.json", "r") \n\
print(f.read())' >> /app/main.py

# The entrypoint or CMD is set to udl. The first argument is the application to run. The second and all subsequent
# arguments are passed to the application defined in the first argument.
CMD [ "/opt/udl", "python", "/app/main.py" ]
```

Here is another example Dockerfile using UDL with Apache, this time using bash to execute UDL before the main app:

```dockerfile
FROM httpd:2.4

RUN apt-get update; apt-get install -y jq curl

# Download the latest version of udl
RUN curl -s https://api.github.com/repos/mcasperson/UltimateDockerLauncher/releases/latest | \
    jq '.assets[] | select(.name|match("udl$")) | .browser_download_url' | \
    xargs -I {} curl -L -o /opt/udl {}
RUN chmod +x /opt/udl

# UDL_WRITEFILE[filename] environment variables are used to save files
ENV UDL_WRITEFILE[/usr/local/apache2/htdocs/config.json]='{"whatever": ["hello"]}'

# UDL_SETVALUE[file][key] environment variables are used to set values inside configuration files like JSON, YAML, INI etc
ENV UDL_SETVALUE[/usr/local/apache2/htdocs/config.json][whatever:0]="world"

# Here we use bash to call UDL before calling the main application
CMD [ "/bin/bash", "-c", "/opt/udl; httpd-foreground" ]
```

Build the image with:

```bash
docker build . -t udltest
```

Run the image with:

```bash
docker run udltest
```

## Writing files

The values assigned to environment variables in the format `UDL_WRITEFILE[FILENAME]` are written to the file `FILENAME`.
For example, UDL will save the contents of the environment variable `UDL_WRITEFILE[/etc/myapp/config.json]` to
the file `/etc/myapp/config.json` during initialization.

To write complex files, use an env var with the format `UDL_WRITEB64FILE[FILENAME]`, which decodes the base64 value
assigned to it and writes it to a file.

## Manipulating files

UDL understands a number of file formats, including:

* JSON
* YAML
* TOML
* INI
* XML (not implemented yet)

The values assigned to the environment variables in the format `UDL_SETVALUE[FILENAME][KEY]`  are inserted into the file
`FILENAME` creating or overwriting the value found at `KEY`. 

The format of `KEY` depends on the file being edited:

* JSON, YAML: Key is a colon seperated path e.g. `first` or `first:second`. Integer values are used to index into an array e.g. `first:second:0`.
* XML: Key is an xpath
* INI: Key is a colon separated path with optional group e.g. `value` or `group:value`

For example, given a JSON file like this at `/etc/myapp/config.json`:

```json
{
    "entry1": "value1",
    "entry2": {
        "entry3": "value2"
    },
    "entry4": ["value3", "value4"]
}
```

* `UDL_SETVALUE[/etc/myapp/config.json][entry1]` set to `newvalue` replaces `value1` with `newvalue`
* `UDL_SETVALUE[/etc/myapp/config.json][entry2:entry3]` set to `newvalue` replaces `value2` with `newvalue`
* `UDL_SETVALUE[/etc/myapp/config.json][entry4:1]` set to `newvalue` replaces `value4` with `newvalue`
* `UDL_SETVALUE_1` with a value of `[/etc/myapp/config.json][entry1]newvalue` replaces `value1` with `newvalue`
* `UDL_SETVALUE_WHATEVER` with a value of `[/etc/myapp/config.json][entry2:entry3]newvalue` replaces `value2` with `newvalue`
* `UDL_SETVALUE_ANY_STRING-WITH.ALPHA_NUMERIC.CHARS-DASHES_OR.UNDERSCORES` with a value of `[/etc/myapp/config.json][entry4:1]newvalue` replaces `value4` with `newvalue`

## Type retention

Where possible, the type of the replaced value is retained. Numbers, strings, booleans, arrays, and objects are 
recognized.

Where the replacement value is unable to be cast to the value at the destination, the replacement value is inserted
as a string.

Given the following JSON blob at `/etc/myapp/config.json`

```json
{
    "entry1": "value1",
    "entry2": {
        "entry3": true
    },
    "entry4": [1, 2]
}
```

Then setting the env var `UDL_SETVALUE[/etc/myapp/config.json][entry1]` with a value of `test` results in:

```JSON
{
    "entry1": "test",
    "entry2": {
        "entry3": true
    },
    "entry4": [1, 2]
}
```

The env var `UDL_SETVALUE[/etc/myapp/config.json][entry2:entry3]` with a value of `false` results in:

```json
{
    "entry1": "test",
    "entry2": {
        "entry3": false
    },
    "entry4": [1, 2]
}
```

The env var `UDL_SETVALUE[/etc/myapp/config.json][entry4:1]` with a value of `10` results in:

```json
{
    "entry1": "test",
    "entry2": {
        "entry3": false
    },
    "entry4": [1, 10]
}
```

The env var `UDL_SETVALUE[/etc/myapp/config.json][entry2]` with a value of `{"entry5": "value1"}` results in:

```json
{
    "entry1": "test",
    "entry2": {
        "entry5": "value1"
    },
    "entry4": [1, 2]
}
```

The env var `UDL_SETVALUE[/etc/myapp/config.json][entry4]` with a value of `["value6", "value7"]` results in:

```json
{
    "entry1": "test",
    "entry2": {
        "entry5": "value1"
    },
    "entry4": ["value6", "value7"]
}
```
