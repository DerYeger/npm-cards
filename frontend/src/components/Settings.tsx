import type { FC } from 'react'
import Select from 'react-select'

import type { CardData, Theme } from '@/types'

export interface SettingsProps {
  cardData: CardData
  contain: boolean
  onChange: (cardData: CardData, contain: boolean) => void
  themes: Theme[]
}

const Settings: FC<SettingsProps> = ({
  cardData,
  contain,
  onChange,
  themes,
}) => {
  const themeOptions = themes.map((theme) => ({
    value: theme,
    label: theme,
  }))

  return (
    <div className="settings">
      <div>
        <label
          htmlFor="packageName"
          className="flex items-center justify-center"
        >
          Package
        </label>
        <input
          name="packageName"
          className="flex-1"
          type="text"
          value={cardData.packageName}
          onChange={(event) =>
            onChange({ ...cardData, packageName: event.target.value }, contain)
          }
        />
      </div>
      <div>
        <label htmlFor="size">Size ({cardData.size}px)</label>
        <input
          name="size"
          type="range"
          min="128"
          max="1024"
          value={cardData.size}
          onChange={(event) =>
            onChange({ ...cardData, size: +event.target.value }, contain)
          }
          step="1"
        />
      </div>
      <div>
        <label htmlFor="padding">Padding ({cardData.padding}px)</label>
        <input
          name="padding"
          type="range"
          min="0"
          max="128"
          value={cardData.padding}
          onChange={(event) =>
            onChange({ ...cardData, padding: +event.target.value }, contain)
          }
          step="1"
        />
      </div>
      <div>
        <label htmlFor="borderRadius">
          Border Radius ({cardData.borderRadius}px)
        </label>
        <input
          name="borderRadius"
          type="range"
          min="0"
          max="128"
          value={cardData.borderRadius}
          onChange={(event) =>
            onChange(
              { ...cardData, borderRadius: +event.target.value },
              contain
            )
          }
          step="1"
        />
      </div>
      <div>
        <label htmlFor="weeks">Weeks ({cardData.weeks})</label>
        <input
          name="weeks"
          type="range"
          min="2"
          max="200"
          value={cardData.weeks}
          onChange={(event) =>
            onChange({ ...cardData, weeks: +event.target.value }, contain)
          }
          step="1"
        />
      </div>
      <div>
        <label htmlFor="theme" className="flex items-center justify-center">
          Theme
        </label>
        <Select
          name="theme"
          value={{ value: cardData.theme, label: cardData.theme }}
          options={themeOptions}
          styles={{
            option: (provided, option) => ({
              ...provided,
              color: option.isSelected ? provided.color : 'black',
            }),
            control: (provided) => ({ ...provided, width: 200 }),
          }}
          onChange={(change) => {
            if (change == null) {
              return
            }
            onChange({ ...cardData, theme: change.value }, contain)
          }}
        />
      </div>
      <div>
        <label htmlFor="contain">Contain</label>
        <input
          type="checkbox"
          checked={contain}
          onChange={() => onChange(cardData, !contain)}
        />
      </div>
    </div>
  )
}

export default Settings
