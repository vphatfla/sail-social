# Bacend API Document

host: `localhost:3000`
prefix for all api: `/api`
    Example: `localhost:3000/api/helloworld`

Creator APIs:
    Public APIs:
        | API Name | endPoint | Request URL Param | Request JSON PAYLoad | Response Example |
        |----------|----------|-------------------|----------------------|------------------|
        | Sign Up | `/api/creator/sign-up` | N/A | ``` {
    "email": "example4@some.com",
    "password": "abcd12345",
    "phoneNumber": "1213456111"
} ``` | `{"error":"Email Exists"}` or `abc`

