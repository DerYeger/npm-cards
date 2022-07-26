import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import PageViews from '@yeger/page-views'
import React from 'react'
import ReactDOM from 'react-dom/client'

import App from '@/App'

import '@/index.css'

PageViews.autoSubmitViews()

const client = new QueryClient()

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <QueryClientProvider client={client}>
      <App />
    </QueryClientProvider>
  </React.StrictMode>
)
