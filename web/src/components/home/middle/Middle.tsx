// import { useListMessagesQuery } from '@redux/services/message'

//components
import ConversationHeader from './ConversationHeader'
import MessageInput from './MessageInput'

//ui
import MessageBubble from '@ui/MessageBubble'
import ScrollableChat from '@ui/ScrollableChat'
import TypingIndicator from '@ui/TypingIndicator'

//interfaces
import { IMessage } from '@interfaces/message'

//data
import { mockMessages } from '../../../data/messages'

const Middle = () => {
  // const { data: messages, isLoading } = useListMessagesQuery(
  //   { conversationId: activeChat?.id },
  //   { skip: !activeChat }
  // )

  const messages: any = mockMessages
  const isTyping = false

  return (
    <div className="flex-1 h-screen flex flex-col bg-[#e2e8f0]">
      <ConversationHeader />

      <ScrollableChat className="flex-1 p-4 gap-4">
        {messages?.map((message: IMessage, index: number) => (
          <MessageBubble
            key={message.id}
            message={message}
            isOwn={false}
            showAvatar={index === 0 || messages[index - 1]?.senderId !== message.senderId}
          />
        ))}
        {isTyping && <TypingIndicator />}
      </ScrollableChat>

      <MessageInput />
    </div>
  )
}

export default Middle
