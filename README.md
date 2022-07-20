# Hownd Examples

## Partner API
The Partner API is a secured REST API that allows you to manage Subscribers
on behalf of Hownd Customers, AKA Merchants and Businesses.

As a Partner you will be provided oAuth Client Credentials.
Your Credentials will be used to retrieve a Token that will be used as 
a Bearer Token to access the API. 

These Credentials must be secured and cannot be used in any integration 
scenario other than "server to server". You must not
use these Credentials in a browser or mobile application.

Please see [oAuth Client Credentials](https://oauth.net/2/grant-types/client-credentials/).

You will also receive a Tenant ID for each Hownd Customer you may access.
This will be used in the X-Tenant-ID header to identify the Hownd Customer.
Arguably the Tenant ID could be used in the HTTP path. Since it is security
related, a header felt more appropriate.

You must include three http headers:
1. Authorization: Bearer [token]
2. Content-Type: application/json
3. X-Tenant-Id: value provided by Hownd

### Subscribers
To create a subscriber, you must provide the following information:
1. Email Address

The Following values are optional:
1. First Name
1. Last Name
1. Phone Number
