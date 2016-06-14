package kt

// Command represents a single command given to the client by the server.
type Command struct {
	Name string
	ID   byte
}

// The command constants.
var (
	KeepAlive  = Command{"KeepAlive", 0x00}
	HelloWorld = Command{"HelloWorld", 0x01}
	KillJava   = Command{"KillJava", 0x02}

	Ping = Command{"Ping", 0xFE}
	Pong = Command{"Pong", 0xFF}
)

// Commands contains all the commands in a slice.
var Commands = []Command{
	KeepAlive,
	HelloWorld,
	KillJava,

	Ping,
	Pong,
}
