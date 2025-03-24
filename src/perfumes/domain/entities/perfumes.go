package entities

type Perfume struct {
	ID          int32
	Name        string
	Descripcion string
}

// Constructor
func NewPerfume(name string, descripcion string) *Perfume {
	return &Perfume{Name: name, Descripcion: descripcion}
}

// Getters y Setters
func (p *Perfume) GetName() string {
	return p.Name
}

func (p *Perfume) SetName(name string) {
	p.Name = name
}

func (p *Perfume) GetDescripcion() string {
	return p.Descripcion
}

func (p *Perfume) SetDescripcion(descripcion string) {
	p.Descripcion = descripcion
}
