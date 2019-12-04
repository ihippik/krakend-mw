# KrakenD middleware

* Relying party (Token-Based Authentication with user roles)

middleware extra config:
```json
  "extra_config": {
    "github_com/ihippik/krakend-mw/relyingparty": {
      "token_secret": "my-token-secret"
    }
  },
```

endpoint extra config:
```json
  "extra_config": {
    "github.com/ihippik/krakend-mw/relyingparty": {
      "roles": ["admin","member"]
    }
  },
```