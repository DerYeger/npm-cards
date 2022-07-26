import { useQuery } from '@tanstack/react-query'
import type { FC } from 'react'

import CopyToClipboardButton from '@/components/CopyToClipboardButton'
import Skeleton from '@/components/Skeleton/Skeleton'
import lib from '@/lib'
import type { CardData } from '@/types'

export interface PreviewProps {
  cardData: CardData
  contain: boolean
}

const Preview: FC<PreviewProps> = ({ cardData, contain }) => {
  const { isLoading, error, data } = useQuery(
    ['card', cardData],
    () => lib.fetchCard(cardData),
    { retry: false }
  )

  if (isLoading) {
    return (
      <Skeleton width={`${cardData.size}px`} height={`${cardData.size}px`} />
    )
  }

  if (error === 404) {
    return <span>Not found</span>
  }

  if (error) {
    const message =
      error instanceof Error ? error.message : 'Something went wrong'
    return <span>{message}</span>
  }

  return (
    <>
      <CopyToClipboardButton cardUrl={lib.getCardUrl(cardData)} />
      <img
        src={data}
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
