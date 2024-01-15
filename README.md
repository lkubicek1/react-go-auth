# Full Stack Authentication Demo with React, Vite, Go, and Auth0

## Overview
This project demonstrates a full-stack application using React (front-end), Vite, Go (back-end), and Auth0 for user authentication. It serves as a starting point for developers looking to implement user authentication in their applications.

## Prerequisites
- [Node.js](https://nodejs.org/en/download/)
- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started) (optional for containerization)
- An Auth0 account (Create one at [Auth0](https://auth0.com))

## Auth0 Setup
1. **Create an Auth0 Tenant:**
   - Sign up/Login to Auth0.
   - Create a new tenant by providing a unique name and region.

2. **Register an Application:**
   - Go to 'Applications' → 'Applications' in the Auth0 dashboard.
   - Click 'Create Application'.
   - Name your application, select 'Single Page Web Applications', and click 'Create'.

3. **Register an API:**
   - In the Auth0 dashboard, go to 'Applications' → 'APIs'.
   - Click 'Create API'.
   - Enter a name, identifier (this will be your API's audience), and select a signing algorithm.

4. **Configuration:**
   - In your application settings, set the 'Allowed Callback URLs', 'Allowed Logout URLs', and 'Allowed Web Origins' to match your application's URLs.

## Environment Setup
Create `.env` files in both the front-end and back-end directories. Include the following:

### Frontend `.env`

``` sh 
VITE_AUTH0_DOMAIN=<Your Auth0 Domain>
VITE_AUTH0_CLIENT_ID=<Your Auth0 Client ID>
VITE_AUTH0_AUDIENCE=<Your API Identifier>
```

### Backend `.env`
```
AUTH0_DOMAIN=<Your Auth0 Domain>
AUTH0_AUDIENCE=<Your API Identifier>
```


## Running the Application
1. **Backend:**
   - Navigate to the backend directory.
   - Run `go build` to build the Go server.
   - Start the server using `./[binary-name]`.

2. **Frontend:**
   - Navigate to the frontend directory.
   - Install dependencies with `npm install`.
   - Run `npm run dev` to start the development server.

## Docker (Optional)
If you prefer to use Docker:
- Use the provided `Dockerfile` and `docker-compose.yml` in the respective directories to build and run containers for both the front-end and back-end.

## License
This project is licensed under the Apache License 2.0. For the full license text, please refer to the `LICENSE` file in the repository or visit the [official Apache License 2.0 webpage](https://www.apache.org/licenses/LICENSE-2.0).

## Support
For support, open an issue in the GitHub repository or refer to the Auth0 documentation for specific authentication queries.

---

This README aims to assist developers in understanding and implementing a full-stack authentication system using React, Vite, Go, and Auth0, under the Apache 2.0 license.
