import { useState } from 'react'
import Modal from '@ui/Modal'
import Input from '@ui/Input'
import Avatar from '@ui/Avatar'

interface User {
  id: string
  name: string
  avatarURL: string
  nickname?: string
}

interface ChangePrivateNicknameModalProps {
  isOpen: boolean
  onClose: () => void
  onConfirm: (userId: string, nickname: string) => void
  currentUser: User
  friendUser: User
}

const ChangePrivateNicknameModal = ({ 
  isOpen, 
  onClose, 
  onConfirm, 
  currentUser,
  friendUser 
}: ChangePrivateNicknameModalProps) => {
  const [selectedUser, setSelectedUser] = useState<'me' | 'friend'>('me')
  const [nickname, setNickname] = useState('')

  const handleUserSelect = (userType: 'me' | 'friend') => {
    setSelectedUser(userType)
    setNickname(userType === 'me' ? currentUser.nickname || '' : friendUser.nickname || '')
  }

  const handleSubmit = () => {
    if (nickname.trim()) {
      const userId = selectedUser === 'me' ? currentUser.id : friendUser.id
      onConfirm(userId, nickname.trim())
      setNickname('')
    }
  }

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Change Nickname">
      <div className="py-4">
        <div className="flex gap-4 mb-6">
          <div
            className={`flex-1 p-3 rounded-lg cursor-pointer transition-colors ${
              selectedUser === 'me' ? 'bg-base-200' : 'hover:bg-base-100'
            }`}
            onClick={() => handleUserSelect('me')}
          >
            <div className="flex items-center gap-3">
              <Avatar src={currentUser.avatarURL} size="sm" />
              <div>
                <div className="font-medium">Me ({currentUser.name})</div>
                {currentUser.nickname && (
                  <div className="text-sm text-base-content/60">
                    Nickname: {currentUser.nickname}
                  </div>
                )}
              </div>
            </div>
          </div>

          <div
            className={`flex-1 p-3 rounded-lg cursor-pointer transition-colors ${
              selectedUser === 'friend' ? 'bg-base-200' : 'hover:bg-base-100'
            }`}
            onClick={() => handleUserSelect('friend')}
          >
            <div className="flex items-center gap-3">
              <Avatar src={friendUser.avatarURL} size="sm" />
              <div>
                <div className="font-medium">{friendUser.name}</div>
                {friendUser.nickname && (
                  <div className="text-sm text-base-content/60">
                    Nickname: {friendUser.nickname}
                  </div>
                )}
              </div>
            </div>
          </div>
        </div>

        <Input
          label={`Nickname for ${selectedUser === 'me' ? 'yourself' : friendUser.name}`}
          placeholder="Enter nickname"
          value={nickname}
          onChange={(e) => setNickname(e.target.value)}
        />

        <div className="flex justify-end gap-2 mt-6">
          <button 
            className="btn btn-ghost" 
            onClick={() => {
              onClose()
              setNickname('')
              setSelectedUser('me')
            }}
          >
            Cancel
          </button>
          <button
            className="btn btn-primary"
            onClick={handleSubmit}
            disabled={!nickname.trim()}
          >
            Save
          </button>
        </div>
      </div>
    </Modal>
  )
}

export default ChangePrivateNicknameModal
