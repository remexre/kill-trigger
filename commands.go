package kt

// Command represents a single command given to the client by the server.
type Command struct {
	Name string
	ID   byte
}

// The command constants.
var (
	HelloWorld = Command{"HelloWorld", 0x00}
)

// Commands contains all the commands in a slice.
var Commands = []Command{
	HelloWorld,
}
