package encoding

type IEncoder interface {
	PushNamespace(namespaceUri string)
	PopNamespace() string

	WriteBoolean(fieldName string, value bool)
	WriteByte(fieldName string, value byte)
}
