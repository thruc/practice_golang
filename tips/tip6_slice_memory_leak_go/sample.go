package main

func receiveMessage() []byte {
	return []byte("hello world")
}

func storeMessageType(msgType []byte) {
	// store message type
}

func consumeMessage() {
	for {
		msg := receiveMessage()
		storeMessageType(badGetMessageType(msg))
	}
}

// bad
func badGetMessageType(msg []byte) []byte {
	return msg[:5]
}

// good
func goodGetMessageType(msg []byte) []byte {
	msgType := make([]byte, 5)
	copy(msgType, msg)
	return msgType
}
