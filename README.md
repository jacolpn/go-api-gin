# GO-API-GIN

- **Subir o GO:** go run main.go
- **Primeira vez que sobe o docker:** docker-compose build
- **Subir o DOCKER:** docker-compose up
- **Subir os testes:** go test
- **Executar somente um teste:** go test -run 'nome da função do teste'
- **Criar uma nova database a partir do zero:** http://localhost:54321/browser/ > Server (Register > Server) > Geral|Name: alunos > Connection|Host name/addres: 172.19.0.2 (docker-compose exec postgres sh | hostname -i) > Connection|Maintenance database: root > Connection|Username: root > Connection|Password: root > Save
- 