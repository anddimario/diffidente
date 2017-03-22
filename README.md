# Diffidente
Share config for developers and devops 

### Features
- no specific requirements (use file for storage)
- basic auth
- ci/cd ready
- multiprojects and multiusers

### Requirements
GoLang

### Installation
clone the repository    
`go build -o diffidente *.go`    
Create a first admin with: `cd scripts && go run add_admin.go admin test`    
`./diffidente`    

### Curl endpoint example
**ADMIN METHODS**
- add user: `curl -d "username=test&password=test" http://admin:test@localhost:3000/admins/users`
- user list: `curl http://admin:test@localhost:3000/admins/users`
- get user: `curl http://admin:test@localhost:3000/admins/users?username=test`
- delete user: `curl -XDELETE http://admin:test@localhost:3000/admins/users?username=test`
- add a policy (update if already exists): `curl -d "username=test&app=ecommerce&keys=HOST=localhost\nPORT=8080" http://admin:test@localhost:3000/admins/policies`
- policy list: `curl http://admin:test@localhost:3000/admins/policies`
- get policy: `curl http://admin:test@localhost:3000/admins/policies?policy=policy_test_ecommerce`
- delete policy: `curl -XDELETE http://admin:test@localhost:3000/admins/policies?policy=policy_test_ecommerce`

**USER_METHODS**
- policy list: `curl http://test:test@localhost:3000/users/policies`
- get policy: `curl http://test:test@localhost:3000/users/policies?policy=policy_test_ecommerce`