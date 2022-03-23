Jinx: a jinkies coordination tool

Usage:

jinkies serve start - starts a jinkies on localhost 😯
jinkies serve start -o hostconfig.yml - the `-o` flag allows the end user to specify any of the docker engine [https://pkg.go.dev/github.com/docker/docker@v20.10.12+incompatible/api/types/container#HostConfig] hostconfig options  😯
jinkies serve start -c containerconfig.yml - the `-c` flag allows the end user to specify any of the docker engine [https://pkg.go.dev/github.com/docker/docker@v20.10.12+incompatible/api/types/container#Config] containerconfig options  😯

Examples of how these yml files are structured are in the examples/ directory.