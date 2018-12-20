# Number Log Application
Application opens a socket and restricts input to at most 5 concurrent clients. Clients will connect to the Application and write any number of 9 digit numbers, and then close the connection. The Application must write a de-duplicated list of these numbers to a log file in no particular order.

### Prerequisites

- Go must be downloaded and configured https://golang.org/doc/install
- Make sure your gopath is configured properly https://github.com/golang/go/wiki/SettingGOPATH
- This project should be located along the following filepath: src/newrelic-interviews/Alison.  If not, move the project or rename the package input references in the import statement to mirror your local environment.

### Running the code

- Clone the code from this repository.
- Run from the /Alison folder using the command `go run main.go`


### Assumptions 

- I assumed that I would be the only one working on this code sample.  As such, I pushed to master.  If I had been working on a team, I would have created a branch for each feature to avoid conflicts.  

### To Do for future versions

- Add test coverage to address unpredicted issues, make the code more readable, and durable.
- Add functionality to log numbers and capture the report in two more places to increase data resiliency.
- Add monitoring and test with 2M numbers per 10-second reporting period to identify opportunities to optimize for maximum throughput.
