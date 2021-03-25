# Tech-assessment

This is a basic user management API implemented using Golang and Gorilla Framework.

 **Set Environment:**

* Download and install Go: https://golang.org/doc/install
* Download and install Goland: https://www.jetbrains.com/go/promo/?gclid=CjwKCAjw6fCCBhBNEiwAem5SO4fpcqL9GiPr-a1rwU1iSF4DOxom6rZxLccA9Y4xBZrMMjizLnnOzhoCLbEQAvD_BwE  
* Download and install MongoDB: https://docs.mongodb.com/manual/installation
* Clone the above repo: `git clone https://github.com/wiseman-ska/tech-assessment.git`

 **Run API:** 
 
* Open the project using Goland IDE
* Use any mongo client to create the following DB: `users-min-db`
* Open the terminal and run `go get` to download all project dependencies
* Execute the program using: `go build main.go`

 

**Register users**
----
Creates new user and returns json object containing user details
* **URL**

     _/users/create_

* **Method:**

    _`POST`_
  
*  **URL Params**

   **Required:** `  ""  `
 
   **Optional:** `  ""  `
    
* **Data Params**

  _Example_
  ```json
  {
      "data" : {
          "firstName": "Wiseman",
          "lastName": "Qamata",
          "email": "qamata@mail.com",
          "mobileNumber": "0617069570",
          "idNumber": "9512216126086",
          "physicalAddress": "Westlake Eco-Estate, 90 Westlake Avanue; Modderfontan; 1604",
          "password": "12345"
       }
   }
  ```

* **Success Response:**
  
  * **Code:** 200 <br />
    **Content:** 
    
    ```json
            {
                "data": {
                     "id": "",
                     "firstName": "Wiseman",
                     "lastName": "Qamata",
                     "email": "qamata@mail.com",
                     "mobileNumber": "0617069570",
                     "idNumber": "9512216126086",
                     "physicalAddress": "Westlake Eco-Estate, 90 Westlake Avanue, Modderfontan, 1604",
                     "password": "",
                     "hashPassword": null
                }
            }

     ```
 
* **Error Response:**

  * **Code:** 500 Internal Server Error <br />
    **Content:** 
    
    ```json
        {
            "data": {
                "error": "E11000 duplicate key error collection: users-min-db.users index: email_1 dup key: { email: \"wiseman@mail.com\" }",
                "message": "An unexpected error has occurred",
                "status": 500
            }
        }
    ```
**Login users**
----
Login existing user and returns json object containing user details plus auth token
* **URL**

     _/users/login_

* **Method:**

    _`POST`_
  
*  **URL Params**

   **Required:** `  ""  `
 
   **Optional:** `  ""  `
    
* **Data Params**

  _Example_
  ```json
        {
            "data" : {
                "email": "wiseman@mail.com",
                "password": "012345"
             }
         }
  ```

* **Success Response:**
  
  * **Code:** 200 <br />
    **Content:** 
    
    ```json
           {
            "data": {
                "user": {
                     "id": "60598cae51d0944937f4d58b",
                     "firstName": "Wiseman",
                     "lastName": "Qamata",
                     "email": "qamata@mail.com",
                     "mobileNumber": "0617069570",
                     "idNumber": "9512216126086",
                     "physicalAddress": "Westlake Eco-Estate, 90 Westlake Avanue, Modderfontan, 1604",
                     "password": "",
                     "hashPassword": null
                },
                "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySW5mbyI6eyJOYW1lIjoid2lzZW1hbkBtYWlsLmNvbSIsIlJvbGUiOiJtZW1iZXIifSwiZXhwIjoxNjE2Njc5NjgxLCJpc3MiOiJhZG1pbiJ9.Af9PiNtb7cPpJYDiSt2g8sXbg_j4gq1iy6HlH_5841CrLp3dWbXmbQ-foS_emthPnyyNWxiVgKzV3okzPmP8A-sCvceF0mPd4-oQ8tE1-hTa50Od4nmg6bat4WWBfqfmNnrDFzolJu0F5ADZDt2QpMIXtnA-wANmPL3vlcHjTVw"
            }
        }

     ```
 
* **Error Response:**

  * **Code:** 401 Unauthorized <br />
    **Content:** 
    
    ```json
          {
              "data": {
                  "error": "crypto/bcrypt: hashedPassword is not the hash of the given password",
                  "message": "Invalid login credentials",
                  "status": 401
              }
          }
    ```

**Retrieve users**
----
Retrieves existing users and returns json object containing their details
* **URL**

     _/api/v1/users/all_

* **Method:**

    _`GET`_
  
*  **URL Params**

   **Required:** `  ""  `
 
   **Optional:** `  ""  `
    
* **Data Params**
    
   ```json

   ```

* **Success Response:**
  
  * **Code:** 200 <br />
    **Content:** 
    
    ```json
          {
          "data": [
              {
                  "id": "60598cae51d0944937f4d58b",
                  "firstName": "Wiseman",
                  "lastName": "Qamata",
                  "email": "qamata@mail.com",
                  "mobileNumber": "0617069570",
                  "idNumber": "9512216126086",
                  "physicalAddress":"Westlake Eco-Estate, Westlake Avanue, Modderfontan, 1604",
                  "password": "",
                  "hashPassword": null
              },
              {
                  "id": "605c8c2d0c5167aa21487897",
                  "firstName": "Arthur",
                  "lastName": "Skuku",
                  "email": "arthur@mail.com",
                  "mobileNumber": "0607069540",
                  "idNumber": "0112245830083",
                  "physicalAddress": "Westlake Eco-Estate, Westlake Avanue, Modderfontan, 1604",
                  "password": "",
                  "hashPassword": null
              }
          ]
      }

     ```
 
* **Error Response:**

  * **Code:** 500 Internal Server Error <br />
    **Content:** 
    
    ```json
        {
            "data": {
                "error": "Some error",
                "message": "An unexpected error has occurred",
                "status": 500
            }
        }
    ```

**Notes:** Please use bearer token to access other end-points. I didn't cover user update, user delete and get user by id in this doc, but the implementation is provided.

