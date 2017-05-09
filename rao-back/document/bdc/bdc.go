package bdc

import "github.com/Zenika/RAO/document"

type BdcDocument struct {
	document.IDocument
	document.BusinessDocument
	Projet     string
	Consultant string
}

func (bdc BdcDocument) GetConsultant() string {
	return bdc.Consultant
}

func (bdc BdcDocument) SetConsultant(c string) {
	bdc.Consultant = c
}

func (bdc BdcDocument) GetProjet() string {
	return bdc.Projet
}

func (bdc BdcDocument) SetProjet(p string) {
	bdc.Projet = p
}
