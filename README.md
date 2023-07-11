
# API Rest CRUD of User for Verify

Aplicação desenvolvida com o proposito de testar conhecimentos, será desenvolvido um CRUD de usuários com as tecnologias 

Golang 1.20 e MySQL 8.0.
Docker
Frameworks: GIN, GORM

As atividades, especificações e acompanhamento do projeto é possivel acompanhar através desse Link: [https://trello.com/b/BFL4WdlW/api-verifymy-crud](https://trello.com/b/BFL4WdlW/api-users-crud-verifymy)

A atividade foi iniciada com o intuito de ter uma curva de aprendizagem na linguagem GO. Para isso iniciei um projeto teste em uma arquitetura MVC que tinha mais familiaridade para colocar em prática o desenvolvimento e iniciar alguns testes de conexão entre o Model e o banco MySql. Configuração de Docker e testes de Frameworks e bibliotecas para API Rest. Isso gerou o projeto teste abaixo para titulo de informação:

-Projeto Teste em MVC: https://github.com/mvzcanhaco/api-verifymy-crud-test

Esse projeto está inacabado e faltando diversas funcionalidades, somente foi gerado para titudo de estudo e incluido aqui, como outro modelo de arquitetura de mais simples construção, porém com alto acoplamento e dificuldade de expansão e troca de componentes. Após esse estudo inicial, planejei uma refatoração de arquitetura e foi criado esse projeto utilizando os conceitos de Clean Archtecture. 

## Arquitetura:

A arquitetura utilizada nesse projeto é conhecida como Clean Architecture, que é um estilo de arquitetura de software que visa separar e organizar as responsabilidades de um sistema de forma clara e modular. Ela é baseada em princípios como a independência de frameworks, testabilidade e isolamento de dependências.

#### Principais Camadas Criadas:

1. Camada de Apresentação (UI): Essa camada é responsável por lidar com a interação do usuário, exibição de informações e coleta de entradas. Ela inclui componentes de interface do usuário, como telas, formulários ou páginas web. Na arquitetura Clean, essa camada não deve conter lógica de negócios complexa, apenas chama os serviços da camada de lógica de negócios para processar as ações do usuário. Pasta que representa essa camada no projeto é /delivery : Essa pasta contém os manipuladores HTTP (handlers) que lidam com as requisições HTTP, traduzem as entradas e chamam os serviços da camada de lógica de negócios. Ela representa a interface entre o usuário e o sistema.

2. Camada de Lógica de Negócios (Business Logic Layer - BLL): Essa camada contém a lógica central do sistema. Ela é responsável por implementar as regras de negócio, processar os dados, realizar cálculos, validações e qualquer outra operação relacionada à lógica do domínio. A camada de lógica de negócios é independente de qualquer detalhe de implementação, como banco de dados ou tecnologias externas. Pasta /usecase: Essa pasta contém as regras de negócio e a lógica central do sistema. Os arquivos dentro dessa pasta representam os casos de uso (use cases) do sistema. Eles encapsulam as operações relacionadas a um determinado contexto ou funcionalidade do sistema.
   
3. Camada de Domínio (Domain Layer): Essa camada contém as entidades de negócio, agregados, serviços e objetos de valor que representam o núcleo do domínio da aplicação. Pasta domain/entity: Essa pasta contém as definições das entidades do sistema, que representam os objetos de negócio. Elas podem ser estruturas de dados simples ou structs que representam conceitos mais complexos. Pasta domain/utils: Essa pasta contém funções impotante para compor, definir, validar e gerar regras de negocios exclusivas para as entidades. 

4. Camada de Acesso a Dados (Data Access Layer - DAL): Essa camada é responsável por acessar e persistir os dados do sistema. Ela inclui as operações de leitura e gravação no banco de dados ou em outras fontes de dados. A camada de acesso a dados abstrai os detalhes de como os dados são armazenados, permitindo que a lógica de negócios se concentre apenas na manipulação dos dados. Pasta /repository: Essa pasta contém as implementações dos repositórios, que são responsáveis por acessar e persistir os dados. Os repositórios definem as operações de leitura e gravação no banco de dados ou em outras fontes de dados.
  
5. Camada de Infraestrutura (Infrastructure Layer): Essa camada é responsável por implementar os detalhes técnicos e infraestruturais, como o acesso a bancos de dados, envio de e-mails, chamadas a APIs externas, entre outros. Pasta /db: Essa pasta contém os arquivos relacionados à configuração e conexão com o banco de dados.      

#### Vantagens de Ter Utilizado Essa Arquitetura:

1. Separação de Responsabilidades: A arquitetura Clean promove uma separação clara das responsabilidades em diferentes camadas. Isso torna o sistema mais organizado, modular e facilita a manutenção e evolução do código.

2. Testabilidade: A separação das camadas e a dependência de inversão facilitam a criação de testes automatizados. Cada camada pode ser testada de forma isolada, permitindo testes unitários eficientes e facilitando a detecção de erros.

3. Independência de Frameworks: A arquitetura Clean incentiva a independência de frameworks e tecnologias específicas. Isso torna o sistema menos acoplado a bibliotecas externas e mais flexível para trocar ou atualizar componentes.

4. Facilidade de Manutenção: Com a estrutura clara e bem definida, a arquitetura Clean facilita a manutenção do código ao longo do tempo. As alterações em uma camada não devem afetar as outras, desde que as interfaces entre as camadas sejam respeitadas.

## Para Executar

Para executar o projeto é necessário ter o docker e docker compose instalados e executar o seguinte comando:
```
docker compose up --build
```

Ao iniciar o MySql pela primeira vez, é criado um usuário administrador inicial para utilizar a aplicação.  
As credenciais são:
```
EMAIL = admin@example.com
PASSWORD = 123456
```

O arquivo API-VerifyMy-CRUD_2023-07-11.json, contém uma collection gerada no Insomnia para testar os principais endpoints.

## Endpoints ```/api/v1```

#### POST ```/users```
Para cadastrar usuário pode usar essa rota, sem nenhum uso de autenticação. Seria uma rota aberta ao publico para se cadastrar a plataforma, porém essa rota somente cria usuários comuns (profile:user)

**Body:**
```
{
  "name": "Elon Musk",
  "email": "elonmusk@example.com",
  "password": "password123",
  "birthDate": "1971-06-28",
  "address": {
    "street": "123 Tesla Ave",
    "city": "Los Angeles",
    "state": "CA",
    "country": "USA"
  }
}
```

#### POST ```/login```
Cria uma sessão autenticada (login) e retorna o token de acesso para as rotas GET, UPDATE e DELETE

**Body:**
```
{
    "email": "admin@example.com",
    "password": "123456"
}
```
*Se utilizar o login Inicial acima, é gerado um token com Perfil "admin". Esse tem acesso a todas as rotas. Se caso logar com outro usuario cadastrado esse terá acesso "user" e somente terá acesso as rotas GET


#### GET ```/users/:id```
Obtém usuário a partir do seu ID.   
*É necessário enviar o token gerado em ```POST /login``` como Bearer Token.
**Perfis 'user' podem visualizar os usuários

#### GET ```/users```
Obtém usuários. Possuí paginação.   
*É necessário enviar o token gerado em ```POST /login``` como Bearer Token.  
**Perfis 'user' podem visualizar os usuários

**Exemplo:**
```
/users?page=1&pageSize=5
```

#### PATCH ```/users/:id```
Atualiza usuário a partir de seu ID.   
*É necessário enviar o token gerado em ```POST /login``` como Bearer Token.  
** Somente  Perfil 'admin' podem atualizar os usuários

**Body:**
```
{
    "name": "Nome Usuário"
}
```

#### DELETE ```/users/:id```
Deleta usuário a partir de seu ID.
*É necessário enviar o token gerado em ```POST /login``` como Bearer Token.  
** Somente  Perfil 'admin' podem deletar os usuários
OBS.: É possível deletar o próprio usuário.

### Testes:

Os testes rodam com o seguinte comando:
```
go test -v -coverprofile cover.out ./...
go tool cover -html cover.out 
```
Abra o arquivo gerado ```cover.html``` no navegador para checar a cobertura.


### Em Construção: 
Pode ser acompanhado em: https://trello.com/b/BFL4WdlW/api-users-crud-verifymy

Swager

Versão Release

Testes Automatizados

Logs estruturados

Monitoramento

Melhorias Segurança
