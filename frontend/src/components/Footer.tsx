import type { FC } from 'react'

import '@/components/Footer.css'

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
    </footer>
  )
}

export default Footer
