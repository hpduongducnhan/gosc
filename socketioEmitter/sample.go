package socketioemitter

func ExampleSentMessage() {
	// socketioemitter.SetupEmitter("redis://localhost:6379/0")
	e, err := NewEmitter("redis://localhost:6379/0", "socket.io")
	if err != nil {
		panic(err)
	}

	newMessage := NewWsMessage("test-event", "hello", "/")
	newMessage.ToRoom("shared-room")
	e.Publish(*newMessage)
}

func ExampleSentMessageToRoom() {
	// socketioemitter.SetupEmitter("redis://localhost:6379/0")
	e, err := NewEmitter("redis://localhost:6379/0", "socket.io")
	if err != nil {
		panic(err)
	}

	newMessage := NewWsMessage("test-event", "hello", "/")
	newMessage.ToRoom("shared-room")
	e.Publish(*newMessage)
}
