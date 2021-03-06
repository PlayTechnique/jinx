* Jinx has a global configuration that's available to all commands. I'm trying to keep this small, but at a minimum
    if you are going to support custom container names then we need to be able to run commands directly against that
    custom name, so this is necessary.
* Some of the commands have a _lot_ of configuration options. In general, I prefer to supply an option for a config 
    file over having a specific flag per config option. This might change; kubernetes has tons of options, for
    example, and they're all exposed on the cli. For now, config files are simpler for me to reason about.
* Because commands need access to both the global runtime object and any custom options, each command has a custom
  runtime struct. Because it's nicer to not use free floating, global variables, each command is a method on this
  custom runtime struct. There's one or two small hoops to jump through implementing this, but it keeps everything
  pretty tidy and simple to reason about.
* Errors are returned from functions in order to facilitate testability. To make this work with cobra, use `RunE`, not
  `Run` to register your command.
* Remember that we're not even at version 0.5 yet. Lots of this thing is in flux. Improvements are beautiful and welcome.
* In general, I'm comfortable handling flag/argument validation inside the commands.
