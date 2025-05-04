//hooks
import { useState } from 'react'

//components
import DeleteGroupConfirmationModal from '../modal/DeleteGroupConfirmationModal'
import ExitGroupConfirmationModal from '../modal/ExitGroupConfirmationModal'
import ChangeNicknameModal from '../modal/ChangeNicknameGroupModal'
import ChangeGroupNameModal from '../modal/ChangeGroupNameModal'
import ChangeGroupAvatarModal from '../modal/ChangeGroupAvatarModal'
import AddGroupMemberModal from '../modal/AddGroupMemberModal'
import ChangePrivateNicknameModal from '../modal/ChangePrivateNicknameModal'
import DeleteChatConfirmationModal from '../modal/DeleteChatConfirmationModal'

//ui
import Avatar from '@ui/Avatar'
import IconButton from '@ui/IconButton'
import Dropdown from '@ui/Dropdown'

//assets
import logo_img from '@assets/images/logo.png'

//icons
import { BiPhone, BiVideo } from 'react-icons/bi'
import { CgMore, CgProfile } from 'react-icons/cg'
import { MdOutlineEdit, MdDelete } from 'react-icons/md'
import { IoPersonAdd } from 'react-icons/io5'
import { IoMdExit } from 'react-icons/io'

//data
import { mockGroupMembers } from '../../../data/group_members'

const ConversationHeader = () => {
  const [showDeleteModal, setShowDeleteModal] = useState<boolean>(false)
  const [showExitModal, setShowExitModal] = useState<boolean>(false)
  const [showNicknameModal, setShowNicknameModal] = useState<boolean>(false)
  const [showGroupNameModal, setShowGroupNameModal] = useState<boolean>(false)
  const [showAvatarModal, setShowAvatarModal] = useState<boolean>(false)
  const [showAddMemberModal, setShowAddMemberModal] = useState<boolean>(false)
  const [showPrivateNicknameModal, setShowPrivateNicknameModal] = useState<boolean>(false)

  const handleViewProfile = () => {}
  const handleSetNickname = () => {
    setShowPrivateNicknameModal(true)
  }

  const mockCurrentUser = {
    id: '1',
    name: 'Current User',
    avatar_url: 'https://api.dicebear.com/7.x/avataaars/svg?seed=CurrentUser',
    nickname: 'Me',
  }

  const mockFriendUser = {
    id: '2',
    name: 'Anh Quoc',
    avatar_url: logo_img,
    nickname: 'Friend',
  }

  const handleConfirmPrivateNickname = (userId: string, nickname: string) => {
    console.log('User ID:', userId, 'New nickname:', nickname)
    setShowPrivateNicknameModal(false)
  }

  const handleDeleteChat = () => {
    setShowDeleteModal(true)
  }

  const handleChangeNicknameInGroup = () => {
    setShowNicknameModal(true)
  }

  const handleChangeAvatarGroup = () => {
    setShowAvatarModal(true)
  }

  const handleConfirmAvatar = (file: File) => {
    console.log('New avatar file:', file)
    setShowAvatarModal(false)
    // Add your avatar change logic here
  }

  const handleChangeNameGroup = () => {
    setShowGroupNameModal(true)
  }

  const handleConfirmGroupName = (name: string) => {
    console.log('New group name:', name)
    setShowGroupNameModal(false)
    // Add your group name change logic here
  }

  const handleDeleteGroup = () => {
    setShowDeleteModal(true)
  }

  const mockAvailableUsers = [
    {
      id: '3',
      name: 'Bob Wilson',
      email: 'bob@example.com',
      avatar_url: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Bob',
    },
    {
      id: '4',
      name: 'Emma Davis',
      email: 'emma@example.com',
      avatar_url: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Emma',
    },
  ]

  const handleAddMember = () => {
    setShowAddMemberModal(true)
  }

  const handleConfirmAddMembers = (selectedUsers: string[]) => {
    console.log('Selected users:', selectedUsers)
    setShowAddMemberModal(false)
  }

  const handleConfirmDelete = () => {
    setShowDeleteModal(false)
  }

  const handleExitGroup = () => {
    setShowExitModal(true)
  }

  const handleConfirmExit = () => {
    setShowExitModal(false)
  }

  const handleConfirmNickname = (memberId: string, nickname: string) => {
    console.log('Member ID:', memberId, 'New nickname:', nickname)
    setShowNicknameModal(false)
    // Add your nickname change logic here
  }

  const handleConfirmDeleteChat = () => {
    console.log('Deleting chat...')
    setShowDeleteModal(false)
  }

  return (
    <>
      <div className="h-[72px] border-b border-base-200 px-4 flex items-center justify-between mt-6">
        <div className="flex items-center gap-3">
          <Avatar src={logo_img} size="md" online={true} />
          <div>
            <h2 className="font-semibold">Anh Quoc</h2>
            {true && <span className="text-sm text-base-content/60">Active now</span>}
          </div>
        </div>

        <div className="flex items-center gap-2">
          <IconButton tooltip="Voice Call">
            <BiPhone className="w-5 h-5" />
          </IconButton>
          <IconButton tooltip="Video Call">
            <BiVideo className="w-5 h-5" />
          </IconButton>
          <Dropdown
            trigger={
              <IconButton>
                <CgMore className="w-5 h-5" />
              </IconButton>
            }
            items={
              false
                ? [
                    { icon: <CgProfile />, label: 'View Profile', onClick: handleViewProfile },
                    { icon: <MdOutlineEdit />, label: 'Nickname', onClick: handleSetNickname },
                    { icon: <MdDelete />, label: 'Delete', onClick: handleDeleteChat },
                  ]
                : [
                    { icon: <MdOutlineEdit />, label: 'Nickname', onClick: handleChangeNicknameInGroup },
                    { icon: <MdOutlineEdit />, label: 'Name', onClick: handleChangeNameGroup },
                    { icon: <MdOutlineEdit />, label: 'Change Avatar', onClick: handleChangeAvatarGroup },
                    { icon: <IoPersonAdd />, label: 'Add Member', onClick: handleAddMember },
                    { icon: <IoMdExit />, label: 'Exit', onClick: handleExitGroup },
                    { icon: <MdDelete />, label: 'Delete', onClick: handleDeleteGroup },
                  ]
            }
          />
        </div>
      </div>
      <DeleteGroupConfirmationModal
        isOpen={showDeleteModal}
        onClose={() => setShowDeleteModal(false)}
        onConfirm={handleConfirmDelete}
      />
      <ExitGroupConfirmationModal
        isOpen={showExitModal}
        onClose={() => setShowExitModal(false)}
        onConfirm={handleConfirmExit}
      />
      <ChangeNicknameModal
        isOpen={showNicknameModal}
        onClose={() => setShowNicknameModal(false)}
        onConfirm={handleConfirmNickname}
        groupMembers={mockGroupMembers}
        currentUserId="1"
      />
      <ChangeGroupNameModal
        isOpen={showGroupNameModal}
        onClose={() => setShowGroupNameModal(false)}
        onConfirm={handleConfirmGroupName}
        currentName="Current Group Name"
      />
      <ChangeGroupAvatarModal
        isOpen={showAvatarModal}
        onClose={() => setShowAvatarModal(false)}
        onConfirm={handleConfirmAvatar}
        currentAvatar={logo_img}
      />
      <AddGroupMemberModal
        isOpen={showAddMemberModal}
        onClose={() => setShowAddMemberModal(false)}
        onConfirm={handleConfirmAddMembers}
        availableUsers={mockAvailableUsers}
        currentMembers={mockGroupMembers.map((member) => member.id)}
      />
      <ChangePrivateNicknameModal
        isOpen={showPrivateNicknameModal}
        onClose={() => setShowPrivateNicknameModal(false)}
        onConfirm={handleConfirmPrivateNickname}
        currentUser={mockCurrentUser}
        friendUser={mockFriendUser}
      />
      <DeleteChatConfirmationModal
        isOpen={showDeleteModal}
        onClose={() => setShowDeleteModal(false)}
        onConfirm={handleConfirmDeleteChat}
        friendName={mockFriendUser.name}
      />
    </>
  )
}

export default ConversationHeader
