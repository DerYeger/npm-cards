import PageViews from '@yeger/page-views'
import React from 'react'
import ReactDOM from 'react-dom/client'

import App from '@/App'

import '@/index.css'

PageViews.autoSubmitViews()

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)
