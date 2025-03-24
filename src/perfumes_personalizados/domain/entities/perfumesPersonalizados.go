package entities

type PerfumePersonalizado struct {
	ID int32
	Name string
	Descripcion string
	Ingredientes string
	Fecha_creacion string
}

func NewPerfumePersonalizado (name string, descripcion string, ingredientes string , fecha_creacion string)*PerfumePersonalizado{
	return&PerfumePersonalizado{Name:name, Descripcion: descripcion, Ingredientes: ingredientes, Fecha_creacion: fecha_creacion}
}


func (pp*PerfumePersonalizado)GetName()string{
	return pp.Name
}

func (pp*PerfumePersonalizado)SetName(name string){
	pp.Name = name
}

func (pp*PerfumePersonalizado)GetDescripcion()string{
	return pp.Descripcion
}

func (pp*PerfumePersonalizado)SetDescripcion(descripcion string){
	pp.Descripcion = descripcion
}

func (pp*PerfumePersonalizado)GetIngredientes()string {
	return pp.Ingredientes
}

func (pp*PerfumePersonalizado)SetIngredientes(ingredientes string){
	pp.Ingredientes = ingredientes
}

func (pp*PerfumePersonalizado)GetFecha_creacion()string{
	return pp.Fecha_creacion
}

func (pp*PerfumePersonalizado)SetFecha_creacion(fecha_creacion string){
	pp.Fecha_creacion = fecha_creacion
}