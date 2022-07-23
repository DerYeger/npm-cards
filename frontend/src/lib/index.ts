import type { CardData } from '@/types'

const apiEndpoint = import.meta.env.VITE_BACKEND_URL || 'http://localhost:8080'

const getCardUrl = ({
  packageName,
  size,
  padding,
  borderRadius,
  weeks,
}: CardData) => {
  return `${apiEndpoint}/api/packages/${packageName}?size=${size}&padding=${padding}&borderRadius=${borderRadius}&weeks=${weeks}`
}

const isCardDataComplete = ({
  packageName,
  size,
  padding,
  borderRadius,
  weeks,
}: CardData) => {
  return (
    (packageName && size !== undefined) ||
    padding !== undefined ||
    borderRadius !== undefined ||
    weeks !== undefined
  )
}

export default {
  getCardUrl,
  isCardDataComplete,
}
