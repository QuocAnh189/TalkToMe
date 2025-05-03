import { cn } from '@utils/cn'
import { IConversation } from '@interfaces/conversation'
import Avatar from './Avatar'

interface ConversationItemProps {
  conversation: IConversation
  active?: boolean
  onClick?: () => void
  className?: string
}

const ConversationItem = ({ conversation, active, onClick, className }: ConversationItemProps) => {
  return (
    <div
      className={cn(
        'flex items-center gap-3 p-3 cursor-pointer hover:bg-base-200 transition-colors',
        active && 'bg-base-200',
        className,
      )}
      onClick={onClick}
    >
      <Avatar src={conversation.partner?.avatarURL} alt={conversation.partner?.name} online={false} size="md" />
      <div className="flex-1 min-w-0">
        <div className="flex items-center justify-between">
          <h3 className="font-semibold truncate">{conversation.partner?.name}</h3>
          {conversation.lastMessage && (
            <span className="text-xs text-base-content/60">
              {new Date(conversation.lastMessage.createdAt).toLocaleDateString()}
            </span>
          )}
        </div>
        {conversation.lastMessage && (
          <p className="text-sm text-base-content/60 truncate">{conversation.lastMessage.message}</p>
        )}
      </div>
      {conversation.unreadCount > 0 && <div className="badge badge-primary badge-sm">{conversation.unreadCount}</div>}
    </div>
  )
}

export default ConversationItem
