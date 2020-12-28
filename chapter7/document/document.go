package document

type Document struct {
	ID     string  `json:"id"`
	Fields []Field `json:"fields"`
}

func NewDocument(id string) *Document {
	return &Document{
		ID:     id,
		Fields: make([]Field, 0),
	}
}

func (d *Document) AddField(f Field) *Document {
	d.Fields = append(d.Fields, f)
	return d
}