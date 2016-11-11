goserver
--------------------------------------------------------
A server framework by registing command to implement one feature.


Usage
---------------------------------------------------------
```
	command.RegistCommand(command.NewCommand("hello", func(clnt *client.Client, args ...string) string {
		return "hello,I am a robot"
	},"say hello"))


	s:=goserver.NewServer()
	s.Start("127.0.0.1","5050")
	defer s.Stop()
```

Example
---------------------------------------------------------
There are some examples in the examples directory
