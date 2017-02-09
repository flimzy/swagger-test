package foo

// Response is a response for foo
// swagger:response FooResponse
type Response struct {
	Tickets Tickets `json:"tickets"`
}

type Tickets []*Ticket

type Ticket struct {
	// The assignee
	AssignedTo *Foo `json:"assignedTo"`
}

type Foo struct{}
