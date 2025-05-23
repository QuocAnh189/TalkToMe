import { cn } from '@utils/cn'
import { IMessage } from '@interfaces/message'
import Avatar from './Avatar'
import Tooltip from './Tooltip'

interface MessageBubbleProps {
  message: IMessage
  isOwn?: boolean
  showAvatar?: boolean
  className?: string
}

const MessageBubble = ({ message, isOwn = false, showAvatar = true, className }: MessageBubbleProps) => {
  return (
    <div className={cn('flex gap-2', isOwn ? 'flex-row-reverse' : 'flex-row', className)}>
      {showAvatar && (
        <Avatar src={message.sender?.avatar_url} alt={message.sender?.name} size="sm" className="flex-shrink-0" />
      )}
      <div className={cn('flex flex-col gap-1 max-w-[70%]', isOwn && 'items-end')}>
        {!isOwn && <span className="text-xs text-base-content/60">{message.sender?.name}</span>}
        <Tooltip content={new Date(message.createdAt).toLocaleTimeString()}>
          <div className={cn('px-4 py-2 rounded-2xl', isOwn ? 'bg-primary text-primary-content' : 'bg-base-200')}>
            {message.message}
          </div>
        </Tooltip>
      </div>
    </div>
  )
}

export default MessageBubble
