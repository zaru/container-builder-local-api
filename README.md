# Google container builder local REST API

GCP container-builder-local REST API mock server.

### Prerequisites

- [container-builder-local](https://github.com/GoogleCloudPlatform/container-builder-local)

### Installation

- [Download binary](https://github.com/zaru/container-builder-local-api/releases)

### Usage

```
./container-builder-local-api
```

#### Create builds

```
curl -X POST -H 'Content-Type:application/json' \
-d '{
  "steps": [
    {
      "name": "gcr.io/cloud-builders/git",
      "entrypoint": "bash",
      "args": ["-c", "date"]
    }
  ],
  "logsBucket": "foobar"
}' \
http://localhost:1323/v1/projects/example-prj/builds | jq .
```

- ref: https://cloud.google.com/container-builder/docs/api/reference/rest/v1/projects.builds/create


### Development

```
dep ensure
```

```
realize run
```
