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


## Migrations

- the postgres DB has been versioned and migration files have been created and run through `go-migrate`
- To run a db migration on its own `docker-compose up migrate`
   - this will read the migrations/ file in the folder and build a new version of the system (if changes have been present)
- As the migrations were not generated through GORM (typeORM used), a change to the model requires changes in both the appropriate files in the migrations/ folder and the structures in models/ folder