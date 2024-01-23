# Bookstore API with GO
***
## Explore with Postman: [Postman collection link](https://github.com/jahangir1x/book-crud/blob/main/postman_collection/book-crud.postman_collection.json)
***
## Endpoints
*JSON standard does not allow comments. The comments are used here for clear demonstration purposes.*
#### 1. Create Author
- **Path:** `/bookstore/authors`
- **Method:** `POST`
- **Authorization required:** `YES`
- **Request Body:**
  ```json
  {
    "authorName":  "John Doe",       // required
    "address":     "123 Main St.",   // optional
    "phoneNumber": "+12 34 5678 910" // optional
  }
  ```
- **Example:**
    - `POST /bookstore/authors`
      ```json
      {
        "authorName":  "Example Author",
        "address":     "Example Address",
        "phoneNumber": "+1234567890"
      }
      ```
    - `POST /bookstore/authors`
      ```json
      {
        "authorName": "Example Author"
      }
      ```
***
#### 2. Get Authors

- **Path:** `/bookstore/authors`
- **Method:** `GET`
- **Authorization required:** `NO`
- **Example:**
    - `GET /bookstore/authors`
***
#### 3. Get Author by ID
- **Path:** `/bookstore/authors/:id`
- **Method:** `GET`
- **Example:**
    - `GET /bookstore/authors/1`
    - `GET /bookstore/authors/123`
***
#### 4. Delete Author by ID
- **Path:** `/bookstore/authors/:id`
- **Method:** `DELETE`
- **Example:**
    - `DELETE /bookstore/authors/1`
    - `DELETE /bookstore/authors/123`
***
#### 5. Update Author by ID
- **Path:** `/bookstore/authors/:id`
- **Method:** `PUT`
- **Request Body:**
  ```json
  {
    "authorName":  "Jane Smith",     // optional
    "address":     "987 Main St.",   // optional
    "phoneNumber": "+10 98 7654 321" // optional
  }
  ```
- **Example:**
    - `PUT /bookstore/authors/1`
      ```json
      {
        "authorName": "New Name",
        "address":    "New Address"
      }
      ```
    - `PUT /bookstore/authors/123`
      ```json
      {
        "authorName": "New Name"
      }
      ```
***
#### 6. Create Book
- **Path:** `/bookstore/books`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "bookName":    "Book Name", // required
    "authorID":    99,          // required
    "publication": "Pub Name"   // optional
  }
  ```
- **Example:**
    - `POST /bookstore/books`
      ```json
      {
        "bookName":    "Example Book",
        "authorID":    42,
        "publication": "Example Publication"
      }
      ```
    - `POST /bookstore/books`
      ```json
      {
        "bookName": "Example Book",
        "authorID": 42
      }
      ```
***
#### 7. Get Books

- **Path:** `/bookstore/books`
- **Method:** `GET`
- **Query Parameters:**
    - `bookName`: *(optional)* Filter books by name.
    - `id`: *(optional)* Filter books by ID.
    - `authorID`: *(optional)* Filter books by author ID.
    - `publication`: *(optional)* Filter books by publication.
- **Example:**
    - `GET /bookstore/books`
    - `GET /bookstore/books?bookName=ExampleBook&authorID=42`
    - `GET /bookstore/books?authorID=42`
***
#### 8. Get Book by ID

- **Path:** `/bookstore/books/:id`
- **Method:** `GET`
- **Example:**
    - `GET /bookstore/books/1`
    - `GET /bookstore/books/123`
***
#### 9. Delete Book by ID

- **Path:** `/bookstore/books/:id`
- **Method:** `DELETE`
- **Example:**
    - `DELETE /bookstore/books/1`
    - `DELETE /bookstore/books/123`
***
#### 10. Update Book by ID

- **Path:** `/bookstore/books/:id`
- **Method:** `PUT`
- **Request Body:** *(all fields are optional)*
  ```json
  {
    "bookName":    "New Name", // optional
    "authorID":    99,         // optional
    "publication": "New Pub"   // optional
  }
  ```
- **Example:**
    - `PUT /bookstore/books/1`
      ```json
      {
        "bookName":    "New Name",
        "publication": "New Pub"
      }
      ```
    - `PUT /bookstore/books/123`
      ```json
      {
        "bookName": "New Name"
      }
      ```
***

### Note

- Ensure to replace `:id` in the paths with the actual ID of the book when making requests.