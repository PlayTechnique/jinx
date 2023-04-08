Jinx: a jinkies coordination tool

# Quick Usage:

Usage:
## Global default file
On startup, jinx defaults to assuming the container can be called `jinkies` and that you _do_ want to pull the latest
version of images from docker hub. This container can be tweaked:
- If a file named `jinx.yml` exists in the pwd of the calling process, it is read in and overrides the default configuration.
- If the `-j` option is supplied like `-j path/to/globalconfig.yml`, then this file is used for the global config.

The options it controls are documented in examples/jinx.yml.

jinx -j path/to/jinx.yml -

## Quick docs
### `jinx serve`


`jinkies serve start` - starts a jinkies on localhost 😯

`jinkies serve start -o hostconfig.yml` - the `-o` flag allows the end user to specify any of the docker engine [hostconfig options](https://pkg.go.dev/github.com/docker/docker@v20.10.12+incompatible/api/types/container#HostConfig) 😯

`jinkies serve start -c containerconfig.yml` - the `-c` flag allows the end user to specify any of the docker engine [containerconfig options](https://pkg.go.dev/github.com/docker/docker@v20.10.12+incompatible/api/types/container#Config) 😯

Examples of how these yml files are structured are in the examples/ directory.


## Developers
Check out [Developer.md](DEVELOPER.md) for hints and tricks for getting started. We want your PR!
