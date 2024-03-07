Steps:
1. Make sure all dependencies are downloaded.
2. Current directory must be `blogging-platform`.
3. Start the server: `go run server/main.go`.
4. Start the client: `go run client/main.go`.
5. Go to POSTMAN and do the following:


Get all Blogs:
GET http://127.0.0.1:5001/blogs

Get a Blog with ID 1:
GET http://127.0.0.1:5001/blog/1

Create a new Blog:
POST http://127.0.0.1:5001/blog
Use the JSON Body below.

Update a Blog with ID 2:
PUT http://127.0.0.1:5001/blog/2
Use the JSON Body below.

Delete a Blog with ID 1:
DEL http://127.0.0.1:5001/blog/1
Use the JSON Body below.


Standard JSON Body:
{
    "Title": "Chamunda",
    "Content": "Pundamanman is a gun man",
    "PublicationDate": "2009-11-10 23:00:00 +0000 UTC m=+0.000000001",
    "Tags": [
        "Chamundamananam",
        "Pundamanman",
        "worstliked"
    ]
}