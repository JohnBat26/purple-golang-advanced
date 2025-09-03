# Login validation 1
POST http://localhost:8081/send
{
    "email": "johnbat26@gmail.com"
}
HTTP 201
[Asserts]
body contains "johnbat26@gmail.com"
