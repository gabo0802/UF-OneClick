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



## Helpful Info for SQL (Remove in Future)
How to get MySQLStuff.go to work with Go and Visual Studios (Windows 10):
1) Install Go: https://go.dev/dl/
2) Install MySQL: https://dev.mysql.com/downloads/installer/ 
3) Run command in Command Prompt terminal:  <b> go env -w GO111MODULE="off" </b>
4) Run command in Visual Studio Code terminal:  <b> go get github.com/go-sql-driver/mysql </b>
5) Run Server using MySQL Workbench (one of apps installed from Step 1)
5) Part 2: Password when setting up should be "MySQLP@ssw0rd"
6) Run command in MySQL terminal (possibly optional): <b> \sql  </b>
7) Run command in MySQL terminal (possibly optional): <b> \connect root@localhost  </b>
8) Run command in MySQL terminal (optional): <b> CREATE DATABASE user;  </b>
9) Run command in MySQL terminal (possibly optional): <b> USE user; </b>
10) Run MySQLStuff.go in Visual Studio terminal (like a normal Go program)

## Go-Angular Tutorial/Test Code (Remove in Future)

* https://github.com/gabo0802/Go-AngularTest 
* https://www.globallogic.com/insights/blogs/develop-restful-api-with-go-and-gin/
