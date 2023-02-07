# UF-OneClick

## Project Name
UF OneClick: Subscription Manager

## Project Description
A subscription manager that allows users to track their subscriptions in one click. It has a login page so that every individual user's subscriptions can tracked by month. Actions that can be done include: adding a new subscription, calculating the total cost of subscriptions per month and by year, tracking when the subscription has been used, etc.

## Project Members
### Front-End
Gabriel Castejon (gabo0802) <br>
Matthew Denslinger (mslinger) <br><br>

### Back-End
Vladimir Alekseev (valekseev03) <br>
Mason Enojo (enojom) <br>

## User Stories
* As a site member, I want a login page so I can access the website.
* As a site member, I want a way to track the cost of each subscription I have so that I can see which subscriptions I spend the most on in a year.
* As a site member, I want a profile page so that I can update or view my information.
* As a site visitor, I want a sign up page so that I can become a site member

## Project Setup
1) Install the go programming language https://go.dev/dl/
2) Install Node.js https://nodejs.org/
3) Install Angular via the command line <code>npm install -g @angular/cli</code>
4) Install MySQL (might be optional, see <i> How to Install MySQL </i> for more info)

## How to Run Project (Using Visual Studio)
1) Clone repository
2) Terminal -> New Terminal
3) Terminal -> Split Terminal
4) Run <code> cd Client </code> then <code> npm start </code> in first terminal
5) Run <code> make dev </code> in second terminal

## How to Install MySQL
How to get MySQLStuff.go to work with Go and Visual Studios (Windows 10):
1) Install Go: https://go.dev/dl/
2) Install MySQL: https://dev.mysql.com/downloads/installer/ 
3) Run command in Command Prompt terminal:  <code> go env -w GO111MODULE="off" </code>
4) Run command in Visual Studio Code terminal: <code> go get github.com/go-sql-driver/mysql </code>

Access MySQL Database:
* (For Windows): <code> mysql.exe -h oneclickserver.mysql.database.azure.com -u adminUser -p </code>
* (For Mac): <code> /usr/local/mysql/bin/mysql -h oneclickserver.mysql.database.azure.com -u adminUser -p </code>
* Hostname: oneclickserver.mysql.database.azure.com
* Username: adminUser
* Password: MySQLP@ssw0rd

## Go-Angular Tutorial/Test Code (Remove in Future)

* https://github.com/gabo0802/Go-AngularTest 
* https://www.globallogic.com/insights/blogs/develop-restful-api-with-go-and-gin/
