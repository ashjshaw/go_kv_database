# go_kv_database
Key value database in Go

## GET /[key]
Returns the value for the provided key, if the key does not exist, a 404 status code is returned.

## PUT /[key]
This sets the value for the provided key to the contents of the request body, if the key already exists it is updated in place

## DELETE /[key]
This deletes the value for the provided key. If the key does not exist, a 404 is returned. 

## GET /
This returns a list of all keys as a JSON array.