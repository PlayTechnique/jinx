Jinx: a containerised Jenkins coordination tool

# It's too easy to let Jenkins get crufty
Jenkins is a beast! Most of us run it from the UI, backing up the jenkins data folder and hoping that a restore will
save us when we need to rebuild. We configure a plugin once, forget how it was done and then are frightened to upgrade
because something might get misconfigured, so plugins versions get old and out of date.

## The Solution is CI/CD
There's a better way: use YAML with JCASC where possible, use groovy to programatically configure plugins where it JCASC
doesn't give you the right tools. All the state for configuring Jenkins is now in version control, allowing for a CI/CD
approach where new versions of Jenkins trigger a rebuild, and a rebuild guides plugin upgrades.

## There's a lot of heavy lifting to get that working
Jinx is a cli tool that takes the heavy lifting out of the dev workflow, so Jenkins can live in CI/CD.
- How do you get plugins to have autocomplete in your IDE, so you're supported in writing groovy to configure them? Jinx
can inspect a running container, see what plugins you added, and get autocomplete working for intellij idea.
- Where do you put Jenkins startup code anyway? Jinx provides startup code and preconfigures your Jenkins with
info about where to find it. Follow the patterns!
- 


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
