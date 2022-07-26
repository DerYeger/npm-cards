import type { FC } from 'react'

import '@/components/Skeleton/Skeleton.css'

const Skeleton: FC<{ width: string; height: string }> = ({ width, height }) => {
  return <div style={{ width, height }} className="skeleton" />
}

export default Skeleton
