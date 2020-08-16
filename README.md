# currencyapp

## Application development plan 
1. Create input screen reader function
2. Create function to download file 
3. Create function to read the currencies for file bind to struct(Model) and load in memory
4. Main function to combine the created functions to retrieve data read user input 


## Design considerations (Functionality)
1. Application loads the currency data once to memory when starting 
2. Application has inbuilt command to refresh the currency list from server when running
3. Application has a gracefull exit command
4. Currency data displayed shows the last time data was fetched from remote location
5. Bulk check of currency codes possible by user input of command separated codes

## Application structure 
1. main.go contains the application specific functionality
2. models.go contains data object definations and validations
3. utilities.go contains reusable fucntinality across other code bases 

## Application commands 
1. reload -> Application to fetch and update the list of currency codes while already running
2. exit -> application exiting gracefully 
3. Application accepts multi input currency code search as long the codes are comma separated


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

