package commnd

type RemoteControl struct {
	control Command
}

func (r *RemoteControl) Press() {
	r.control.Execute()
}

func (r *RemoteControl) SetCommand(command Command) {
	r.control = command
}
