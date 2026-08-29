# Nummus.app

Nummus is a self hosted home finance tracker built with Go, React and Postgres.

> This project was born from the need to track my expenses at home and the fact that i wanted to learn about backend development and the go programming language.

## Tech track
### Server
Backend API built with [Go](https://go.dev/) and [Gin](https://gin-gonic.com/) for the services, [PostgreSQL](https://www.postgresql.org/) for the database and [Goose](https://github.com/pressly/goose) for managing database migrations.

### FrontEnd
Web Application built with [React](https://react.dev/), [react router](https://reactrouter.com/) in the declarative mode, [tailwind](https://tailwindcss.com/) for styles and [shadcn](https://ui.shadcn.com/) for the design system.

### Deployment
The application is designed to be deployed using [Docker](https://www.docker.com/) with a single [Docker Compose](https://docs.docker.com/compose/) file allowing it to be ran in any local machine with ease.

