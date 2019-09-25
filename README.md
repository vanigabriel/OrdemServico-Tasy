# OrdemServico-Tasy
 API para criar uma OS dentro do Tasy.
 Implementado em Golang e utilizando o framework Gin.

# Documentação API
https://documenter.getpostman.com/view/7995657/SVn2PbnR

# Features
- Tratativa de erros;
- Princípios de Clean Architecture;
- Govalidator para conferir campos obrigatórios;
- Log salvo em arquivo;
- TDD;
- Funcionamento com o banco de dados Oracle;

# Rotas
- POST /ordemservico : insere ordem de serviço;

# Atenção:
- Necessário preencher o .env com os dados de conexão com o banco de dados;
- Necessário criar a procedure que se encontra no /sql para que o sistema funcione;
 
# Backlog
- Docker
- Possibilidade de inserir anexos, utilizando a configuração do Tasy;
