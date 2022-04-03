Jinx: a jinkies coordination tool


Usage:
## Global default file
On startup, jinx defaults to assuming the container can be called `jinkies` and that you _do_ want to pull the latest
version of images from docker hub. This container can be tweaked:
- If a file named `jinx.yml` exists in the pwd of the calling process, it is read in and overrides the default configuration.
- If the `-j` option is supplied like `-j path/to/globalconfig.yml`, then this file is used for the global config. 

The options it controls are documented in examples/jinx.yml.

jinx -j path/to/jinx.yml - 

## `jinx serve`
`jinx serve start` - starts a jinkies on localhost 😯
`jinx serve stop`  - stops the jinkies on localhost

The `start`  subcommand further exposes all of the options that docker engine allows you to configure via two configuration
files, called the host config file and the container config file.

jinx serve start -o hostconfig.yml - the `-o` flag allows the end user to specify any of the docker engine [https://pkg.go.dev/github.com/docker/docker@v20.10.12+incompatible/api/types/container#HostConfig] hostconfig options
jinx serve start -c containerconfig.yml - the `-c` flag allows the end user to specify any of the docker engine [https://pkg.go.dev/github.com/docker/docker@v20.10.12+incompatible/api/types/container#Config] containerconfig options

Examples of how these yml files are structured are in the examples/ directory.

Config File:
For custom jinx options, we support a config file called jinx.yml. It must be present in the same directory as the jinx binary is invoked from.
The flags this supports are documented in the examples/ directory.
