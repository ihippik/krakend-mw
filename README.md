# KrakenD middleware

* Relying party (Token-Based Authentication with user roles)

    middleware will check for each request the validity of the jwt-token (expiration time, signature), as well as the correspondence of the user's role with the allowed roles for this endpoint.

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