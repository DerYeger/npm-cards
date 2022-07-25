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
    (packageName && size !== undefined) ||
    padding !== undefined ||
    borderRadius !== undefined ||
    weeks !== undefined ||
    theme !== undefined
  )
}

export default {
  getCardUrl,
  isCardDataComplete,
}
