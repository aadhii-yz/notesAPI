package main

type Notes struct {
	data   map[int]Note
	nextID int
}

func newNotes() *Notes {
	return &Notes{
		data:   make(map[int]Note),
		nextID: 0,
	}
}

func (n *Notes) Append(newNote Note) Note {
	newNote.ID = n.nextID
	n.data[n.nextID] = newNote
	n.nextID++
	return newNote
}
