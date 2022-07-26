import type { CardData } from '@/types'

const apiEndpoint = import.meta.env.VITE_BACKEND_URL || 'http://localhost:8080'

const getCardUrl = ({
  packageName,
  size,
  padding,
  borderRadius,
  weeks,
  theme,
}: CardData) => {
  return `${apiEndpoint}/api/packages/${packageName}?size=${size}&padding=${padding}&borderRadius=${borderRadius}&weeks=${weeks}&theme=${theme}`
}

const isCardDataComplete = ({
  packageName,
  size,
  padding,
  borderRadius,
  weeks,
  theme,
}: CardData) => {
  return (
    packageName &&
    size !== undefined &&
    padding !== undefined &&
    borderRadius !== undefined &&
    weeks !== undefined &&
    theme !== undefined
  )
}

const fetchCard = async (cardData: CardData) => {
  if (!isCardDataComplete(cardData)) {
    throw new Error('Missing Input')
  }

  const cardUrl = getCardUrl(cardData)

  const res = await fetch(cardUrl, { method: 'GET' })

  if (res.status === 429) {
    throw new Error('We are being rate-limited :(')
  }

  if (res.status !== 200) {
    throw new Error(res.statusText)
  }

  return URL.createObjectURL(await res.blob())
}

export default {
  getCardUrl,
  isCardDataComplete,
  fetchCard,
}
