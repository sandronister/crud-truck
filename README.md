# Crud-Truck

Crud-Truck é um projeto desenvolvido para fornecer uma solução robusta e escalável. Este documento fornece instruções sobre como configurar e executar o projeto localmente.

## Pré-requisitos

Antes de iniciar, certifique-se de ter o Docker instalado em sua máquina. Se você não tem o Docker instalado, visite [Docker](https://www.docker.com/get-started) para instruções de instalação.

# Configuração do Projeto

## Configuração do Arquivo .env

Não se esqueça de configurar o arquivo `.env` antes de iniciar o projeto. Este arquivo contém variáveis de ambiente essenciais para a configuração correta da aplicação, incluindo strings de conexão com o banco de dados, chaves de API, e outras configurações sensíveis e específicas do ambiente. Siga os passos abaixo para configurar o seu arquivo `.env`:

1. No diretório `./cmd `do projeto, tem um exemplo de arquivo env chamado `env.example` 
2. Renomeime para `.env` 
3. Abra o arquivo `.env` em um editor de texto de sua escolha.
4. Configure as variáveis de ambiente conforme as suas configurações


### Rodando as Migrações

Para preparar seu banco de dados, é necessário rodar o script SQL na pasta migrations em um banco de dados mysql


# Construindo a Imagem Docker
Para construir a imagem Docker do projeto a partir do Dockerfile, execute o seguinte comando no diretório raiz do projeto:

docker build -t gobrax .

Este comando construirá uma imagem Docker chamada gobrax baseada nas instruções fornecidas no Dockerfile.

# Rodando a Imagem Docker
Após construir a imagem, você pode rodar a aplicação executando:

docker run -p 8000:8000 gobrax

Este comando irá executar a aplicação em um contêiner Docker, mapeando a porta 8000 do contêiner para a porta 8000 do host.

# URLs Disponíveis

O projeto GOBRAX fornece várias URLs para interação com a aplicação. Abaixo estão detalhadas as operações disponíveis para cada endpoint:

### Motoristas

- `POST /api/v1/drivers` - Cadastra um novo motorista.
- `GET /api/v1/drivers` - Lista todos os motoristas.
- `GET /api/v1/drivers/{id}` - Retorna detalhes de um motorista específico.
- `PUT /api/v1/drivers/{id}` - Atualiza os dados de um motorista específico.
- `DELETE /api/v1/drivers/{id}` - Remove um motorista específico.

### Veículos (Caminhões)

- `POST /api/v1/trucks` - Cadastra um novo veículo.
- `GET /api/v1/trucks` - Lista todos os veículos.
- `GET /api/v1/trucks/{id}` - Retorna detalhes de um veículo específico.
- `PUT /api/v1/trucks/{id}` - Atualiza os dados de um veículo específico.
- `DELETE /api/v1/trucks/{id}` - Remove um veículo específico.

### Link (Vinculo)

- `POST /api/v1/links` - Vincula um caminhão a um motorista
- `GET /api/v1/liks/{id}` - Retorna uma lista de caminhões especifica de um motorista
- `DELETE` /api/v1/links/{driver_id}/{truck_id} - Remove o vinculo entre caminhão e motorista

Lembre-se de substituir `{id},{driver_id} e {truck_id}` pelo identificador real do motorista ou veículo que deseja consultar, atualizar ou remover. Estas URLs permitem uma gestão completa dos recursos de motoristas e veículos dentro da aplicação.
