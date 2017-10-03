# Google container builder local REST API

```
dep ensure
```

```
godo server --watch
```

### Usage

```
curl -X POST -H 'Content-Type:application/json' \
-d '{
  "steps": [
    {
      "name": "gcr.io/cloud-builders/git",
      "entrypoint": "bash",
      "args": "date"
    }
  ],
  "logsBucket": "hogehoge"
}' \
http://localhost:1323/users
```