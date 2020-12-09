package core

type Size int64

const (
	BYTE Size = 1
	KB   Size = BYTE << 10
	MB   Size = KB << 10
	GB   Size = MB << 10
	TB   Size = GB << 10
)

func (s Size) Add(arg Size) Size {
	return s + arg
}

func (s Size) Del(arg Size) Size {
	return s - arg
}

func (s Size) TB() Size {
	return s / TB
}

func (s Size) GB() Size {
	return s / GB
}

func (s Size) MB() Size {
	return s / MB
}

func (s Size) KB() Size {
	return s / KB
}
