package renderer

// Renderer delegates the call to a passed strategy
type Renderer struct{}

// Call renders the data
func (r Renderer) Call(strategy base, data [][]interface{}) {
	strategy.call(data)
}
