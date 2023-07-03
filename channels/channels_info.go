package channels

/*
	Channels are a built-in language construct used for communication and
	synchronization between goroutines(concurrently executing functions or methods).
	Channels provide a way for goroutines to send and receive valuesi allowing them to safely communicate
	and coordinate their actions.

	Channels can be thoght of as pipes or conduits through which data flows.
	 They have a specific type associated with them, which defines the type of data that can be sent or received on the channel.
	 The type is specified using the channel operator 'chan' followed by the desired data type.
	 To create a channel, we use 'make()' func with the channel type as an argument.

	Channels have two main operations: sending and receiving. The '<-' operator is used for
	 sending and receiving values on a channel. When sending a value to a channel,
	 the value is placed into the channel using the syntax 'channelName <- value'.
	 When receiving a value from channel, the syntax is ' variable := <- channelName'.

	Channels can also be used for synchronization. Ny default, sending and receiving opearations on a channel block until both the sender and receiver are ready.
	 This allows goroutines to synchronize their execution. For example, a goroutine may wait
	 for a value to be sent on a channel before proceeding.

	Go also supports buffered channels, which have a capacity and allow multiple values to be sent awithout an immediate receiver.
	 Buffered channels provide a level of decoupling between senders and receivers, allowing them to opeate at different speeds.

	Channels are a powerful mechanism for cordinating concurrent operation in Go,
	 enabling safe communication and synchronization between goroutines.


*/
