package main

import "CommandPattern/src/command"

func main() {
	tv := &command.Tivi{}

	onCommand := &command.OnCommand{
		Device: tv,
	}

	offCommand := &command.OffCommand{
		Device: tv,
	}

	onButton := command.Button{
		Command: onCommand,
	}

	offButton := command.Button{
		Command: offCommand,
	}

	onButton.Press()
	offButton.Press()
}
