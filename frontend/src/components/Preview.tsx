import CopyField from '@/components/CopyField'
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

  const cardUrl = lib.getCardUrl({
    packageName,
    size,
    padding,
    borderRadius,
    weeks,
  })

  const [isValid, setIsValid] = useState(false)
  useEffect(() => {
    fetch(cardUrl)
      .then((res) => setIsValid(res.status === 200))
      .catch(console.info)
  }, [cardUrl])

  if (!isValid) {
    return <span>Not found</span>
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
