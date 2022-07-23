import debounce from 'lodash.debounce'
import { useEffect, useMemo, useState } from 'react'
import { ToastContainer } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'

import '@/App.css'
import Footer from '@/components/Footer'
import Preview from '@/components/Preview'
import type { CardData } from '@/types'

function App() {
  const params = new URLSearchParams(window.location.search)

  const [packageName, setPackageName] = useState(
    params.get('package') ?? '@yeger/vue-masonry-wall'
  )
  const [size, setSize] = useState(+(params.get('size') ?? 512))
  const [padding, setPadding] = useState(+(params.get('padding') ?? 0))
  const [borderRadius, setBordeRadius] = useState(
    +(params.get('borderRadius') ?? 16)
  )
  const [weeks, setWeeks] = useState(+(params.get('weeks') ?? 16))

  const [cardData, setCardData] = useState<CardData>({
    packageName,
    size,
    padding,
    borderRadius,
    weeks,
  })
  const debouncedSetCardData = useMemo(
    () =>
      debounce((cardData: CardData) => {
        params.set('package', cardData.packageName)
        params.set('size', cardData.size.toString())
        params.set('padding', cardData.padding.toString())
        params.set('borderRadius', cardData.borderRadius.toString())
        params.set('weeks', cardData.weeks.toString())
        window.history.replaceState({}, 'NPM Cards', `?${params.toString()}`)
        setCardData(cardData)
      }, 300),
    []
  )

  useEffect(() => {
    return () => {
      debouncedSetCardData.cancel()
    }
  }, [])

  const [contain, setContain] = useState(true)

  return (
    <div className="App">
      <main>
        <h1>NPM Cards</h1>
        <div className="settings">
          <div>
            <label
              htmlFor="packageName"
              style={{
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
              }}
            >
              Package
            </label>
            <input
              name="packageName"
              type="text"
              value={packageName}
              onChange={(event) => {
                setPackageName(event.target.value)
                debouncedSetCardData({
                  ...cardData,
                  packageName: event.target.value,
                })
              }}
              style={{
                flexGrow: '1',
              }}
            />
          </div>
          <div>
            <label htmlFor="size">Size ({size}px)</label>
            <input
              name="size"
              type="range"
              min="128"
              max="1024"
              value={size}
              onChange={(event) => {
                setSize(+event.target.value)
                debouncedSetCardData({ ...cardData, size: +event.target.value })
              }}
              step="1"
            />
          </div>
          <div>
            <label htmlFor="padding">Padding ({padding}px)</label>
            <input
              name="padding"
              type="range"
              min="0"
              max="128"
              value={padding}
              onChange={(event) => {
                setPadding(+event.target.value)
                debouncedSetCardData({
                  ...cardData,
                  padding: +event.target.value,
                })
              }}
              step="1"
            />
          </div>
          <div>
            <label htmlFor="borderRadius">
              Border Radius ({borderRadius}px)
            </label>
            <input
              name="borderRadius"
              type="range"
              min="0"
              max="128"
              value={borderRadius}
              onChange={(event) => {
                setBordeRadius(+event.target.value)
                debouncedSetCardData({
                  ...cardData,
                  borderRadius: +event.target.value,
                })
              }}
              step="1"
            />
          </div>
          <div>
            <label htmlFor="weeks">Weeks ({weeks})</label>
            <input
              name="weeks"
              type="range"
              min="2"
              max="200"
              value={weeks}
              onChange={(event) => {
                setWeeks(+event.target.value)
                debouncedSetCardData({
                  ...cardData,
                  weeks: +event.target.value,
                })
              }}
              step="1"
            />
          </div>
          <div>
            <label htmlFor="contain">Contain</label>
            <input
              type="checkbox"
              checked={contain}
              onChange={() => setContain(!contain)}
            ></input>
          </div>
        </div>
        <Preview {...cardData} contain={contain} />
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
