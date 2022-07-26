import 'react-toastify/dist/ReactToastify.css'
import { useState } from 'react'
import { ToastContainer } from 'react-toastify'
import { useDebounce } from 'use-debounce'

import '@/App.css'
import Footer from '@/components/Footer'
import Preview from '@/components/Preview'
import Settings from '@/components/Settings'
import type { CardData, Theme } from '@/types'
import { themes } from '@/types'

function App() {
  const params = new URLSearchParams(window.location.search)

  const preselectedTheme = themes.includes(params.get('theme') as Theme)
    ? (params.get('theme') as Theme)
    : 'dark'

  const [cardData, setCardData] = useState<CardData>({
    packageName: params.get('package') ?? '@yeger/vue-masonry-wall',
    size: +(params.get('size') ?? 512),
    padding: +(params.get('padding') ?? 0),
    borderRadius: +(params.get('borderRadius') ?? 16),
    weeks: +(params.get('weeks') ?? 16),
    theme: preselectedTheme,
  })

  const [contain, setContain] = useState(true)

  const onSettingsChange = (cardData: CardData, contain: boolean) => {
    params.set('package', cardData.packageName)
    params.set('size', cardData.size.toString())
    params.set('padding', cardData.padding.toString())
    params.set('borderRadius', cardData.borderRadius.toString())
    params.set('weeks', cardData.weeks.toString())
    params.set('theme', cardData.theme)
    window.history.replaceState({}, 'NPM Cards', `?${params.toString()}`)
    setCardData(cardData)
    setContain(contain)
  }

  const [debouncedCardData] = useDebounce(cardData, 300)

  return (
    <div className="App">
      <main>
        <h1>NPM Cards</h1>
        <Settings
          cardData={cardData}
          contain={contain}
          themes={[...themes]}
          onChange={onSettingsChange}
        />
        <Preview cardData={debouncedCardData} contain={contain} />
      </main>
      <Footer />
      <ToastContainer
        autoClose={1000}
        pauseOnFocusLoss={false}
        hideProgressBar
      />
    </div>
  )
}

export default App
