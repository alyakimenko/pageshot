<p align="center">
<img
    src="assets/logo.png"
    width="350" height="112" border="0" alt="pageshot">
</p>

Pageshot is a simple self-hosted tool that enables you to take a screenshot of any webpage.

## Features

- Containerized (browser is delivered within a Docker image).
- Configurable width and height parameters.
- Ability to take a full page screenshot.
- Support for a delayed screenshot.
- Support for three output image formats (`PNG`, `JPEG` and `WEBP`).
- Flexible scale and quality factors.
- Cache screenshots with local file storage or S3.

## API

`GET /screenshot`

| Param      |   Type   | Description                                               |   Default    |
| :--------- | :------: | :-------------------------------------------------------- | :----------: |
| `url`      | `string` | Target webpage URL                                        | **Required** |
| `width`    |  `int`   | Viewport width                                            |     1440     |
| `height`   |  `int`   | Viewport height                                           |     900      |
| `scale`    |  `int`   | Viewport scale factor                                     |     1.0      |
| `format`   | `string` | Output image format (png, jpeg, webp)                     |     png      |
| `quality`  |  `int`   | Output image quiality                                     |      70      |
| `delay`    |  `int`   | Delay in milliseconds, to wait before taking a screenshot |      0       |
| `fullpage` |  `bool`  | Capture full page screenshot                              |    false     |

## Config

Config is based on environmental variables.

| Environmental variable               |    Type    |                                                           Description                                                            |    Default    |
| :----------------------------------- | :--------: | :------------------------------------------------------------------------------------------------------------------------------: | :-----------: |
| `SERVER_PORT` or `PORT` (for Heroku) |   `int`    |                                                          Server's port                                                           |     8000      |
| `SERVER_READ_TIMEOUT`                | `duration` |                               Maximum duration for reading the entire request, including the body                                |      5s       |
| `SERVER_WRITE_TIMEOUT`               | `duration` |                                    Maximum duration before timing out writes of the response                                     |      15s      |
| `SERVER_IDLE_TIMEOUT`                | `duration` |                         Maximum amount of time to wait for the next request when keep-alives are enabled                         |      5s       |
| `BROWSER_WIDTH`                      |   `int`    |                                                 Initial browser's viewport width                                                 |     1440      |
| `BROWSER_HEIGHT`                     |   `int`    |                                                Initial browser's viewport height                                                 |      900      |
| `BROWSER_URL`                        |  `string`  | Remote browser's URL. If specified, pageshot will try to connect to a remote browser, otherwise will try to allocate a local one |               |
| `STORAGE_TYPE`                       |  `string`  |                                  Type of a storage. **local** and **s3** are currently allowed                                   |               |
| `STORAGE_LOCAL_DIRECTORY`            |  `string`  |                           Path to local storage directory. Applicable if **local** storage is selected                           | temp dir path |
| `STORAGE_S3_BUCKET`                  |  `string`  |                                                            S3 bucket                                                             |               |
| `STORAGE_S3_ENDPOINT`                |  `string`  |                                                           S3 endpoint                                                            |               |
| `STORAGE_S3_ACCESS_KEY_ID`           |  `string`  |                                                         S3 Access Key ID                                                         |               |
| `STORAGE_S3_SECRET_ACCESS_KEY`       |  `string`  |                                                       S3 Secret Access Key                                                       |               |
| `STORAGE_S3_SSL`                     | `boolean`  |                                                      Use SSL for S3 storage                                                      |     false     |
| `LOGGER_LEVEL`                       |  `string`  |                                                          Logger's level                                                          |     INFO      |
