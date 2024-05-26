package config

type ParameterSection struct {
	Text       string
	Parameters []Parameter
}

type Parameter struct {
	Template    string
	Parameter   string
	Units       string
	SliderInt   *SliderInt
	SliderFloat *SliderFloat
	Toggle      *Toggle
}

type SliderInt struct {
	Min int
	Max int
}

type SliderFloat struct {
	Min       float64
	Max       float64
	Precision float64
}

type Toggle struct{}
