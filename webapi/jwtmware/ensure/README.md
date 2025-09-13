An example of HTTP server with authentication powered by JWT.

JWT is issued upon successful login via response body (JSON).

Server expects JWT to be provided by client via Authorization request header (basic authentication).

JWT secret key is provided via environment variable JWT_SECRET

Inspired by [cybertec-postgresql/pgwatch](https://github.com/cybertec-postgresql/pgwatch)