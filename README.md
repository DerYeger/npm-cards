# NPM Cards

> Configurable downloads graphs for NPM packages.

## Backend

The lightweight Go backend uses [Gin](https://github.com/gin-gonic/gin) and in-memory or Redis-based [caching](https://github.com/chenyahui/gin-cache) for low latency.

Behind the scenes, the NPM API is queried to gather weekly download statistics for graph generation.

[SVGo](https://github.com/ajstarks/svgo) is used for generating SVGs programmatically.

## Frontend

The frontend is a React SPA built with Vite 3.
It serves as a playground for the backend by enabling experimentation with the configurable parameters.

It presents previews of cards and an option to copy card links to the clipboard.

## License

[MIT](./LICENSE) - Copyright &copy; Jan Müller
