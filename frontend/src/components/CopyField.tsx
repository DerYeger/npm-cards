import { FC } from 'react'

import { CopyToClipboard } from 'react-copy-to-clipboard'

const CopyField: FC<{ cardUrl: string }> = ({ cardUrl }) => {
  return (
    <CopyToClipboard text={cardUrl}>
      <button
        style={{
          display: 'block',
          fontSize: '0.75rem',
          marginLeft: 'auto',
          marginRight: 'auto',
        }}
      >
        Copy to Clipboard
      </button>
    </CopyToClipboard>
  )
}

export default CopyField
