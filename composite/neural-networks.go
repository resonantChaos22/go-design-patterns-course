package main

type NeuronInterface interface {
	Iter() []*Neuron
}

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}
func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

type NeuronLayer struct {
	Neurons []Neuron
}

func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for _, neuron := range n.Neurons {
		result = append(result, &neuron)
	}

	return result
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{
		Neurons: make([]Neuron, count),
	}
}

func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}

func TestNeuralNetwork() {
	n1, n2 := &Neuron{}, &Neuron{}
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	Connect(n1, n2)
	Connect(n1, layer1)
	Connect(layer2, n1)
	Connect(layer1, layer2)
}
