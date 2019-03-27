# simple quiz
An app with few questions.
* Frontend - Vuejs
* Backend - Golang
* Database - Just in-memory , so no database

## Install dependencies
```
# Download and install gorilla mux router
go get github.com/gorilla/mux

# Download and install client NPM packages
cd web/
npm install
npm install vue-resource
npm install bootstrap-vue
```

# Run
### To run app Go server and Vue should be started
```
# Run Go server at port :8888
from root package run:
go run quiz.go

# Run Vue at port :8080
cd web/
npm run start
```

Open link http://localhost:8080/.
