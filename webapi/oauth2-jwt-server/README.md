```
http://localhost:9090/authorize?response_type=code&client_id=clientID&redirect_uri=http://localhost:8080/callback
```

Provide username/password

```
curl -X POST http://localhost:9090/token \
  -d "grant_type=authorization_code" \
  -d "client_id=clientID" \
  -d "client_secret=clientSecret" \
  -d "code=YOUR_CODE_HERE" \
  -d "redirect_uri=http://localhost:8080/callback"
```

Decode returned JWT