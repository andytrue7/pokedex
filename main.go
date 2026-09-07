package main

func main() {
	config := &ReplStateConfig{
		commands: getCommands(),
	}

	startRepl(config)
}
