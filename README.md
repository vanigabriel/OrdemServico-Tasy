# OrdemServico-Tasy
 API para criar uma OS dentro do Tasy.
 Implementado em Golang e utilizando o framework Gin.

# Documentação API
https://documenter.getpostman.com/view/7995657/SVn2PbnR

# Características
- Tratativa de erros;
- Busca de usuário por CPF;
- Possibilidade de inserir vários anexos;
- Docker;
- Princípios de Clean Architecture;
- Govalidator para conferir campos obrigatórios;
- Log enviado para o Timber.io;
- TDD;
- Funcionamento com o banco de dados Oracle;

# Rotas
- POST /ordemservico : insere ordem de serviço;
- POST /ordemservico/:os/files : insere anexos na ordem de serviço os;

# Atenção:
 Necessário preencher o .env com os dados de conexão com o banco de dados;
 Necessário criar a procedure que se encontra no /sql para que o sistema funcione;
 Para funcionar com o Oracle, é necessário ter 3 arquivos baixados para conexão:
- oracle-instantclient19.3-basic-19.3.0.0.0-1.x86_64.rpm 
- oracle-instantclient19.3-sqlplus-19.3.0.0.0-1.x86_64.rpm 
- oracle-instantclient19.3-devel-19.3.0.0.0-1.x86_64.rpm 
 Os mesmos se encontram no seguinte link: http://yum.oracle.com/repo/OracleLinux/OL7/oracle/instantclient/x86_64/index.html
 
# Backlog
- Usuário/Setor/Equipamento padrão quando não localizar o do CPF;
