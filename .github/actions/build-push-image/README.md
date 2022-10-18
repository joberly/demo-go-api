<!-- start title -->

# GitHub Action: build-push-image

<!-- end title -->
<!-- start description -->

Build and push multi-arch image to GHCR using docker buildx

<!-- end description -->
<!-- start contents -->
<!-- end contents -->
<!-- start usage -->

```yaml
- uses: ./.github/actions/build-push-image@main
  with:
    # Name of image to build.
    name: ""

    # Comma separated list of platforms to build
    # Default: linux/amd64
    platforms: ""

    # Context directory for build containing Dockerfile
    # Default: ./
    context: ""

    # Path to Dockerfile
    # Default: ./Dockerfile
    file: ""

    # Set latest tag to this image
    # Default: false
    latest: ""

    # Username for ghcr.io
    username: ""

    # Password for ghcr.io
    password: ""
```

<!-- end usage -->
<!-- start inputs -->

| **Input**       | **Description**                                   | **Default**    | **Required** |
| --------------- | ------------------------------------------------- | -------------- | ------------ |
| **`name`**      | Name of image to build.                           |                | **true**     |
| **`platforms`** | Comma separated list of platforms to build        | `linux/amd64`  | **false**    |
| **`context`**   | Context directory for build containing Dockerfile | `./`           | **false**    |
| **`file`**      | Path to Dockerfile                                | `./Dockerfile` | **false**    |
| **`latest`**    | Set latest tag to this image                      | `false`        | **false**    |
| **`username`**  | Username for ghcr.io                              |                | **true**     |
| **`password`**  | Password for ghcr.io                              |                | **true**     |

<!-- end inputs -->
<!-- start outputs -->
<!-- end outputs -->
<!-- start [.github/ghadocs/examples/] -->
<!-- end [.github/ghadocs/examples/] -->
