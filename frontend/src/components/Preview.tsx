import type { FC } from 'react'
import { useEffect, useState } from 'react'

import CopyToClipboardButton from '@/components/CopyToClipboardButton'
import Spinner from '@/components/Spinner'
import lib from '@/lib'
import type { CardData } from '@/types'

export interface PreviewProps extends CardData {
  contain: boolean
}

const Preview: FC<PreviewProps> = ({
  packageName,
  size,
  padding,
  borderRadius,
  weeks,
  contain,
}) => {
  if (
    !lib.isCardDataComplete({ packageName, size, padding, borderRadius, weeks })
  ) {
    return <span>Missing input</span>
  }

  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<number | undefined>()
  const [image, setImage] = useState<string | undefined>()

  const cardUrl = lib.getCardUrl({
    packageName,
    size,
    padding,
    borderRadius,
    weeks,
  })

  useEffect(() => {
    async function fetchCard() {
      setIsLoading(true)
      setError(undefined)
      try {
        const res = await fetch(cardUrl)
        const blob = URL.createObjectURL(await res.blob())
        if (res.status === 200) {
          setImage(blob)
        } else {
          setError(res.status)
        }
      } catch (err) {}
      setIsLoading(false)
    }
    fetchCard()
  }, [cardUrl])

  if (isLoading) {
    return <Spinner />
  }

  if (error === 404) {
    return <span>Not found</span>
  }

  if (error !== undefined) {
    return <span>Something went wrong</span>
  }

  return (
    <>
      <CopyToClipboardButton cardUrl={cardUrl} />
      <img
        src={image}
        style={{
          maxWidth: contain ? '100%' : 'initial',
          display: 'block',
          marginLeft: 'auto',
          marginRight: 'auto',
        }}
      />
    </>
  )
}

export default Preview
