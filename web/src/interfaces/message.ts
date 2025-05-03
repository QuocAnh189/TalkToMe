import { IUser } from "./user"
import { IMessageAttachment } from "./message_attachments"

export interface IMessage {
  id: string
  message: string
  senderId: string
  sender?: IUser|any
  groupId?: string
  conversationId?: string
  attachments?: IMessageAttachment[]
  createdAt: string
  updatedAt: string
}

export interface SendMessageRequest {
  groupId?: string
  conversationId?: string
  message: string
  attachments?: File[]
}

export interface UpdateMessageRequest {
  message: string
}