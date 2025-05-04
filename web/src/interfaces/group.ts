import { IUser } from "./user"
import { IMessage } from "./message"

export interface IGroup {
  id: string
  name: string
  description: string
  avatar_url: string
  ownerId: string
  owner?: IUser
  members?: IUser[]
  lastMessage?: IMessage
  unreadCount: number
  createdAt: string
  updatedAt: string
}

export interface CreateGroupRequest {
  name: string
  description: string
  avatar?: File
  memberIds: string[]
}

export interface UpdateGroupRequest {
  name?: string
  description?: string
  avatar?: File
}