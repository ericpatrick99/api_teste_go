CREATE TABLE IF NOT EXISTS usuario(
    id BIGINT not null PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    name VARCHAR(255) not null COMMENT 'Nome do usuário',
    email VARCHAR(100) not null COMMENT 'email',
    nick_name VARCHAR(100) not null COMMENT 'apelido',
    password VARCHAR(255) not null COMMENT 'senha',
    created_at DATETIME COMMENT 'data cadastro' DEFAULT now(),
    updated_at DATETIME COMMENT 'data de atualização' DEFAULT now(),
    unique (email, nick_name, password) COMMENT'garante que só terá um email e senha cadastrado',
    unique (email),
    unique (nick_name)
    ) COMMENT 'tabela que registra os usuários';