# 📊 Documentação da Test API de Votação em Enquetes

## 🚀 Introdução

Bem-vindo à documentação da Test API para votações em enquetes. Esta API foi desenvolvida utilizando as seguintes tecnologias:

- [Golang](https://golang.org/): Linguagem de programação principal.
- [PostgreSQL](https://www.postgresql.org/): Banco de dados relacional para armazenar dados persistentes.
- [Redis](https://redis.io/): Armazenamento de cache para otimizar consultas frequentes.
- [Gin](https://gin-gonic.com/): Framework web para construir APIs em Golang.
- [Swagger](https://swagger.io/): Ferramenta para design, construção, documentação e uso de serviços web RESTful.
- [Prisma ORM](https://www.prisma.io/): ORM (Object-Relational Mapping) para comunicação com o banco de dados.
- Autenticação JWT: Autenticação baseada em JSON Web Tokens para garantir segurança nas chamadas da API.

## 🏗️ Padrões de Projeto

A aplicação segue os seguintes padrões de projeto:

1. **Clean Architecture**: A estrutura do projeto é organizada em camadas (entidades, use cases, interfaces) para separar as preocupações e facilitar a manutenção.

2. **Domain Driven Design (DDD)**: O design do software é orientado pelo domínio, com foco nas regras de negócio e nas entidades principais.

3. **Injeção de Dependência**: A inversão de controle e injeção de dependência são utilizadas para garantir a flexibilidade e testabilidade do código.

## 🔄 Domain Events

A aplicação utiliza o conceito de **Domain Events** para atualizar a contagem de votos quando um novo voto é registrado ou alterado. Isso garante que a lógica de negócio relacionada à contagem de votos permaneça consistente.

## 🌐 GitHub

O código-fonte da aplicação pode ser encontrado no GitHub: [Link do Projeto](https://github.com/seu-usuario/nome-do-repositorio)

## 📧 Contato

Em caso de dúvidas ou sugestões, entre em contato através do e-mail: [seu.email@example.com](mailto:seu.email@example.com).
