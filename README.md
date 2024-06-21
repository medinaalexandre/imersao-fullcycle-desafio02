# Desafio02
Neste desafio, você deverá criar uma aplicação Golang que gerará uma API REST.


No repositório da imersão existe um arquivo na pasta nextjs-front-end/node-api/data.js. Você irá pegar este arquivo e convertê-lo para JSON e salvar na raiz do seu projeto Golang como “data.json”.


Este arquivo será o banco de dados do seu projeto, a aplicação irá expor algumas rotas e sua aplicação deverá ler o arquivo JSON ao iniciar o programa com o go run main.go e manipular os dados sempre na memória (não salve nada no arquivo JSON).


Sua aplicação deverá expor os endpoints:

- [x] GET /events - Listar todos os eventos
- [x] GET /events/:eventId - Listar os dados de um evento
- [x] GET /events/:eventId/spots - Listar os lugares de um evento
- [x] POST /event/:eventId/reserve - Reservar um lugar
  - Dados:
spots: array com spots, exemplo [A1, B2];
Não poderá reservar o mesmo spot duas vezes, deverá retornar um erro 400

Compartilhe a aplicação num repositório Git e compartilhe o link!

Bons estudos!
