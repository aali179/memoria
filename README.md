# 🕰️ Memoria

A virtual multi-media scrapbook generator for treasured memories. Add images, songs, maps and more!

## 🤩 Team Members 

|    Member            |         Student Number               |
|--------------------- | ------------------------------------ |
|Ruwani De Alwis       |        300057076                     |
|Afrah Ali             |        300049798                     |

## 🖌️ Features
- Provides a collection of pre-designed scrapbook templates users can choose to use
  - Upload and resize images
  - Stylize text to add to the scrapbook
  - Attach songs (Spotify)
  - Pin locations (using Google Maps)
- Generates a collage page and saves to user's scrapbook
- Supports user account creation and management (DB)

## 🎨 Design System
View the documentation of our Design System [here.](https://github.com/professor-forward/memoria/blob/f/deliverable-2/designSystem/README.md) 

## 🛠 Tools
- **Client-side:** HTML, CSS/SASS
- **Server-side:** Gin (Go)
- **Database:** Postgres
- **Internal API:** REST
- **External APIs:** Google Maps, Spotify


## Installation & Deployment

- Deployment will be done using docker (a requirement to run the application)
- Navigate to the /app directory within the application
- Run `docker-compose build`, this will generate an executable of the Go application
- Run  `docker-compose up -d postgres`, this will start the postgres db in the background
- Run `docker-compose up app migrate`
   - this will start the go application at http://localhost:5000, along with a migration if needed

## Releases
- Memoria releases will be published as Docker images with version tags to our [Docker Hub registry.](https://hub.docker.com/layers/195686374/afrah412000/memoria/1.0/images/sha256-1fcf6955651cecf79e0d7a32fb34344d5f61e9a368f41183af4b8a81afe1c92e?context=repo&tab=layers)

## Migrations

- the postgres DB has been versioned and migration files have been created and run through `go-migrate`
- To run a db migration on its own `docker-compose up migrate`
   - this will read the migrations/ file in the folder and build a new version of the system (if changes have been present)
- As the migrations were not generated through GORM (typeORM used), a change to the model requires changes in both the appropriate files in the migrations/ folder and the structures in models/ folder

## CI/CD
- Our DevOps pipeline is hosted on GitHub Actions where we build and test our code before integration. For release versions, we tag, build and publish Docker images of Memoria to our registry.

## Testing
- Unit testing for CRUD and API handlers are done through our test suite and integrated into our CI pipeline. Tests can be run locally with `go test` from the `app/test` directory.