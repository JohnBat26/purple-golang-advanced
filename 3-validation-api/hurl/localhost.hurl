# Login validation 1
POST http://localhost:8081/auth/login
{
    "email": "a@a.ru",
    "password": "1"
}
HTTP 201

# Login validation 2
POST http://localhost:8081/auth/login
{
    "email": "aa.ru",
    "password": "1"
}
HTTP 402
[Asserts]
body contains "failed"


# Login validation 3
POST http://localhost:8081/auth/login
{
    "email": "a@a.ru"
}
HTTP 402
[Asserts]
body contains "failed"


# Register validation 1
POST http://localhost:8081/auth/register
{
    "email": "a@a.ru",
    "password": "1",
    "name":"john"
}
HTTP 201

# Register validation 2
POST http://localhost:8081/auth/register
{
    "email": "aa.ru",
    "password": "1",
    "name":"john"
}
HTTP 402
[Asserts]
body contains "failed"


# Login validation 3
POST http://localhost:8081/auth/register
{
    "email": "a@a.ru",
    "password": "1"
}
HTTP 402
[Asserts]
body contains "failed"
