package keyvalueeditor

type KeyValueEditorConfig struct {
	Title            string
	KeyPlaceholder   string
	ValuePlaceholder string
	InitialRows      int
	Width            int // it is not the editor container width, this width will be divided by 2 and assigned to key and value input width
}

func NewKeyValueEditorConfig(
	title string,
	keyPlaceholder string,
	valuePlaceholder string,
	initialRows int,
	width int,
) KeyValueEditorConfig {
	return KeyValueEditorConfig{
		Title:            title,
		KeyPlaceholder:   keyPlaceholder,
		ValuePlaceholder: valuePlaceholder,
		InitialRows:      initialRows,
		Width:            width,
	}
}
