# FC4 - OAuth 2 e OpenID Connect com Keycloak

```sh
# Run keycloak
$ docker run -p 8080:8080 -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin quay.io/keycloak/keycloak:11.0.1

# Run main.go
go run client/main.go
```