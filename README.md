# 🕰️ Memoria

A virtual multi-media scrapbook generator for treasured memories. Add images, songs, maps and more!

## Contributers 
Afrah Ali, Ruwani De Alwis 

## 🖌️ Features
- Supports user account creation and management
- Allows users to create new scrapbooks and create virtual collages to add to them
- Users can:
  - Upload images
  - Add image captions and other text
  - Attach songs from Spotify
  - Pin locations using Google Maps)
- Memoria stylizes user input to create and display their collage page

## Gallery
![landing](https://user-images.githubusercontent.com/43392705/161670225-44433895-578a-499d-949e-d3c96a8f0e82.PNG)
![sign up](https://user-images.githubusercontent.com/43392705/161670251-a80adde8-2ea6-4f70-a06a-234fb845c1b8.PNG)
![add new](https://user-images.githubusercontent.com/43392705/161670282-beb5b25d-c01b-4df0-a8a4-29dc87b20225.PNG)
![sbs](https://user-images.githubusercontent.com/43392705/161670297-3dd39cc3-f530-4c5d-837e-5918b4fa574d.PNG)
![0](https://user-images.githubusercontent.com/43392705/161670313-f13da2a5-91ec-4109-934c-9102e891d010.PNG)
![image](https://user-images.githubusercontent.com/43392705/161670535-25bc73fe-c856-4f4e-86a0-b414f5bcaa44.png)
![captions](https://user-images.githubusercontent.com/43392705/161670333-302b68ea-f9f8-435d-9a13-899c16444999.PNG)
![music](https://user-images.githubusercontent.com/43392705/161670397-b079709d-4d32-4917-8c95-f2308f6ac57d.PNG)
![rome](https://user-images.githubusercontent.com/43392705/161670467-3b28b044-0417-4168-bf29-81a036f6ec34.PNG)
![image](https://user-images.githubusercontent.com/43392705/161670605-d466dbcb-9459-41a0-bc7f-9301b2d77d70.png)


## Memoria v1.0
The first release of Memoria has been published to our [Docker Hub registry.](https://hub.docker.com/r/afrah412000/memoria/tags) This is where all future releases will be published, in sync with the `release` branch of this repository.

## Software Documentation
### Installation & Deployment

- Deployment will be done using docker (a requirement to run the application)
- Navigate to the /app directory within the application
- Run `docker-compose build`, this will generate an executable of the Go application
- Run  `docker-compose up -d postgres`, this will start the postgres db in the background
- Run `docker-compose up app migrate`
   - this will start the go application at http://localhost:5000, along with a migration if needed
- Create an account and get started!

### 🎨 Design System
View the documentation of our Design System [here.](https://github.com/professor-forward/memoria/blob/f/deliverable-2/designSystem/README.md) 

### 🛠 Tools
- **Client-side:** HTML, CSS/SASS
- **Server-side:** Gin (Go)
- **Database:** Postgres
- **Internal API:** REST
- **External APIs:** Google Maps, Spotify

### Migrations

- the postgres DB has been versioned and migration files have been created and run through `go-migrate`
- To run a db migration on its own `docker-compose up migrate`
   - this will read the migrations/ file in the folder and build a new version of the system (if changes have been present)
- As the migrations were not generated through GORM (typeORM used), a change to the model requires changes in both the appropriate files in the migrations/ folder and the structures in models/ folder

### CI/CD
- Our DevOps pipeline is hosted on GitHub Actions where we build and test our code before integration. For release versions, we tag, build and publish Docker images of Memoria to our registry.

### Testing
- Unit testing for CRUD and API handlers are done through our test suite and automatically run as part of our CI pipeline. Tests can be run manually with `go test` from the `app/test` directory.

## Feature Development
### 1. Models
- Data models for Memoria exist in the `app/models` directory and are represented as go structs. This is where application entities can be extended for further feature development. 

### 2. Controllers
- User-to-model interaction is handled by the controllers in `app/controllers`, which handle database CRUD operations as well as hit external API endpoints of Google Maps and Spotify for user requests. 
- Routes and handlers are defined in `controllers/routes.go`, which can be extended to support new requests and connect to external APIs. 

### 3. Views
- Views are created with HTML5, CSS and Javascript in `app/views` and are routed through handlers. 

