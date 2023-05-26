Link Shortener API
Introduction

This is a simple API that can be used as a backend for a link shortener website.

Authorization: Bearer <YOUR_API_KEY>

Endpoints


POST /create

This endpoint allows you to create a short link from a long URL.
Request
    Body:

    json

    {
      "linkid": "a-memorable-link-identifier",
      "link": "a-website.com",
      "createdat": "current-time-in-iso"
    }

Response

    Status: 200 OK
    Body:

    json

    {
     "status": "Created"
    }

Retrieve a Shortened URL

bash

GET /link/:id

This endpoint allows you to retrieve a redirect to a previously shortened URL. Intended for browsers.
Request

Response

    Status: 301 MOVED PERMANENTLY
    Body:
	"The link"

