package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (this *Hombre) Respirar()    { this.respirando = true }
func (this *Hombre) Pensar()      { this.pensando = true }
func (this *Hombre) Comer()       { this.comiendo = true }
func (this *Hombre) Sexo() string { return "Hombre" }

// func (this *Hombre) SerVivo() bool { return true}
