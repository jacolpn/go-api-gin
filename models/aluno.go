package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

/*
gorm.Model:

	ID         uint            `gorm: "primaryKey`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt  `gorm: "index"`
*/
type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	RG   string `json:"rg" validate:"len=9,regexp=^[0-9]*$"`
	CPF  string `json:"cpf" validate:"len=11,regexp=^[0-9]*$"`
}

// var Alunos []Aluno

func ValidaDadosDeAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}

	return nil
}
