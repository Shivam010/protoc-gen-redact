package main

// ProtoFileData defines custom data type for Proto File info needed in template
type ProtoFileData struct {
	Source  string
	Package string
	// Imports: alias -> import-path
	Imports    map[string]string
	References []string
	Services   []*ServiceData
	Messages   []*MessageData
}

// ServiceData defines custom data type for Service info needed in template
type ServiceData struct {
	Name    string
	Skip    bool
	Methods []*MethodData
}

// MethodData defines custom data type for Method info needed in template
type MethodData struct {
	Name       string
	Skip       bool
	Input      string
	Output     *MessageData // will only contain name and options (ignore, nil, empty)
	Internal   bool
	StatusCode string
	ErrMessage string
}

// MessageData defines custom data type for Message info needed in template
type MessageData struct {
	Name      string
	WithAlias string

	Fields  []*FieldData
	Ignore  bool
	ToNil   bool
	ToEmpty bool
}

// FieldData defines custom data type for Field info needed in template
type FieldData struct {
	Name string
	// Redact using RedactionValue
	Redact         bool
	RedactionValue string

	IsMap      bool // IsMap: true for Map types
	IsRepeated bool // IsRepeated: true for Repeated types
	IsMessage  bool // IsMessage: true for Message type(& not Repeated/Map)

	// Iterate will only be used for Repeated/Map types and it specifies
	// whether or not to iterate each entry to be redacted
	Iterate bool

	// NestedEmbedCall will only be used for Message Types and it specifies
	// whether or not the embed message should be called for redaction.
	NestedEmbedCall bool

	// EmbedSkip will only be used for Message Types and it specifies
	// whether or not the embed message should be skipped.
	EmbedSkip bool

	// EmbedMessageName: name of embed message which is in case of Repeated or
	// Map or Message type field
	EmbedMessageName          string
	EmbedMessageNameWithAlias string
}
