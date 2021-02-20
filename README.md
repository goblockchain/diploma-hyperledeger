Ubuntu 16.04

Pré requisitos
Instalar os seguintes programas:
- Docker
- Docker-compose
- npm
- node
- curl
- Go Programming Language
- Hyperldger Fabric - curl -sSL http://bit.ly/2ysbOFE | bash -s

To start the network, run ./start-diploma.sh

Run cliente:

cliente/npm install && npm start


***************************
Create Diploma - JSON Model
***************************
{
	"collection": "collectionDiplomaDetailsOrg1",
    "universidade": {
		"universityId": "3",
		"universityName": "FIAP"
	},
	"diploma": {
		"diplomaId": "123456789",
		"diplomaHash": "0x1234567890",
		"curso": {
			"courseName": "Sistema de informação",
			"beginDate": "01/01/2015",
			"endDate": "12/12/2019",
			"professores": [{
					"teacherName": "João",
					"subject": "Matemática"
				}
			]
		},
		"aluno": {
			"studentId": "1234",
			"studentEmail": "fale.henrique@gmail.com",
			"studentName": "Maria Eduarda",
			"studentCpf": "23456789"
		}
	}
}

*************
Query Diploma
Type: Post
Format: JSON
*************

{
    "key": "3",
    "collection": "collectionDiplomaDetailsOrg2"
}

*********
Endpoints
*********
Create diploma for Org1

http://localhost:3000/diploma-create/

Query Diploma for Org1

http://localhost:3000/diploma-get/ 

Create diploma for Org2

http://localhost:3000/diploma-create-org2/ 

Query Diploma for Org2

http://localhost:3000/diploma-get-org2/ 

Query Diploma for Obs1

http://localhost:3000/diploma-get-obs1/ 


*******************************************************
********************Collections************************
*******************************************************

collectionDiploma:

	Essa coleção de dados são acessíveis pelas 3 organizações.
	
	Ela guarda as seguintes informações públicas sobre o diploma:
	
		UniversityId
		UniversityName
		DiplomaId
		StudentName
		CourseName
		EndDate
		StudentCpf
	
collectionDiplomaDetailsOrg1:

	Essa coleção guarda os diplomas da universidade 1, ou seja, org1.
	Todas as informações do JSON são armazenadas aqui e acessíveis
	pela org1 e obs 1

collectionDiplomaDetailsOrg2

	Essa coleção guarda os diplomas da universidade 2, ou seja, org2.
	Todas as informações do JSON são armazenadas aqui e acessíveis
	pela org2 e obs 1



Telas Front End

	Criar Diploma
	
		Descrição:

			Tela para a criação de diplomas da universidade executando a operação

		Campos para preencher e adicionar no JSON: 

			- Nome Aluno
			- Nome Curso
			- Data Início
			- Data Término
			- Nome professores
			- Aula de tais professores
			- ID Aluno
			- CPF Aluno
			- Email Aluno

		Campos inseridos automaticamente para criação do JSON

			- ID universidade
			- Nome Universidade
			- ID diploma (gerado automaticamente eu acho)
			- Hash do diploma (gerado automaticamente)
			- Coleção do diploma (ex. collectionDiplomaDetailsFiap)

	Consultar Diploma

		Descrição:

			Tela onde o aluno ou universidade vai poder consultar o diploma
		
		Campos (Se for universidade):

			- ID do Aluno 
			- CollectionUniversidade (Campo oculto, preencher no json)

		Campos (Se for o aluno):

			- ID Aluno
			- Senha
			- CollectionUniversidade (Campo oculto, preencher no json)

	Detalhes Diploma

		Descrição:

			Página que tratará do JSON recebido pela query e mostrar na tela
			com os detalhes do diploma
		
		Campos:

			- Nome Aluno
			- Nome Curso
			- Data Início
			- Data Término
			- Nome professores
			- Aula de tais professores
			- ID Aluno
			- CPF Aluno
			- Email Aluno
			- Nome Universidade
			- ID Diploma
			- Hash Diploma