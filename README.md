# Ultimate Docker Launcher

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
* `UDL_SETVALUE[FILENAME][KEY]`: Sets a value in a config file e.g. `UDL_SETVALUE[/etc/myapp/config.json][entry2:entry3]`

## Docker CMD Example

Save the following to `Dockerfile`:

```
FROM python:3

COPY udl /opt
ENV UDL_WRITEFILE[/app/main.py]="print('hi')"

CMD [ "/opt/udl", "python", "/app/main.py" ]
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
* YAML (Planned)
* XML (Planned)
* INI (Planned)

The values assigned to the environment variables in the format `UDL_SETVALUE[FILENAME][KEY]`  are inserted into the file
`FILENAME` creating or overwriting the value found at `KEY`. the format of `KEY` depends on the file being edited.

For example, given a JSON file like this at `/etc/myapp/config.json`

```
{
    "entry1": "value1"
    "entry2": {
        "entry3": "value3"
    }
}
```

UDL will overwrite the value for the property `entry1` with the value in the environment variable
`UDL_SETVALUE[/etc/myapp/config.json][entry1]` and overwrite the value for the property `entry2.entry3` with the value 
in the environment variable `UDL_SETVALUE[/etc/myapp/config.json][entry2:entry3]`.