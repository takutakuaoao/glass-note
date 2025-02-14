package domain

type NoteValidator struct {
	title string
}

func NewNoteValidator(title string) *NoteValidator {
	return &NoteValidator{
		title: title,
	}
}

func (v NoteValidator) Validate() bool {
	return v.title != ""
}
