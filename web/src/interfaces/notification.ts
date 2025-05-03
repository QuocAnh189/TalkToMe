export interface INotification {
  id: string
  userId: string
  type: string
  title: string
  content: string
  read: boolean
  data: any
  createdAt: string
  updatedAt: string
}