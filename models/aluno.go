package models

import "gorm.io/gorm"

/*
	gorm.Model:
		ID         uint            `gorm: "primaryKey`
		CreatedAt  time.Time
		UpdatedAt  time.Time
		DeletedAt  gorm.DeletedAt  `gorm: "index"`
*/
type Aluno struct {
	gorm.Model
	Nome string `json: "nome"`
	CPF  string `json: "cpf"`
	RG   string `json: "rg"`
}

// var Alunos []Aluno
