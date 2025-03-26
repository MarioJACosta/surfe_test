# Backend Coding Challenge

This repository contains a solution for a back-end coding challenge using **Go**. It implements a simple API server with in-memory data from two JSON files: `users.json` and `actions.json`.

## Technologies
- Language: Go
- Router: Gorilla Mux

---

## Design Decisions

- **Modular Structure**: Separated concerns into `models`, `handlers`, `router`, and `utils`
- **In-Memory Loading**: Loads JSON into memory at startup, no DB needed
- **Fast Lookup**: Uses maps for O(1) user and action access
- **DFS for Referrals**: Computes referral index using depth-first search to find indirect invites

---

## Features Implemented

### 1. Get User by ID
**GET /users/{id}**

Returns a user object with their ID, name, and created date.

### 2. Get User Action Count
**GET /users/{id}/actions/count**

Returns the total number of actions a user has performed.

### 3. Next Action Probabilities
**GET /actions/next/{type}**

Returns a probability breakdown of what actions users typically take immediately after performing the specified action.

### 4. Referral Index
**GET /users/referral-index**

Returns a mapping of each user ID to their Referral Index, defined as the number of unique users they have referred — either directly or indirectly.

A `REFER_USER` action indicates that a user has invited another user to the platform. 

If **User A** invites **User B**, and **User B** invites **User C**, then **A**'s referral index is **2** (B and C).
A user can only be invited once, so there are no referral cycles.


### Implementation
We build a referral graph: **map[int][]int** where keys are referrers and values are the list of referred users.

We then perform a **Depth-First Search (DFS)** from each referrer to discover all users reachable from them.

The size of this reachable set is their referral index.

**Time and Space Complexity**
|Step	|Time Complexity	|Space Complexity|
|-|-|-|
|Build referral graph|	O(n)	|O(n)|
|DFS per user (worst case)|	O(m * n)	|O(n)|

- n = number of `REFER_USER` actions
- m = number of users who referred someone

In the worst case, a user might indirectly refer every other user (long referral chains)

In practice, the graph is sparse, and referral chains are shallow, so the real performance is much faster.

---

## Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/MarioJACosta/surfe_test.git
cd surfe_test
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Place data files
Download and place `users.json` and `actions.json` into a `data/` directory at the root.

### 4. Run the server
```bash
go run main.go
```

Server will start on:
```
http://localhost:8080
```

---

## Sample Requests

```bash
curl http://localhost:8080/users/1
curl http://localhost:8080/users/1/actions/count
curl http://localhost:8080/actions/next/REFER_USER
curl http://localhost:8080/users/referral-index
```

---