## Work Completed

Front-End
* Added Active and Inactive filtering for subscription table
* Added deactivate and reactivate functionality for subscriptions
* Added paginator functionality to subscription list table
* Added sorting functionality to subscription list table
* Added autocompletion for add active subscription and improved error handling
* Added functionality for adding inactive subscription
* Added delete functionality for inactive subscriptions
* Added side navigation for subscriptions and report

Back-End
* Added UserSubID to MySQL table UserSubs
* Added DeleteUserSub by value UserSubID functionality
* Fixed Bugs in SendAllReminders, DeleteUserSub, and dateToString Function (Angular)
* Added Specific Timezones
* Average Price of Subscriptions and Average Age of Subscriptions Queries Added
* Code cleanup

## Front-End Unit Tests

**Jasmine Tests:**

Add-Subscription Component:
* all Subscriptions should be initially empty
* max Date should be todays date
* minDate should be today date minus 40 years
* form subscription name should initially be empty
* form subscription price should initially be empty
* form subscription date should initially be maxDate
* start date entered beyond maximimum limit throws maxDate error
* start date entered below minimum limit throws minDate error

Add-Inactive Subscription Component:
* all Subscriptions should be initially empty
* max Date should be todays date
* minDate should be today date minus 40 years
* form subscription name should initially be empty
* form subscription price should initially be empty
* form subscription dateadded should initially be null
* form subscription dateremoved should initially be null
* Dates Validator function checking if error thrown on end date form if start date after end date
* Dates Validator function checking if error thrown on start end if start after end date
* start date entered beyond maximimum limit throws maxDate error
* start date entered below minimum limit throws minDate error
* end date entered beyond maximimum limit throws maxDate error
* start date entered below minimum limit throws minDate error

**Cypress tests:**

Subscription-list Component:
* Updated - Add Active Subscription Button has text Add Active Subscription
* Add Inactive Subscription Button has text Add Inactive Subscription