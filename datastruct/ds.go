package datastruct

type Port struct {
	Value int
	Ch    chan Reqest
}

type Reqest struct {
	Value int
	Resp  chan Response
}

type Response struct {
	Error error
	Value int
}
