package reuse

import (
	"container/list"
)

type MyList struct {
	l *list.List
}

func Create() MyList {
	return MyList{l: list.New()}
}

func (m *MyList) PushBack(elt any) *list.Element {
	return m.l.PushBack(elt)
}

func (m *MyList) PushFront(elt any) *list.Element {
	return m.l.PushFront(elt)
}

func (m *MyList) InsertBefore(elt any, mark *list.Element) *list.Element {
	return m.l.InsertBefore(elt, mark)
}

func (m *MyList) InsertAfter(elt any, mark *list.Element) *list.Element {
	return m.l.InsertAfter(elt, mark)
}