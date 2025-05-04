//hooks
import { useContext, useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAppDispatch } from '@hooks/useRedux'
import { useSignOutMutation } from '@redux/services/auth'
// import { useListConversationsQuery } from '@redux/services/conversation'
// import { useListGroupsQuery } from '@redux/services/group'

//interfaces
import { IConversation } from '@interfaces/conversation'
// import { IGroup } from '@interfaces/group'

//components
import CreateGroupModal from '../modal/CreateGroupModal'
import SearchUsersModal from '../modal/SearchUsersModal'
import NotificationModal from '../modal/NotificationModal'
import toast from 'react-hot-toast'

//ui
import Avatar from '@ui/Avatar'
import Input from '@ui/Input'
import IconButton from '@ui/IconButton'
import ConversationItem from '@ui/ConversationItem'
import Dropdown from '@ui/Dropdown'

//icons
import { BiPencil, BiSearch, BiBell } from 'react-icons/bi'
import { CgMore } from 'react-icons/cg'

//data
import { mockConversations } from '../../../data/conversation'
import { mockAvailableUsers, mockAvailableSearchUsers } from '../../../data/user'

//context
import { AppSocketContext } from 'context/socket_io'

//store
import { setAuth } from '@redux/slices/auth.slice'
import Button from '@ui/Button'

const LeftSidebar = () => {
  const navigate = useNavigate()
  const dispatch = useAppDispatch()
  const [searchTerm, setSearchTerm] = useState<string>('')
  const { sendPrivateMessage, isConnected, socket } = useContext(AppSocketContext)
  // const { data: conversations } = useListConversationsQuery({})
  // const { data: groups } = useListGroupsQuery({})

  const user = JSON.parse(localStorage.getItem('user') || '{}')
  const [SignOut] = useSignOutMutation()

  const conversations = mockConversations

  const handleLogout = async () => {
    try {
      const result = await SignOut()
      if (result) {
        dispatch(setAuth(null))
        localStorage.removeItem('token')
        localStorage.removeItem('user')
        toast.success('Logout successfully.')
        navigate('/login')
      }
    } catch (e: any) {
      toast.error('Something went wrong.')
    }
  }

  const [showCreateGroupModal, setShowCreateGroupModal] = useState<boolean>(false)
  const [showSearchUsersModal, setShowSearchUsersModal] = useState<boolean>(false)
  const [showNotificationModal, setShowNotificationModal] = useState<boolean>(false)

  const handleCreateGroup = (data: { name: string; description: string; avatar?: File; memberIds: string[] }) => {
    console.log('Creating group:', data)
    setShowCreateGroupModal(false)
  }

  const handleAddFriend = (userId: string) => {
    console.log('Adding friend:', userId)
  }

  const handleViewProfile = (userId: string) => {
    console.log('Viewing profile:', userId)
    navigate(`/user/${userId}`)
  }

  const [lastMessage, setLastMessage] = useState<string>('')

  useEffect(() => {
    if (!socket) return

    // Listen for private messages
    socket.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data)
        if (data.type === 'private_message') {
          console.log('Received private message:', data)
          setLastMessage(data.content)
          toast.success(`New message: ${data.content}`)
        }
      } catch (error) {
        console.error('Error parsing message:', error)
      }
    }

    return () => {
      if (socket) {
        socket.onmessage = null
      }
    }
  }, [socket])

  const handleSendMessage = (userId: string) => {
    if (!isConnected) {
      console.warn('WebSocket is not connected')
      return
    }
    if (sendPrivateMessage) {
      sendPrivateMessage(userId, 'Hello! This is a test message from Anh Quoc')
      console.log('Sent test message to user:', userId)
    }
  }

  return (
    <>
      <div className="w-[360px] h-screen border-r border-base-200 flex flex-col bg-[#e3f2fd]">
        {/* Header */}
        <div>
          <div className="px-4 pt-4 flex items-center gap-4">
            <Avatar src={user?.avatar_url} size="md" />
            <div>
              <h3 className="font-medium text-success">Hello {user?.name}</h3>
              <h4>Have a good day</h4>
              {lastMessage && <p className="text-sm text-gray-600">Last message: {lastMessage}</p>}
            </div>
            <Button onClick={() => handleSendMessage('08ea9f7b-c6b0-425a-9bda-46fb6419a71a')}>Test Socket</Button>
          </div>
          <div className="px-4 pt-4 pb-4 border-b border-base-200 flex items-center justify-between">
            <div className="flex items-center gap-2">
              <h1 className="text-xl font-bold">Chats</h1>
            </div>
            <div className="flex items-center gap-2">
              <IconButton tooltip="Notifications" onClick={() => setShowNotificationModal(true)}>
                <BiBell className="w-5 h-5" />
              </IconButton>
              <IconButton tooltip="New Group" onClick={() => setShowCreateGroupModal(true)}>
                <BiPencil className="w-5 h-5" />
              </IconButton>
              <IconButton tooltip="Search Users" onClick={() => setShowSearchUsersModal(true)}>
                <BiSearch className="w-5 h-5" />
              </IconButton>
              <Dropdown
                trigger={
                  <IconButton>
                    <CgMore className="w-5 h-5" />
                  </IconButton>
                }
                items={[
                  { label: 'Profile', onClick: () => navigate('/profile') },
                  { label: 'Logout', onClick: handleLogout },
                ]}
              />
            </div>
          </div>
        </div>

        {/* Search */}
        <div className="p-4">
          <Input
            placeholder="Search conversations"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            startIcon={<BiSearch className="w-5 h-5" />}
          />
        </div>

        {/* Conversations List */}
        <div className="flex-1 overflow-y-auto">
          {conversations?.map((conversation: IConversation) => (
            <ConversationItem key={conversation.id} conversation={conversation} active={false} onClick={() => {}} />
          ))}
        </div>
      </div>

      <CreateGroupModal
        isOpen={showCreateGroupModal}
        onClose={() => setShowCreateGroupModal(false)}
        onConfirm={handleCreateGroup}
        availableUsers={mockAvailableUsers}
        currentUserId="1"
      />

      <SearchUsersModal
        isOpen={showSearchUsersModal}
        onClose={() => setShowSearchUsersModal(false)}
        onAddFriend={handleAddFriend}
        onViewProfile={handleViewProfile}
        availableUsers={mockAvailableSearchUsers}
        currentUserId="1"
      />

      <NotificationModal isOpen={showNotificationModal} onClose={() => setShowNotificationModal(false)} />
    </>
  )
}

export default LeftSidebar
