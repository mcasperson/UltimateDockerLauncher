# Ultimate Docker Launcher
![Coverage](https://img.shields.io/badge/Coverage-91.2%25-brightgreen)

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

## Quick Env Var Reference

* `UDL_WRITEFILE[FILENAME]`: Writes a file e.g. `UDL_WRITEFILE[/etc/myapp/config.json]`
* `UDL_SETVALUE[FILENAME][KEY]`: Sets a value in a config file e.g. `UDL_SETVALUE[/etc/myapp/config.json][entry2:entry3]` or `UDL_SETVALUE[/etc/myapp/config.yaml][entry2:entry3:0]`

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

## Manipulating files

UDL understands a number of file formats, including:

* JSON
* YAML
* TOML
* XML (Planned)
* INI (Planned)

The values assigned to the environment variables in the format `UDL_SETVALUE[FILENAME][KEY]`  are inserted into the file
`FILENAME` creating or overwriting the value found at `KEY`. 

The format of `KEY` depends on the file being edited:

* JSON, YAML: Key is a colon seperated path e.g. `first` or `first:second`. Integer values are used to index into an array e.g. `first:second:0`.
* XML: Key is an xpath
* INI: Key is a colon separated path with optional group e.g. `value` or `group:value`

For example, given a JSON file like this at `/etc/myapp/config.json`:

```
{
    "entry1": "value1"
    "entry2": {
        "entry3": "value2"
    },
    "entry4": ["value3", "value4"]
}
```

* `UDL_SETVALUE[/etc/myapp/config.json][entry1]` replaces `value1`
* `UDL_SETVALUE[/etc/myapp/config.json][entry2:entry3]` replaces `value2`
* `UDL_SETVALUE[/etc/myapp/config.json][entry4:1]` replaces `value4`
