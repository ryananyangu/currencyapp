# currencyapp

## Running application
1. Test environment with go compiler  
a) navigate to root directory where main.go file exists.  
b) run command below either command
```bash
go run .
```
or 
```bash
go run main.go
```

2. Production environment executable file (Linux env)  
a) While in the project root directory  
b) Run command below  
```bash
go build -o bin/currencyapp
```
A binary executable file will be created in directory path bin/  
To execute the binary file in the created bin directory
```bash
./currencyapp
```

## Application commands 
1. reload -> Application to fetch and update the list of currency codes while already running
2. exit -> application exiting gracefully 
3. help -> displays instructions for using the application
4. Application accepts multi input currency code search as long the codes are comma separated

## Application structure 
1. main.go contains the application specific functionality
2. models.go contains data object definations and validations
3. utilities.go contains reusable fucntinality across other code bases 