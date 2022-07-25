export const themes = [
  'dark',
  'dark-transparent',
  'light',
  'light-transparent',
] as const

export type Theme = typeof themes[number]

export interface CardData {
  packageName: string
  size: number
  padding: number
  borderRadius: number
  weeks: number
  theme: Theme
}
