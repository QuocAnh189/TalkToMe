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
      <Avatar
        src={conversation.is_group ? conversation.members[0]?.avatarURL : conversation.members[1]?.avatarURL}
        alt={conversation.name}
        online={false}
        size="md"
      />
      <div className="flex-1 min-w-0">
        <div className="flex items-center justify-between">
          <h3 className="font-semibold truncate">{conversation.name}</h3>
          {conversation.last_message && (
            <span className="text-xs text-base-content/60">
              {new Date(conversation.last_message.created_at).toLocaleDateString()}
            </span>
          )}
        </div>
        {conversation.last_message && (
          <p className="text-sm text-base-content/60 truncate">
            {conversation.is_group && `${conversation.last_message.sender.name}: `}
            {conversation.last_message.content}
          </p>
        )}
      </div>
    </div>
  )
}

export default ConversationItem
