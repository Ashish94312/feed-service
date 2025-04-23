

### Steps to Set Up
1. **Clone the Repository**  
   ```bash
   git clone feed-service
   cd feed-service
   ```

2. **Docker Compose Setup**
   ```bash
   docker-compose build
   ```

---

## 🏃‍♂️ How to Run the Service

Run the entire service stack using Docker Compose:
```bash
docker-compose up
```

This will start two servers:
- **GraphQL Server** on `http://localhost:8080`
- **gRPC Server** on port `:50051` internally (used by GraphQL resolver).

---

## 📡 Sample GraphQL Query

Here's a simple GraphQL query to fetch a user's timeline:

```graphql
query GetUserTimeline {
  getTimeline(userID: "u1") {
    id
    authorID
    content
    timestamp
  }
}
```

---

## 📚 Description of Approach

### Architecture Overview
This service uses a **GraphQL** API to provide a structured feed, fetching posts from followed users through a simulated **gRPC** backend.

### Key Components:
- **GraphQL API**: Defined using `gqlgen`, providing a query `getTimeline`.
- **gRPC Server**: Serves mock post data (for simplicity), simulating an external posts service.
- **Concurrency**: Posts are fetched concurrently from each followed user, improving response times.

### Data Model:
- Users, followers, and posts are mocked within an in-memory dataset (`data.go`) for demonstration.
- Posts contain an ID, author ID, content, and a timestamp.

### Resolver Logic:
- The resolver (`GetTimeline`) concurrently retrieves posts from all users followed by the requested user.
- Posts from different users are aggregated and sorted by timestamp in descending order.
- The response is limited to the latest 20 posts.

---

