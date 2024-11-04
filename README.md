# stress-test

**Objetivo**: Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

## Entrada de Parâmetros via CLI:

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requests.
- `--concurrency`: Número de chamadas simultâneas.

## Execução do Teste:

1. Realizar requests HTTP para a URL especificada.
2. Distribuir os requests de acordo com o nível de concorrência definido.
3. Garantir que o número total de requests seja cumprido.

## Geração de Relatório:

Apresentar um relatório ao final dos testes contendo:
- Tempo total gasto na execução.
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Execução da aplicação:

Podemos utilizar essa aplicação fazendo uma chamada via Docker. Exemplo:
```bash
docker run <sua-imagem-docker> --url=http://google.com --requests=1000 --concurrency=10
```

# Lista de Checks para Implementação do Projeto de Teste de Carga

## 1. Setup do Projeto
- [X] Inicializar o módulo Go com `go mod init stress-test`.
- [X] Configurar o Cobra CLI com `cobra-cli init`.
- [ ] Criar comandos principais para receber parâmetros via CLI (`--url`, `--requests`, `--concurrency`).

## 2. Configuração dos Parâmetros CLI
- [ ] Definir e validar os parâmetros `--url`, `--requests`, e `--concurrency`.
- [ ] Garantir que `--url` seja uma URL válida.
- [ ] Verificar se `--requests` e `--concurrency` são números inteiros positivos.
- [ ] Adicionar mensagens de erro para parâmetros inválidos.

## 3. Execução do Teste de Carga
- [ ] Criar uma função para realizar requests HTTP à URL especificada.
- [ ] Implementar concorrência para distribuir requests com base no valor de `--concurrency`.
- [ ] Garantir que o número total de requests enviados corresponda ao valor de `--requests`.
- [ ] Gerenciar tempo limite e falhas nos requests (timeouts, falhas de conexão, etc.).

## 4. Coleta de Resultados
- [ ] Monitorar o tempo total de execução dos requests.
- [ ] Contabilizar o total de requests realizados.
- [ ] Contabilizar o total de requests com status HTTP 200.
- [ ] Registrar e classificar outros códigos de status HTTP (como 404, 500).

## 5. Geração de Relatório
- [ ] Criar uma função para gerar um relatório final após a execução dos testes.
- [ ] Incluir no relatório:
  - Tempo total de execução.
  - Quantidade total de requests realizados.
  - Quantidade de requests com status HTTP 200.
  - Distribuição de outros códigos de status HTTP.
- [ ] Apresentar o relatório no terminal de forma organizada.

## 6. Execução via Docker
- [ ] Criar um `Dockerfile` para o projeto.
- [ ] Configurar o `Dockerfile` para aceitar parâmetros CLI.
- [ ] Testar a execução da aplicação via Docker com um exemplo:
  ```bash
  docker run <sua-imagem-docker> --url=http://google.com --requests=1000 --concurrency=10
  ```

## 7. Testes e Validações
- [ ] Implementar testes para validar a execução dos requests e o processamento concorrente.
- [ ] Testar a coleta de dados e a geração de relatórios para garantir precisão.
- [ ] Validar a funcionalidade em um ambiente Docker para garantir compatibilidade.
