package main

type Conn struct {
	db *DB
}

func NewConn(db *DB) *Conn {
	return &Conn{db: db}
}

func (c *Conn) Open() {
	// do nothing
}

func (c *Conn) Query() int {
	return c.db.Query()
}
