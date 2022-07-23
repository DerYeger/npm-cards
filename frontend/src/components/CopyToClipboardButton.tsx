import type { FC } from 'react'
import { CopyToClipboard } from 'react-copy-to-clipboard'
import { toast } from 'react-toastify'

const CopyToClipboardButton: FC<{ cardUrl: string }> = ({ cardUrl }) => {
  const notify = () => {
    const prefersLight = window.matchMedia(
      '(prefers-color-scheme: light)'
    ).matches
    toast('Copied to Clipboard', {
      pauseOnHover: false,
      theme: prefersLight ? 'light' : 'dark',
    })
  }
  return (
    <CopyToClipboard text={cardUrl} onCopy={notify}>
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

export default CopyToClipboardButton
