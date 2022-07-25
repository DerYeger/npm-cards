import PageViews from '@yeger/page-views'
import type { FC } from 'react'
import { useEffect, useState } from 'react'

import '@/components/Footer.css'

const Footer: FC = () => {
  const [pageViews, setPageViews] = useState<number>()

  useEffect(() => {
    PageViews.getViews().then(setPageViews)
  })

  const pageViewElement =
    pageViews !== undefined ? (
      <span>{pageViews} views</span>
    ) : (
      <span>Loading...</span>
    )

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
      {pageViewElement}
    </footer>
  )
}

export default Footer
