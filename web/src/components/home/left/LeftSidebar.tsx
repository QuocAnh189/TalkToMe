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
import { BiPencil, BiSearch } from 'react-icons/bi'
import { CgMore } from 'react-icons/cg'

//data
import { mockConversations, mockGroups } from '../../../data/conversation'

const LeftSidebar = () => {
  const navigate = useNavigate()
  const [searchTerm, setSearchTerm] = useState('')
  // const { data: conversations } = useListConversationsQuery({})
  // const { data: groups } = useListGroupsQuery({})

  const conversations = mockConversations
  const groups = mockGroups

  const handleLogout = () => {
    navigate('/login')
  }

  return (
    <div className="w-[360px] h-screen border-r border-base-200 flex flex-col bg-[#e3f2fd]">
      {/* Header */}
      <div className="px-4 pt-8 pb-4 border-b border-base-200 flex items-center justify-between">
        <div className="flex items-center gap-2">
          <Avatar src={logo_img} size="md" />
          <h1 className="text-xl font-bold">Chats</h1>
        </div>
        <div className="flex items-center gap-2">
          <IconButton tooltip="New Group">
            <BiPencil className="w-5 h-5" />
          </IconButton>
          <IconButton tooltip="Search Users">
            <BiSearch className="w-5 h-5" />
          </IconButton>
          <Dropdown
            trigger={
              <IconButton>
                <CgMore className="w-5 h-5" />
              </IconButton>
            }
            items={[
              { label: 'Settings', onClick: () => navigate('/settings') },
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
  )
}

export default LeftSidebar
