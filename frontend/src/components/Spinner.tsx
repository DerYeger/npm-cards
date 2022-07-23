import { FC } from 'react'

import '@/components/Spinner.css'

const Spinner: FC = () => {
  return (
    <div className="lds-ripple">
      <div></div>
      <div></div>
    </div>
  )
}

export default Spinner
