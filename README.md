# 📊 Poll Voting

## 🚀 Introdução

Inspirado pela NLW Expect da [Rocketseat](https://rocketseat.com.br), onde um projeto semelhante foi desenvolvido na trilha de NodeJS.
Poll Voting é uma api construída em [Golang](https://golang.org/) para votações em enquetes.
A intenção deste projeto é demonstrar a utilização de ferramentas e técnicas de desenvolvimento web em Golang, utilizando habilidades do desenvolvimento backend aprendidas utilizando NodeJS.

## 👨‍💻 Tecnologias

- [Golang](https://golang.org/): Linguagem de programação principal.
- [PostgreSQL](https://www.postgresql.org/): Banco de dados relacional para armazenar dados persistentes.
- [Redis](https://redis.io/): Armazenamento de cache para otimizar consultas frequentes.
- [Gin](https://gin-gonic.com/): Framework web para construir APIs em Golang.
- [Swagger](https://swagger.io/): Ferramenta para design, construção, documentação e uso de serviços web RESTful.
- [Prisma ORM](https://www.prisma.io/): ORM (Object-Relational Mapping) para comunicação com o banco de dados.
- [Docker](https://www.docker.com/): Ambiente de desenvolvimento para o PostgreSQL e Redis.
- [Autenticação JWT](https://jwt.io/): Autenticação baseada em JSON Web Tokens para garantir segurança nas chamadas da API.

## 🏗️ Padrões de Projeto

A aplicação segue os seguintes padrões de projeto:

1. **Clean Architecture**: A estrutura do projeto é organizada em camadas (entidades, use cases, interfaces) para separar as preocupações e facilitar a manutenção.

2. **Domain Driven Design (DDD)**: O design do software é orientado pelo domínio, com foco nas regras de negócio e nas entidades principais.

3. **Injeção de Dependência**: A inversão de controle e injeção de dependência são utilizadas para garantir a flexibilidade e testabilidade do código.

4. **Testes automatizado**: Os testes unitários, integração e end-2-end são escritos para garantir a qualidade do código.

## 🔄 Domain Events

A aplicação utiliza o conceito de **Domain Events** para atualizar a contagem de votos quando um novo voto é registrado ou alterado. Isso garante que a lógica de negócio relacionada à contagem de votos permaneça consistente.

## 🌐 GitHub

O código-fonte da aplicação pode ser encontrado no GitHub: [Link do Projeto](https://github.com/nitoba/poll-voting)

## 📧 Contato

Em caso de dúvidas ou sugestões, entre em contato através do e-mail: [nito.ba.dev@gmail.com](mailto:nito.ba.dev@gmail.com).
