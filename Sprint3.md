## Work Completed

Front-End
* Profile functionality implemented and improved
* Added UI for the Subscription Manager: Seamless (initial) implementation of adding/deleting subscriptions
* Added the UI for the Subscription Report, doesn't have functionality just yet (First thing we're doing for Sprint 4)
* Worked on unit tests


## Front-End Unit Tests

**Cypress tests:**

Login Component:
* Form should be initially empty
* Form values should be accurate
* Login Button disabled when form is empty
* Login Button disabled when password form is empty
* Login Button disabled when username form is empty
* Login Button enabled when username and password filled in

Sign Up Component:
* Form should be initially empty
* Form values should be accurate
* Sign Up Button disabled when form is empty
* Sign Up Button disabled when password form is empty
* Sign Up Button disabled when username form is empty
* Sign Up Button disabled when email form is empty
* Sign Up Button enabled when username, email, password forms filled in

Profile Component:
* Correct heading for Profile Component
* Correct delete profile button text
* Correct back button text

Username-Field Component:
* mounts Username-Field Component
* Label contains Username
* Edit button contains text Edit
* Edit button should change text to Save when clicked
* Valid Input into form enables button

Email-Field Component:
* mounts Email-Field Component
* Label contains Email:
* Edit button contains text Edit
* Edit button should change text to Save when clicked

Time Zone-Field Component:
* mounts Time Zone-Field Component
* Label contains Time Zone:
* Edit button contains text Edit
* Edit button should change text to Save when clicked

**Jasmine Tests:**

Profile Component:
* password input field hidden by default
* username empty before loading
* email empty before loading
* time zone empty before loading

Username-Field Component:
* editing should be false initially
* oldUsername should be empty before call
* Username form should be disabled initially
* editing variable set to true when editUsername() initially called
* when editUsername() called form is enabled
* when editUsername() called form value is empty string
* when editing if editUsername() called, editing is false and form is disabled
* Username form has duplicate error if same username is entered

Email-Field Component:
* editing should be false initially
* oldEmail should be empty before call
* Email form should be disabled initially
* editing variable set to true when editEmail() initially called
* when editEmail() called form is enabled
* when editEmail() called form value is empty string
* when editing if editEmail() called, editing is false and form is disabled
* Email form has duplicate error if same email is entered

Password-Field Component:
* password input field should be hidden by default
* form should be disabled
* form value should be asterisks

Time Zone-Field Component:
* editing should be false initially
* oldTimeZone should be empty before call
* Time Zone form should be disabled initially
* editing variable set to true when editTimeZone() initially called
* when editTimeZone() called form is enabled
* when editTimeZone() called form value is only "UTC"
* when editing if editTimeZone() called, editing is false and form is disabled
* Time Zone form has duplicate error if same Time Zone is entered

PasswordReset Component:
* oldPassword and newPassword input fields hidden by default
* OldPassword and NewPassword form values should be empty string intially
* if newPassword is idenitical to oldPassword throws form duplicate error

Users Component:
* username variable should initially be empty
* subscriptionList should be initially be empty

Welcome-header Component:
* username should be initially empty
* currentDate should be today

Subscription-List Component:
* subscription list should be initially empty

## Back-End Unit Tests

**MySQL Package Tests**
Each unit test checks the general functionality of their respective function:
* TestMySQLConnect(t *testing.T)
* TestGetDatabaseSize(t *testing.T)
* TestSetUpTables(t *testing.T)
* TestGetTableSize(t *testing.T)
* TestResetTable(t *testing.T)
* TestResetAllTables(t *testing.T)
* TestCreateNewUser(t *testing.T)
* TestCreateAdminUser(t *testing.T)
* TestCreateCommonSubscriptions(t *testing.T)
* TestGetPassword(t *testing.T)
* TestChangePassword(t *testing.T)
* TestGetEmail(t *testing.T)
* TestChangeEmail(t *testing.T)
* TestGetUsername(t *testing.T)
* TestChangeUsername(t *testing.T)
* TestCreateNewSub(t *testing.T)
* TestCreateNewUserSub(t *testing.T)
* TestAddOldUserSub(t *testing.T)
* TestCancelUserSub(t *testing.T)
* TestDeleteUser(t *testing.T)
* TestLogin(t *testing.T)
* TestGetMostUsedSubscription(t *testing.T)
* TestGetPriceForMonth(t *testing.T)

**Handler Package Tests**
The unit tests create mock gin contexts and cookies to test output:
* TestSetDB(t *testing.T)
* TestDailyReminder(t *testing.T)
* TestNewsLetter(t *testing.T)
* TestChangeTimezone(t *testing.T)
* TestTryLogin(t *testing.T)
* TestNewUser(t *testing.T)
* TestGetAllSubscriptionServices(t *testing.T)
* TestGetAllCurrentUserSubscriptions(t *testing.T)
* TestGetMostUsedUserSubscription(t *testing.T)
* TestDeleteUser(t *testing.T)

## Back-End API
<b>[Error Codes Can Be Viewed In UF-OneClick/ProjectNotes/errorcode.txt]</b>

### MySQL
<code> MySQLConnect() *sql.DB </code>
<br> Establishes a connection to the remote database hosted on a Microsoft Azure server. Returns the variable holding the database connection. <br>

<code> GetDatabaseSize(db *sql.DB) int </code>
<br> Returns the number of tables in the main database (called "userdb"). <br>

<code> SetUpTables(db *sql.DB) </code>
<br> Creates the "Users," "Subscriptions," "UserSubs," and "Verification" tables in the database. <br>

<code> GetTableSize(db *sql.DB, tableName string) int </code>
<br> Returns the size of the table of the specified tableName passed in. <br>
  
<code> ResetTable(db *sql.DB, tableName string) </code>
<br> Removes the specified tableName from the database. <br>
  
<code> ResetAllTables(db *sql.DB) </code>
<br> Removes all tables from the database. <br>
  
<code> CreateNewUser(db *sql.DB, username string, password string, email string) int </code>
<br> Inserts a new entry into the User table with an inputted username, password, and email. Cannot use: empty inputs, emails that already exist in the system, usernames that already exist in the system. Returns 1 if successful, and an error code otherwise. <br>
  
<code> CreateAdminUser(db *sql.DB) </code>
<br> Inserts an admin user into the User table for testing. <br>
  
<code> CreateTestUser(db *sql.DB) </code>
<br> Inserts an test user into the User table and adds some random subscriptions for them for testing purposes. <br>

<code> CreateCommonSubscriptions(db *sql.DB) </code>
<br> Inserts many common subscription entries into the Subscriptions table. <br>

<code> GetPassword(db *sql.DB, userID int) string </code>
<br> Returns the password of the specified userID. <br>

<code> ChangePassword(db *sql.DB, userID int, oldPassword string, newPassword string) int </code>
<br> Takes in a userID, old password, and new password as parameters. Changes the password based on the inputted userID. oldPassword and newPassword must not be empty. Returns 1 if successful, and an error code otherwise. <br>

<code> GetEmail(db *sql.DB, userID int) string </code>
<br> Returns the email of the specified userID. <br>
  
<code> ChangeEmail(db *sql.DB, userID int, newEmail string) int </code>
<br> Takes in a userID and new email. The userID specifies which users' email to change. newEmail must not already exist in the database. Returns 1 if successful, and an error code otherwise. <br>

<code> GetUsername(db *sql.DB, userID int) string </code>
<br> Returns the username of the specified userID or an error code otherwise. <br>
  
<code> ChangeUsername(db *sql.DB, userID int, newUsername string) int </code>
<br> Changes the username based on the userID. newUsername must not already exist in the database. Returns 1 if successful, and an error code otherwise. <br>
  
<code> CreateNewSub(db *sql.DB, name string, price string) int </code>
<br> Inserts a new subscription into the Subscriptions table. name must not already exist in database. Returns 1 if successful, and an error code otherwise. <br>
  
<code> CreateNewUserSub(db *sql.DB, userID int, subscriptionName string) int </code>
<br> Inserts a new userSub into the UserSubs table. subscriptionName must: not be empty, already exist in the Subscriptions table, and not currently associated with the userID. DateAdded column in the table is based on the current time. Returns 1 if successful and an error code otherwise. <br>
  
<code> AddOldUserSub(db *sql.DB, userID int, subscriptionName string, dateAdded string, dateCanceled string) int </code>
<br> Inserts a userSub into the UserSubs table, but with a specified dateAdded and dateCancelled entry. This is used for old subscriptions that want to be entered. Returns 1 if successful and an error code otherwise. <br>
  
<code> CancelUserSub(db *sql.DB, userID int, subscriptionName string) int </code>
<br> Updates a userSub's subscriptionName DateRemoved value based on the current time.  Returns 1 if successful, and an error code otherwise. <br>
  
<code> DeleteUser(db *sql.DB, ID int) </code>
<br> Removes a user entry from all associated tables based on the inputted userID. <br>
  
<code> Login(db *sql.DB, username string, password string) int </code>
<br> Checks to see if inputted username and password already exist in the database. Returns the current userID if successful, and an error code otherwise. <br>
  
<code> GetMostUsedSubscription(db *sql.DB, currentID int, isContinuous bool, isCurrentlyActive bool) (string, int) </code>
<br> Calculates the most used subscription based on the userID and if the subscription is continuous and/or currently active on the plan. Returns the subscription name and the time used with the subscription in seconds. <br>
  
<code> GetPriceForMonth(db *sql.DB, currentID int, monthNumber int, yearNumber int) string </code>
<br> Returns the price of an inputted month and year of a specified user's subscription total. <br>
  
<code> ManuallyTestBackend(db *sql.DB) </code>
<br> Allows manual testing of functions that manipulate the database with a simple user interface based on inputs of 1-10. <br>

### Handler
<code> SetDB(db *sql.DB) </code> 
<br> <b>Must be done after database connection established using MySQL.MySQLConnect().</b> Sets the global variable of currentDB to db.<br>

<code> DailyReminder(c *gin.Context) </code> 
<br> <b>Works with any HTTP request.</b> Sends an email to users once per day that contains all of the subsubscription charges that will happen in 1 day and 1 week with the email format being <i> Subscription Name, Subscription Price, Date to Be Renewed </i>. <br>

<code> NewsLetter(c *gin.Context) </code> 
<br> <b>Works only with a POST or PUT request with the message parameter being not null.</b> Sends an email to all verified users with the header of the email being "UF-OneClick Newsletter <i>Current Date</i>" and the message being the message parameter.<br>

<code> VerifyEmail(c *gin.Context) </code> 
<br> <b>Works only with a GET request with the /:code part of the URL being not null.</b> Checks if the code is a valid code and verifies the user the code matches to. Deletes any users that are unverified and those code has expired. <br>

<code> ChangeTimezone(c *gin.Context) </code>
<br> <b>Works only with a POST or PUT request with the timezonedifference parameter being not null.</b> Sets the global variable of currentTimezone to timezonedifference.<br>

<code> TwoFactorAuthentication() </code> 
<br> <b> Currently not functional </b> <br>

<code> TryLogin(c *gin.Context) </code> 
<br> <b>Works only with a POST or PUT request with the username and password parameters being not null.</b> Calls MySQL.Login() and if there is no error code returns a JSON object with a success message, updates the currentUserID to the returned userID, and creates a cookie with the userID that lasts 1 hour. Returns a JSON object with an error corresponding to the proper error code.<br>

<code> NewUser(c *gin.Context) </code>
<br> <b>Works only with a POST or PUT request with the username, email, and password parameters being not null.</b> Calls MySQL.CreateNewUser() and if there is no error code returns a JSON object with a success message and starts the email verification process. Returns a JSON object with an error corresponding to the proper error code.<br>

<code> GetAllUserSubscriptions() gin.HandlerFunc </code> 
<br> <b>Works with any HTTP request.</b> Returns all of the values in the table UserSubs joined with the table Subscriptions where the UserID is equal to the currentUserID. The values are formatted in the form <i> Subscription Name, Subscription Price, Date Added, Date Removed </i> and are ordered by the DateAdded. <br>

<code> GetMostUsedUserSubscription(isContinuous bool, isActive bool) gin.HandlerFunc </code> 
<br> <b>Works with any HTTP request.</b> If isContinuous is false returns the most used user subscription based on currentUserID. If isContinuous is true returns the most used <b>continuous</b> user subscription based on currentUserID. If isActive is true returns the most used <b>currently active</b> user subscription based on currentUserID.Format of return is <i> Subscription Name, Time Used (in seconds)</i>. <br>

<code> Logout() </code> 
<br> <b> Currently not functional </b> <br>

<code> DeleteUser(c *gin.Context) </code> 
<br> <b>Works with any HTTP request.</b> Calls MySQL.DeleteUser(). Deletes the user whose userID is equal to currentUserID. Also deletes the cookie containing currentuserID. <br>

<code> NewUserSubscription(c *gin.Context) </code> 
<br> <b>Works only with a POST or PUT request with the name parameter not null.</b> Calls MySQL.CreateNewUserSub(). Attempts to create a new user subscription to the currentUserID to the subscription whose name is equal to the paramter name, dateadded is equal to the current data and time, and dateremoved is null. Returns JSON object with success message if successful and returns proper error message corresponding to the error code if unsuccessful. <br>

<code> NewPreviousUserSubscription(c *gin.Context) </code> 
<br> <b>Works only with a POST or PUT request with the name and dateadded parameters being not null. The dateremoved parameter does not have to be not null.</b> Calls MySQL.AddOldUserSub(). Attempts to create a new user subscription to the currentUserID to the subscription whose name is equal to the paramter name, dateadded is equal to the dateadded paramter, and dateremoved is null or equal to the dateremoved parameter. Returns JSON object with success message if successful and returns proper error message corresponding to the error code if unsuccessful. <br>

<code> NewSubscriptionService(c *gin.Context) </code> 
<br> <b>Works only with a POST or PUT request with the name and price parameters being not null.</b> Calls MySQL.CreateNewSub(). Attempts to create a new subscription service whose name is equal to name and price is equal to price. Returns JSON object with success message if successful and returns proper error message corresponding to the error code if unsuccessful. <br>

<code> CancelSubscriptionService(c *gin.Context) </code> 
<br> <b>Works only with a POST or PUT request with the name parameter not null.</b> Calls MySQL.CreateNewUserSub(). Attempts to update dateremoved to the current date and time in an existing user subscription of the currentUserID and subscription whose name is equal to the paramter name. Returns JSON object with success message if successful and returns proper error message corresponding to the error code if unsuccessful. <br>

<code> ChangeUserPassword(c *gin.Context) </code> 
<br> <b>Works only with a POST or PUT request with the oldPassword and newPassword parameters being not null.</b> Calls MySQL.ChangePassword(). Attempts to update the old password of the currentUserID to the newpassword. Returns JSON object with success message if successful and returns proper error message corresponding to the error code if unsuccessful.<br>

<code> ChangeUserUsername(c *gin.Context) </code> 
<br> <b>Works only with a POST or PUT request with the email parameter not null.</b> Calls MySQL.ChangeEmail(). Attempts to update the username of the currentUserID to the new username. Returns JSON object with success message if successful and returns proper error message corresponding to the error code if unsuccessful.<br>

<code> ChangeUserEmail(c *gin.Context) </code> 
<br> <b>Works only with a POST or PUT request with the email parameter not null.</b> Calls MySQL.ChangeEmail(). Attempts to update the email of the currentUserID to the new email. Returns JSON object with success message if successful and returns proper error message corresponding to the error code if unsuccessful.<br>

<code> GetUserInfo(c *gin.Context) </code> 
<br> <b>Works with any HTTP request.</b> Calls MySQL.GetUsername() and MySQL.GetEmail(). Returns JSON object in the format <i>username, email</i>. <br>

<code> GetTimezone(c *gin.Context) </code> 
<br> <b>Works with any HTTP request.</b> Converts currentTimezone global variable to a string. Returns a JSON object in the format "CurrentTimezone:", <i>currentTime</i>. <br>

<code> ResetALL(c *gin.Context) </code> 
<br> <b>Works with any HTTP request. currentUserID must be equal to 1.</b> Calls MySQL.ResetAllTables(), MySQL.SetUpTables(), MySQL.CreateAdminUser(), and MySQL.CreateCommonSubscriptions(). Removes all cookies. Redirects to http:/localhost:4200/login<br>

<code> GetAllUserData(c *gin.Context) </code> 
<br> <b>Works with any HTTP request. currentUserID must be equal to 1.</b> Returns a JSON object with all of the user data, subscriptions data, and usersubs data from the database.<br>
<br>

## Video
