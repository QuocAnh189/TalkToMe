//hooks
import { useState, useRef } from 'react'
import { useSendMessageMutation } from '@redux/services/message'

//ui
import Input from '@ui/Input'

//icon
import IconButton from '@ui/IconButton'
import { GiPaperClip } from 'react-icons/gi'
import { BsSendFill } from 'react-icons/bs'

const MessageInput = () => {
  const [message, setMessage] = useState('')
  const fileInputRef = useRef<HTMLInputElement>(null)
  const [sendMessage] = useSendMessageMutation()

  const handleSend = async () => {
    if (!message.trim()) return

    await sendMessage({
      conversationId: '123',
      message: message.trim(),
    })

    setMessage('')
  }

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files

    if (!files) return
  }

  return (
    <div className="p-4 border-t border-base-200">
      <div className="flex items-center gap-2">
        <IconButton tooltip="Add Files" onClick={() => fileInputRef.current?.click()}>
          <GiPaperClip className="w-5 h-5" />
        </IconButton>
        <Input
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          placeholder="Aa"
          className="flex-1"
          onKeyDown={(e) => e.key === 'Enter' && handleSend()}
        />
        <IconButton variant="primary" disabled={!message.trim()} onClick={handleSend}>
          <BsSendFill className="w-5 h-5" />
        </IconButton>
      </div>
      <input type="file" ref={fileInputRef} className="hidden" multiple onChange={handleFileChange} />
    </div>
  )
}

export default MessageInput
