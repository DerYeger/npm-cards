import { useEffect, useMemo, useState } from 'react'
import debounce from 'lodash.debounce'

import Preview from '@/components/Preview'

import '@/App.css'
import { CardData } from '@/types'

function App() {
  const [packageName, setPackageName] = useState('@yeger/vue-masonry-wall')

  const [size, setSize] = useState(512)

  const [padding, setPadding] = useState(0)

  const [borderRadius, setBordeRadius] = useState(16)

  const [weeks, setWeeks] = useState(16)

  const [cardData, setCardData] = useState<CardData>({
    packageName,
    size,
    padding,
    borderRadius,
    weeks,
  })
  const debouncedSetCardData = useMemo(() => debounce(setCardData, 300), [])

  useEffect(() => {
    return () => {
      debouncedSetCardData.cancel()
    }
  }, [])

  const [contain, setContain] = useState(true)

  return (
    <div className="App">
      <h1>NPM Cards</h1>
      <div className="settings">
        <div>
          <label htmlFor="padding">Padding</label>
          <input
            name="name"
            type="text"
            value={packageName}
            onChange={(event) => {
              setPackageName(event.target.value)
              debouncedSetCardData({
                ...cardData,
                packageName: event.target.value,
              })
            }}
          />
        </div>
        <div>
          <label htmlFor="size">Size</label>
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
          <label htmlFor="padding">Padding</label>
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
          <label htmlFor="borderRadius">Border Radius</label>
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
          <label htmlFor="weeks">Weeks</label>
          <input
            name="weeks"
            type="range"
            min="2"
            max="200"
            value={weeks}
            onChange={(event) => {
              setWeeks(+event.target.value)
              debouncedSetCardData({ ...cardData, weeks: +event.target.value })
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
    </div>
  )
}

export default App
