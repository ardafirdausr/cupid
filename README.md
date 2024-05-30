# Cupid

Cupid is Tinder-like backend service for matchmaking. It is a RESTful API that allows users to create accounts, view other users, and match with other users.

## Software Architecture

This project uses monolithic architecture for the mvp version. The reason for choosing a monolithic architecture is because it is easier to develop and deploy. The app is also small enough that a monolithic architecture is sufficient.  

For the software architecture, we will be using clean architecture. Clean architecture is a software design philosophy that separates the software into layers. Each layer has a specific responsibility and interacts with other layers in a specific way. The layers are:

- Entities
- Use Cases
- Interface Adapters
- Frameworks and Drivers

For more information on clean architecture, see [cleanarchitecture.io](https://www.cleancoders.com/episode/clean-code-episode-42/show).

## Git Branching Strategy

This repository uses trunk-based development branching strategy. The main branch is `main`. All feature branches are branched off from `main` and merged back into `main`.

For more information on trunk-based development, see [trunkbaseddevelopment.com](https://trunkbaseddevelopment.com/).
