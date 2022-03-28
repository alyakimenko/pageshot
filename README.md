# 📸 Pageshot

Pageshot is a simple self-hosted tool that enables you to take a screenshot of any webpage.

## Features

- Containerized (browser is delivered within a Docker image).
- Supports configured width and height.
- Supports fullpage screenshots.
- Supports delayed screenshots.
- Supports three basic image formats (PNG, JPEG and WEBP).
- Supports scale and quality factors.
- Supports local file storage.

## TODO

- [ ] Screenshots by specific CSS selector.
- [ ] Scroll entire webpage option.
- [ ] Scroll to specific element option.
- [ ] Custom User-Agent and others HTTP headers.
- [ ] Proxy support.
- [ ] Execute custom JavaScript before taking a screenshot.
- [ ] Setup building, publishing, and deploying Docker images with Github Actions.

## API

`GET /screenshot`

| Param         |   Type    | Description                                                   |   Default    |
| :------------ | :-------: | :------------------------------------------------------------ | :----------: |
| `url`         | `string`  | Target webpage URL                                            | **Required** |
| `width`       |   `int`   | Viewport width                                                |     1440     |
| `height`      |   `int`   | Viewport height                                               |     900      |
| `scale`       |   `int`   | Viewport scale factor                                         |     1.0      |
| `format`      | `string`  | Output image format (png, jpeg, webp)                         |     png      |
| `quality`     |   `int`   | Output image quiality                                         |     70       |
| `delay`       |   `int`   | Delay in milliseconds, to wait before taking a screenshot     |     0        |
| `fullpage`    |   `bool`  | Capture full webpage screenshot                               |    false     |