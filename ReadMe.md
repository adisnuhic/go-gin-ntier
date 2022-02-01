This is basic setup of N tier arcitecture with Gin framework.

Arch is consisted of 4 layers:
- Controllers
- Business
- Services
- Repositories


1. Migrations are implemented with go migrate and they are executed when you run the project.
2. GORM is used as a ORM.
3. Entry point of the app is main.go file where everything starts. Golobby container is used to setup DI container.




