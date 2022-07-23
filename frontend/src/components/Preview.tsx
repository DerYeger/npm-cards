import CopyField from '@/components/CopyField'
import Spinner from '@/components/Spinner'
import lib from '@/lib'
import { CardData } from '@/types'
import { FC, useEffect, useState } from 'react'

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
  const [isValid, setIsValid] = useState(false)

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
      try {
        const res = await fetch(cardUrl)
        setIsValid(res.status === 200)
      } catch (err) {}
      setIsLoading(false)
    }
    fetchCard()
  }, [cardUrl])

  if (!isValid) {
    return <span>Not found</span>
  }

  if (isLoading) {
    return <Spinner />
  }

  return (
    <>
      <CopyField cardUrl={cardUrl} />
      <img
        src={cardUrl}
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
