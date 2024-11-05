# Stress Test

O **Stress Test** é uma aplicação desenvolvida em Go para realizar testes de carga em serviços web. Esta aplicação permite que você especifique a URL do serviço, o número total de requisições a serem feitas e o número de requisições simultâneas.

## Como Rodar

### Testes

Para rodar os testes da aplicação, use o seguinte comando:

```bash
docker compose up tests
```

### Aplicação

Para rodar a aplicação e realizar um teste de carga, utilize o comando abaixo, substituindo a URL, o número de requisições e a concorrência conforme necessário:

```bash
docker compose run app --url=http://google.com --requests=1000 --concurrency=100
```
