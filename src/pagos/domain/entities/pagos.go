package entities

type Pagos struct{
	ID int32
	Descripcion string
}

func NewPago(descripcion string) *Pagos{
	return &Pagos{ Descripcion:descripcion}
}

func (pay* Pagos) GetDescripcion()string{
	return pay.Descripcion
}

func (pay *Pagos) SetDescripcion(descripcion string){
	pay.Descripcion = descripcion
}