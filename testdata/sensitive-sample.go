func connect()  {
	user := "root"
	password:= "supersecret" // Sensitive
  
	url := "login=" + user + "&passwd=" + password
  }

  func run() {
	prepare("This should be a constant")  // Noncompliant; 'This should ...' is duplicated 3 times
	execute("This should be a constant")
	release("This should be a constant")
}

func fun1() (x, y int) {
	a, b := 1, 2
	b, a = a, b
	return a, b
  }
  
  func fun2() (x, y int) {  // Noncompliant; fun1 and fun2 have identical implementations
	a, b := 1, 2
	b, a = a, b
	return a, b
  }

  var (
	ip   = "192.168.12.42"
	port = 3333
  )
  
  SocketClient(ip, port)

  if true {
	doSomething()
}

if false {
	doSomething()
}