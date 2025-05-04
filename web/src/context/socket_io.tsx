import { PropsWithChildren, createContext, useEffect, useRef, useState } from 'react'

interface Message {
  type: string
  target_user_id?: string
  group_id?: string
  content?: string
  requester_user_id?: string
  sender_user_id?: string
}

export interface AppSocketIOContextProps {
  socket: WebSocket | null
  isConnected: boolean
  connect: () => void
  disconnect: () => void
  sendPrivateMessage: (targetUserId: string, content: string) => void
  sendGroupMessage: (groupId: string, content: string) => void
  joinGroup: (groupId: string) => void
  sendFriendRequest: (targetUserId: string) => void
  acceptFriendRequest: (requesterUserId: string) => void
  addUserToGroup: (groupId: string, targetUserId: string) => void
  removeUserFromGroup: (groupId: string, targetUserId: string) => void
}

export const AppSocketContext = createContext<Partial<AppSocketIOContextProps>>({})

const AppSocketIOProvider = ({ children }: PropsWithChildren) => {
  const user = JSON.parse(localStorage.getItem('user') || '{}')
  const socketRef = useRef<WebSocket | null>(null)
  const [isConnected, setIsConnected] = useState<boolean>(false)

  const connect = () => {
    if (!socketRef.current && user?.id) {
      const wsUrl = `${import.meta.env.VITE_SOCKET_URL}/ws?user_id=${user.id}`
      socketRef.current = new WebSocket(wsUrl)

      socketRef.current.onopen = () => {
        console.log('WebSocket connected')
        setIsConnected(true)
      }

      socketRef.current.onclose = () => {
        console.log('WebSocket disconnected')
        setIsConnected(false)
        // Attempt to reconnect after 5 seconds
        // setTimeout(() => {
        //   socketRef.current = null
        //   connect()
        // }, 5000)
      }

      socketRef.current.onerror = (error) => {
        console.error('WebSocket error:', error)
        setIsConnected(false)
      }

      socketRef.current.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data)
          console.log('Received message:', data)
          // Handle different message types here based on data.type
          // You can emit events or update state based on the message type
        } catch (error) {
          console.error('Error parsing message:', error)
        }
      }
    }
  }

  const disconnect = () => {
    if (socketRef.current) {
      socketRef.current.close()
      socketRef.current = null
      setIsConnected(false)
    }
  }

  const sendMessage = (message: Message) => {
    if (socketRef.current && isConnected) {
      try {
        socketRef.current.send(JSON.stringify(message))
      } catch (error) {
        console.error('Error sending message:', error)
      }
    }
  }

  const sendPrivateMessage = (targetUserId: string, content: string) => {
    sendMessage({
      type: 'private_message',
      target_user_id: targetUserId,
      content: content,
    })
  }

  const sendGroupMessage = (groupId: string, content: string) => {
    sendMessage({
      type: 'group_message',
      group_id: groupId,
      content: content,
    })
  }

  const joinGroup = (groupId: string) => {
    sendMessage({
      type: 'join_group',
      group_id: groupId,
    })
  }

  const sendFriendRequest = (targetUserId: string) => {
    sendMessage({
      type: 'friend_request',
      target_user_id: targetUserId,
    })
  }

  const acceptFriendRequest = (requesterUserId: string) => {
    sendMessage({
      type: 'accept_friend_request',
      requester_user_id: requesterUserId,
    })
  }

  const addUserToGroup = (groupId: string, targetUserId: string) => {
    sendMessage({
      type: 'add_to_group',
      group_id: groupId,
      target_user_id: targetUserId,
    })
  }

  const removeUserFromGroup = (groupId: string, targetUserId: string) => {
    sendMessage({
      type: 'remove_from_group',
      group_id: groupId,
      target_user_id: targetUserId,
    })
  }

  useEffect(() => {
    if (user?.id) {
      connect()
    }

    return () => {
      disconnect()
    }
  }, [user?.id])

  const value = {
    socket: socketRef.current,
    isConnected,
    connect,
    disconnect,
    sendPrivateMessage,
    sendGroupMessage,
    joinGroup,
    sendFriendRequest,
    acceptFriendRequest,
    addUserToGroup,
    removeUserFromGroup,
  }

  return <AppSocketContext.Provider value={value}>{children}</AppSocketContext.Provider>
}

export default AppSocketIOProvider
