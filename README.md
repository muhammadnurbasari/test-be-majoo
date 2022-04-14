# test-be-majoo

This application is for the selection test at Majoo Indonesia


required:
1. golang version 1.17
2. mySQL

steps:
1. fill in sample.env in the config folder according to your environment
2. run the command "go mod vendor"
3. run the command "go run database/main.go --up"
4. run the command "go run main.go" to run the application
5. to see the documentation at {base_url}/docs/swagger/index.html
6. To login using basic authentication, refer to the client_id and client_secret in the sample.env  file