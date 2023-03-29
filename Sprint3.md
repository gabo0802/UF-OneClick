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

