package components

type Component interface {
	Render() string // Render the component
	String() string
	Html() string
	Iam() (Component, string)
}