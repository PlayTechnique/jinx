package jinxtypes

//Global jinx cli options useful across multiple commands.
//ContainerName the name of a container that the current command should operate against.
//PullImages should an image be pulled, or not? If not, we are going to assume the image already
//exists in the current docker image cache.
type JinxData struct {
	ContainerName string
	PullImages    bool
}
