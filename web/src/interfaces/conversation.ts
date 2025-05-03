import { IUser } from "./user"
import { IMessage } from "./message"

export interface IConversation {
  id: string
  userId: string
  partnerId: string
  partner?: IUser
  lastMessage?: IMessage
  unreadCount: number
  createdAt: string
  updatedAt: string
}

export interface CreateConversationRequest {
  partnerId: string
}