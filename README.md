![Build](https://img.shields.io/badge/build-passing-brightgreen.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)

## How to run

First, clone the repository:

```bash
git clone https://github.com/AFK068/designing-software-hw2.git
```

```bash
cd designing-software-hw2
```

Then, use the provided `Makefile` to run the application:

```bash
make run
```

## URL Endpoints

#### Animals

- **GET /animals**  
    Retrieve a list of all animals in the zoo.

- **POST /animals**  
    Add a new animal to the zoo.  
    **Request Body (JSON):**  
    ```json
    {
            "enclosureId": "string (uuid)",
            "species": "string",
            "name": "string",
            "birthDate": "string (date-time)",
            "gender": "Male or Female",
            "favoriteFood": "string",
            "status": "Healthy or Sick"
    }
    ```  
    (The ID of the newly created animal will be returned in the response.)

- **GET /animals/{animalId}**  
    Retrieve details of a specific animal by its ID.  
    **Path Parameter:**  
    - `animalId`: `string (uuid)`

- **DELETE /animals/{animalId}**  
    Remove an animal from the zoo by its ID.  
    **Path Parameter:**  
    - `animalId`: `string (uuid)`

- **POST /animals/{animalId}/move**  
    Move an animal to a new enclosure.  
    **Path Parameter:**  
    - `animalId`: `string (uuid)`  
    **Request Body (JSON):**  
    ```json
    {
            "newEnclosureId": "string (uuid)"
    }
    ```

#### Enclosures

- **GET /enclosures**  
    Retrieve a list of all enclosures in the zoo.

- **POST /enclosures**  
    Add a new enclosure to the zoo.  
    **Request Body (JSON):**  
    ```json
    {
            "type": "string",
            "size": "integer",
            "maxCapacity": "integer"
    }
    ```

- **GET /enclosures/{enclosureId}**  
    Retrieve details of a specific enclosure by its ID.  
    **Path Parameter:**  
    - `enclosureId`: `string (uuid)`

- **DELETE /enclosures/{enclosureId}**  
    Remove an enclosure from the zoo by its ID.  
    **Path Parameter:**  
    - `enclosureId`: `string (uuid)`

#### Feeding Schedules

- **GET /feeding-schedules**  
    Retrieve a list of all feeding schedules.

- **POST /feeding-schedules**  
    Add a new feeding schedule.  
    **Request Body (JSON):**  
    ```json
    {
            "animalId": "string (uuid)",
            "feedingTime": "string (date-time)",
            "foodType": "string"
    }
    ```
    (The ID of the newly created animal will be returned in the response.)

- **GET /feeding-schedules/{scheduleId}**  
    Retrieve details of a specific feeding schedule by its ID.  
    **Path Parameter:**  
    - `scheduleId`: `string (uuid)`

- **DELETE /feeding-schedules/{scheduleId}**  
    Remove a feeding schedule by its ID.  
    **Path Parameter:**  
    - `scheduleId`: `string (uuid)`

#### Statistics

- **GET /statistics**  
    Retrieve zoo statistics, including total animals, enclosures, and feeding schedules.


## Domain-Driven Design (DDD) Concepts

- **Entities**:  
  - `Animal`, `Enclosure`, and `FeedingSchedule` represent core domain entities, encapsulating attributes and behaviors specific to the zoo's business logic.

- **Value Objects**:  
  - Immutable concepts like `AnimalStatus`, `Gender`, and `Food` are modeled as Value Objects to ensure consistency and encapsulation.

- **Domain Events**:  
  - Events such as `AnimalMovedEvent` and `FeedingTimeEvent` are used to represent significant domain actions and ensure decoupling between components.

- **Encapsulation of Business Rules**:  
  - Business rules are encapsulated within domain objects, ensuring that logic like feeding schedules or enclosure capacity is managed directly by the entities.

---

## Clean Architecture Principles

- **Layered Dependencies**:  
  - All layers depend only inward. The `Domain` layer is completely independent and does not rely on any other layer.

- **Interface-Based Dependencies**:  
  - All dependencies between layers are managed through interfaces, ensuring loose coupling and flexibility.

- **Business Logic Isolation**:  
  - Business logic is fully isolated in the `Domain` and `Application` layers.

- **Mapping Between Layers**:
  -  Mappers handle conversion between domain models and API models.