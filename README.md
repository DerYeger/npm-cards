# NPM Cards

> Configurable download graphs for NPM packages.

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/new/template/pk27Nm?referralCode=FzqVFW)

## Backend

The lightweight Go backend uses [Gin](https://github.com/gin-gonic/gin) and in-memory or Redis-based [caching](https://github.com/chenyahui/gin-cache) for low latency.

Behind the scenes, the NPM API is queried to gather weekly download statistics for graph generation.

[SVGo](https://github.com/ajstarks/svgo) is used for generating SVGs programmatically.

The backend is hosted by [Railway](https://railway.app) at [npm-cards.up.railway.app](https://npm-cards.up.railway.app).

## Frontend

The frontend is a React SPA built with Vite 3.
It serves as a playground for the backend by enabling experimentation with the configurable parameters.

It presents previews of cards and an option to copy card links to the clipboard.

The frontend is hosted by [Vercel](https://vercel.com) at [npm-cards.vercel.app](https://npm-cards.vercel.app).

## License

[MIT](./LICENSE) - Copyright &copy; Jan Müller
