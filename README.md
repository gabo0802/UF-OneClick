# UF-OneClick

## Project Name
UF OneClick: Subscription Manager <br>

## Project Description
A subscription manager that allows users to track their subscriptions in one click. It has a login page so that every individual user's subscriptions can tracked by month. Actions that can be done include: adding a new subscription, calculating the total cost of subscriptions per month and by year, tracking when the subscription has been used, etc. <br>

## Project Members
### Front-End
Gabriel Castejon (gabo0802) <br>
Matthew Denslinger (mslinger) <br><br>

### Back-End
Vladimir Alekseev (valekseev03) <br>
Mason Enojo (enojom) <br><br>

## Project Setup
1. Install the go programming language https://go.dev/dl/
2. Install Node.js https://nodejs.org/
3. Install Angular via the command line <code>npm install -g @angular/cli</code>
4. Install MySQL (might be optional, see <i> How to Install MySQL </i> for more info) <br>
5. Clone respository <code>git clone https\://github.com/gabo0802/UF-OneClick.git</code> or via SSH
6. Run <code> go get github.com/gin-gonic/gin </code>
7. Navigate to the client folder <code>cd Client</code> and run the command <code>npm install</code> (Libraries to install: ng2-charts, cypress) 

## How to Run Project (Using Visual Studio)
1. Terminal -> New Terminal
2. Terminal -> Split Terminal
3. Run <code> make first </code> in first terminal
4. Run <code> make second </code> in second terminal
5. Go to: http://localhost:4200/ <br>

## How to Install MySQL
### How to get MySQLStuff.go to work with Go and Visual Studios (Windows 10):
1. Install Go: https://go.dev/dl/
2. Install MySQL: https://dev.mysql.com/downloads/installer/ 
3. Run command in Command Prompt terminal:  <code> go env -w GO111MODULE="off" </code>
4. Run command in Visual Studio Code terminal: <code> go get github.com/go-sql-driver/mysql </code> <br>

### Access MySQL Database:
* (For Windows): <code> mysql.exe -h oneclickserver.mysql.database.azure.com -u adminUser -p </code>
* (For Mac): <code> /usr/local/mysql/bin/mysql -h oneclickserver.mysql.database.azure.com -u adminUser -p </code>
* Hostname: oneclickserver.mysql.database.azure.com
* Username: adminUser
* Password: MySQLP@ssw0rd
