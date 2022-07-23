<h1 align="center">NPM Cards</h1>

<p align="center">
Configurable download graphs for NPM packages.
</p>

<p align="center">
  <a href="https://railway.app/new/template/pk27Nm?referralCode=FzqVFW">
    <img alt="Deploy on Railway" src="https://railway.app/button.svg">
  </a>
</p>

<p align="center">
  <a href="https://npm-cards.vercel.app">
    Frontend
  </a>
  —
  <a href="https://npm-cards.up.railway.app">
    Backend
  </a>
</p>

<p align="center">
  <a href="https://npm-cards.vercel.app/?package=react&size=256&padding=0&borderRadius=16&weeks=64">
    <img alt="React" src="https://npm-cards.up.railway.app/api/packages/react?size=256&padding=0&borderRadius=16&weeks=64">
  </a>
  <a href="https://npm-cards.vercel.app/?package=vite&size=256&padding=0&borderRadius=16&weeks=64">
    <img alt="Vite" src="https://npm-cards.up.railway.app/api/packages/vite?size=256&padding=0&borderRadius=16&weeks=64">
  </a>
  <a href="https://npm-cards.vercel.app/?package=@yeger/vue-masonry-wall&size=256&padding=0&borderRadius=16&weeks=64">
    <img alt="@yeger/vue-masonry-wall" src="https://npm-cards.up.railway.app/api/packages/@yeger/vue-masonry-wall?size=256&padding=0&borderRadius=16&weeks=64">
  </a>
</p>

## Backend

The lightweight Go backend uses [Gin](https://github.com/gin-gonic/gin) and in-memory or Redis-based [caching](https://github.com/chenyahui/gin-cache) for low latencies.

Behind the scenes, the NPM API is queried to gather weekly download statistics for graph generation.

[SVGo](https://github.com/ajstarks/svgo) is used for generating SVGs programmatically.

The backend is hosted by [Railway](https://railway.app) at [npm-cards.up.railway.app](https://npm-cards.up.railway.app).

## Frontend

The frontend is a React SPA built with Vite 3.
It serves as a playground for the backend by enabling experimentation with the configurable parameters.

It presents previews of cards and an option to copy card links to the clipboard.

The frontend is hosted by [Vercel](https://vercel.com) at [npm-cards.vercel.app](https://npm-cards.vercel.app).

## Usage

1. Visit [npm-cards.vercel.app](https://npm-cards.vercel.app)
2. Configure the card to fit your needs. While the frontend only allows for certain input ranges concerning size, padding, etc., modyfing the output URL allows for additional customization.
3. Press "Copy to Clipboard".
4. Embed the image URL in HTML, Markdown, etc.

## License

[MIT](./LICENSE) - Copyright &copy; Jan Müller
