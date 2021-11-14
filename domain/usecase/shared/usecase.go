package shared

type Reader interface {
	Read() (usecaseInput interface{})
}

type Writer interface {
	Write(usecaseOutput interface{})
}

type Usecase = func(reader Reader, writer Writer)
