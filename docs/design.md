# Requirements

This requirement is only for the MVP version of the app. The diagrams can be found in this [draw.io](https://drive.google.com/file/d/1mZI80SHS0B5LvmvcMb81fN2onL_vn0EX/view?usp=sharing).

## Functional

- User able to sign up to the app
- User able to login to the app
- User able to swipe(like/dislike) other users
  - User can only swipe on a profile once.
  - User can only swipe on 10 profiles total (pass + like) in 1 day (free-plan).
  - User can swipe on unlimited profiles in 1 day (premium-plan).
- User able to view other users
  - Can filter by gender, age, and location
  - Same profiles can’t appear twice.
- User able to purchase premium packages that unlocks one premium feature
  - Options:
    - No swipe quota for user
    - Verified label for user

## Data Entity

- User
  - Attributes:
    - `name`: string
    - `email`: string
    - `password`: string-hased
    - `gender`: enum
      - `male`
      - `female`
    - `birth_date`: date
    - `latitude`: float
    - `longitude`: float

- User Photo
  - Attributes:
    - `url`: string

- Matching
  - `user_one_id`: int
  - `user_two_id`: int
  - `status`: int
    - 0 -> `pending`
    - 1 -> `rejected`
    - 2 -> `matched`

- Subscription
  - `user_id`: int
  - `plan_id`: int
  - `start_date`: date
  - `end_date`: date

- Subscription Plan
  - `name`: string
  - `price`: float
  - `duration`: int
  - `limit`: int

## Data Relationship

- User has many User Photos
- User has many Matchings
- User has many Subscriptions

## Software Architecture

The software architecture will be a monolithic architecture. The reason for choosing a monolithic architecture is because it is easier to develop and deploy. The app is also small enough that a monolithic architecture is sufficient.

## Technical Stack

- **Programming Language:** Go  
  The reason for choosing Go is because it is a statically typed language that is easy to read and write. It is also a compiled language which means it is faster than interpreted languages like Python.
- **Database:** MongoDB
  In this usecase MongoDB is a good choice because it is a NoSQL database that is easy to scale horizontally. It is also a good choice because it is a document-based database which means it is easy to store and retrieve data.
