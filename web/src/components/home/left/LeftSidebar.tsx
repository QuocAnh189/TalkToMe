//hooks
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
// import { useListConversationsQuery } from '@redux/services/conversation'
// import { useListGroupsQuery } from '@redux/services/group'
// import { IConversation } from '@interfaces/conversation'
// import { IGroup } from '@interfaces/group'

//ui
import Avatar from '@ui/Avatar'
import Input from '@ui/Input'
import IconButton from '@ui/IconButton'
import ConversationItem from '@ui/ConversationItem'
import Dropdown from '@ui/Dropdown'

//assets
import logo_img from '@assets/images/logo.png'

//icons
import { BiPencil, BiSearch, BiBell } from 'react-icons/bi'
import { CgMore } from 'react-icons/cg'

//data
import { mockConversations, mockGroups } from '../../../data/conversation'
import CreateGroupModal from '../modal/CreateGroupModal'
import SearchUsersModal from '../modal/SearchUsersModal'
import NotificationModal from '../modal/NotificationModal'

const LeftSidebar = () => {
  const navigate = useNavigate()
  const [searchTerm, setSearchTerm] = useState<string>('')
  // const { data: conversations } = useListConversationsQuery({})
  // const { data: groups } = useListGroupsQuery({})

  const conversations = mockConversations
  const groups = mockGroups

  const handleLogout = () => {
    navigate('/login')
  }

  const [showCreateGroupModal, setShowCreateGroupModal] = useState<boolean>(false)
  const [showSearchUsersModal, setShowSearchUsersModal] = useState<boolean>(false)
  const [showNotificationModal, setShowNotificationModal] = useState<boolean>(false)

  const handleCreateGroup = (data: { name: string; description: string; avatar?: File; memberIds: string[] }) => {
    console.log('Creating group:', data)
    setShowCreateGroupModal(false)
    // Add your group creation logic here
  }

  const mockAvailableUsers = [
    {
      id: '3',
      name: 'Bob Wilson',
      email: 'bob@example.com',
      avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Bob',
    },
    {
      id: '4',
      name: 'Emma Davis',
      email: 'emma@example.com',
      avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Emma',
    },
    {
      id: '5',
      name: 'Emma Davis',
      email: 'emma@example.com',
      avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Json',
    },
    {
      id: '6',
      name: 'Emma Davis',
      email: 'emma@example.com',
      avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Xml',
    },
    {
      id: '7',
      name: 'Emma Davis',
      email: 'emma@example.com',
      avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Yasou',
    },
    {
      id: '8',
      name: 'Emma Davis',
      email: 'emma@example.com',
      avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Rakan',
    },
    {
      id: '9',
      name: 'Emma Davis',
      email: 'emma@example.com',
      avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Darius',
    },
    {
      id: '10',
      name: 'Emma Davis',
      email: 'emma@example.com',
      avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Garen',
    },
  ]

  const handleAddFriend = (userId: string) => {
    console.log('Adding friend:', userId)
    // Add your friend request logic here
  }

  const handleViewProfile = (userId: string) => {
    console.log('Viewing profile:', userId)
    // Add your profile view logic here
    navigate(`/user/${userId}`)
  }

  return (
    <>
      <div className="w-[360px] h-screen border-r border-base-200 flex flex-col bg-[#e3f2fd]">
        {/* Header */}
        <div className="px-4 pt-8 pb-4 border-b border-base-200 flex items-center justify-between">
          <div className="flex items-center gap-2">
            <Avatar src={logo_img} size="md" />
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
          {conversations?.map((conversation: any) => (
            <ConversationItem
              key={conversation.id}
              conversation={conversation}
              // active={activeChat?.id === conversation.id}
              // onClick={() => setActiveChat({ id: conversation.id, type: 'conversation' })}
              active={false}
              onClick={() => {}}
            />
          ))}
          {groups?.map((group: any) => (
            <ConversationItem
              key={group.id}
              conversation={group}
              // active={activeChat?.id === group.id}
              // onClick={() => setActiveChat({ id: group.id, type: 'group' })}
              active={false}
              onClick={() => {}}
            />
          ))}
        </div>
      </div>

      <CreateGroupModal
        isOpen={showCreateGroupModal}
        onClose={() => setShowCreateGroupModal(false)}
        onConfirm={handleCreateGroup}
        availableUsers={mockAvailableUsers}
        currentUserId="1" // Replace with actual current user ID
      />

      <SearchUsersModal
        isOpen={showSearchUsersModal}
        onClose={() => setShowSearchUsersModal(false)}
        onAddFriend={handleAddFriend}
        onViewProfile={handleViewProfile}
        availableUsers={mockAvailableUsers.map((user) => ({ ...user, isFriend: false }))}
        currentUserId="1" // Replace with actual current user ID
      />

      <NotificationModal isOpen={showNotificationModal} onClose={() => setShowNotificationModal(false)} />
    </>
  )
}

export default LeftSidebar
