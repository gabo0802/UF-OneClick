# UF-OneClick

## Project Name
UF OneClick: Subscription Manager <br>

## Project Description
A subscription manager that allows users to track their subscriptions in one click. It has a login page so that every individual user's subscriptions can tracked by month. Actions that can be done include: adding a new subscription, calculating the total cost of subscriptions per month and by year, tracking when the subscription has been used, etc. <br>

## Project Features

### Subscription List
![image](https://user-images.githubusercontent.com/92041237/232912902-e5d79f0d-fafb-4926-a775-c85bde788a6e.png)
<br>
### Subscription Report
![image](https://user-images.githubusercontent.com/92041237/232912604-6c50400c-4d3c-4bbb-a5b5-2b7df7ddaaaf.png)
<br>
### Subscription Graph
![image](https://user-images.githubusercontent.com/92041237/232935983-aeb65c51-4890-4757-bd40-53af3735d64f.png)
<br>

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
7. Navigate to the client folder <code>cd Client</code> and run the command <code>npm install ng2-charts</code> and <code>npm install cypress</code>

## How to Run Project (Using Visual Studio)
1. Have the MySQLPassword.txt and EmailCode.txt files in same folder as makefile
2. Terminal -> New Terminal
3. Terminal -> Split Terminal
4. Run <code> make first </code> in first terminal
5. Run <code> make second </code> in second terminal
6. Go to: http://localhost:4200/ <br> (should open automatically)

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
* Password: <b>Found in MySQLPassword.txt</b>
