import PageViews from '@yeger/page-views'
import type { FC } from 'react'
import { useEffect, useState } from 'react'

import '@/components/Footer.css'

const ViewCounter: FC = () => {
  const [pageViews, setPageViews] = useState<number>()

  useEffect(() => {
    PageViews.getViews().then(setPageViews)
  })

  if (pageViews === undefined) {
    return <span>Loading...</span>
  }

  return <span>{pageViews} views</span>
}

const Footer: FC = () => {
  return (
    <footer>
      <a
        href="https://github.com/DerYeger/npm-cards"
        target="_blank"
        rel="noreferrer"
      >
        GitHub
      </a>
      <span>Made by Jan with 🐹 and ⚛️</span>
      <ViewCounter />
    </footer>
  )
}

export default Footer
