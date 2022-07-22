import { CardData } from '@/types'
import { FC, useEffect, useState } from 'react'

export interface PreviewProps extends CardData {
  contain: boolean
}

const apiEndpoint = import.meta.env.VITE_BACKEND_URL || 'http://localhost:8080'

const Preview: FC<PreviewProps> = ({
  packageName,
  size,
  padding,
  borderRadius,
  weeks,
  contain,
}) => {
  if (
    !packageName ||
    size === undefined ||
    padding === undefined ||
    borderRadius === undefined ||
    weeks === undefined
  ) {
    return <span>Missing input</span>
  }

  const cardSrc = `${apiEndpoint}/api/packages/${packageName}?size=${size}&padding=${padding}&borderRadius=${borderRadius}&weeks=${weeks}`
  const [isValid, setIsValid] = useState(false)
  useEffect(() => {
    fetch(cardSrc)
      .then((res) => setIsValid(res.status === 200))
      .catch(console.info)
  }, [cardSrc])

  if (!isValid) {
    return <span>Not found</span>
  }

  return (
    <img
      src={cardSrc}
      style={{
        maxWidth: contain ? '100%' : 'initial',
      }}
    />
  )
}

export default Preview
