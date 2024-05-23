package repositorios

import (
	"api_teste_go/src/modelos"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuario(db *sql.DB) *Usuarios {
	return &Usuarios{db: db}
}
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuario(name, nick_name, email, password) values (?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Name, usuario.NickName, usuario.Email, usuario.Password)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil
}
