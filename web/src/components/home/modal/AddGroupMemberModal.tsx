import { useState } from 'react'
import Modal from '@ui/Modal'
import Input from '@ui/Input'
import Avatar from '@ui/Avatar'

interface User {
  id: string
  name: string
  avatarURL: string
  email: string
}

interface AddGroupMemberModalProps {
  isOpen: boolean
  onClose: () => void
  onConfirm: (selectedUsers: string[]) => void
  availableUsers: User[]
  currentMembers: string[]
}

const AddGroupMemberModal = ({
  isOpen,
  onClose,
  onConfirm,
  availableUsers,
  currentMembers,
}: AddGroupMemberModalProps) => {
  const [searchTerm, setSearchTerm] = useState('')
  const [selectedUsers, setSelectedUsers] = useState<string[]>([])

  const filteredUsers = availableUsers.filter(
    (user) =>
      !currentMembers.includes(user.id) &&
      (user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        user.email.toLowerCase().includes(searchTerm.toLowerCase())),
  )

  const handleToggleUser = (userId: string) => {
    setSelectedUsers((prev) => (prev.includes(userId) ? prev.filter((id) => id !== userId) : [...prev, userId]))
  }

  const handleSubmit = () => {
    if (selectedUsers.length > 0) {
      onConfirm(selectedUsers)
      setSelectedUsers([])
      setSearchTerm('')
    }
  }

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Add Group Members">
      <div className="py-4">
        <Input
          placeholder="Search users..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="mb-4"
        />

        <div className="max-h-[300px] overflow-y-auto mt-4">
          {filteredUsers.map((user) => (
            <div
              key={user.id}
              className={`flex items-center gap-3 p-2 rounded-lg cursor-pointer hover:bg-base-200 transition-colors ${
                selectedUsers.includes(user.id) ? 'bg-base-200' : ''
              }`}
              onClick={() => handleToggleUser(user.id)}
            >
              <div className="flex-shrink-0">
                <Avatar src={user.avatarURL} size="sm" />
              </div>
              <div className="flex-grow">
                <div className="font-medium">{user.name}</div>
                <div className="text-sm text-base-content/60">{user.email}</div>
              </div>
              <div className="flex-shrink-0">
                <input
                  type="checkbox"
                  checked={selectedUsers.includes(user.id)}
                  onChange={() => handleToggleUser(user.id)}
                  className="checkbox"
                />
              </div>
            </div>
          ))}

          {filteredUsers.length === 0 && <div className="text-center text-base-content/60 py-4">No users found</div>}
        </div>

        <div className="flex justify-between items-center mt-4">
          <span className="text-sm text-base-content/60">{selectedUsers.length} users selected</span>
          <div className="flex gap-2">
            <button
              className="btn btn-ghost"
              onClick={() => {
                onClose()
                setSelectedUsers([])
                setSearchTerm('')
              }}
            >
              Cancel
            </button>
            <button className="btn btn-primary" onClick={handleSubmit} disabled={selectedUsers.length === 0}>
              Add Members
            </button>
          </div>
        </div>
      </div>
    </Modal>
  )
}

export default AddGroupMemberModal
