//hooks
import { useState } from 'react'

//ui
import Modal from '@ui/Modal'
import Input from '@ui/Input'
import Avatar from '@ui/Avatar'

//interfaces
import { IUser } from '@interfaces/user'

//assets
import group_default_img from '@assets/images/group_default.jpg'

interface CreateGroupModalProps {
  isOpen: boolean
  onClose: () => void
  onConfirm: (data: { name: string; description: string; avatar?: File; memberIds: string[] }) => void
  availableUsers: IUser[]
  currentUserId: string
}

const CreateGroupModal = ({ isOpen, onClose, onConfirm, availableUsers, currentUserId }: CreateGroupModalProps) => {
  const [name, setName] = useState('')
  const [description, setDescription] = useState('')
  const [selectedUsers, setSelectedUsers] = useState<string[]>([])
  const [avatar, setAvatar] = useState<File | null>(null)
  const [searchTerm, setSearchTerm] = useState('')
  const [previewUrl, setPreviewUrl] = useState('')

  const filteredUsers = availableUsers.filter(
    (user) =>
      user.id !== currentUserId &&
      (user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        user.email.toLowerCase().includes(searchTerm.toLowerCase())),
  )

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0]
    if (file) {
      setAvatar(file)
      const reader = new FileReader()
      reader.onloadend = () => {
        setPreviewUrl(reader.result as string)
      }
      reader.readAsDataURL(file)
    }
  }

  const handleToggleUser = (userId: string) => {
    setSelectedUsers((prev) => (prev.includes(userId) ? prev.filter((id) => id !== userId) : [...prev, userId]))
  }

  const handleSubmit = () => {
    if (name.trim() && description.trim() && selectedUsers.length > 0) {
      onConfirm({
        name: name.trim(),
        description: description.trim(),
        avatar: avatar || undefined,
        memberIds: selectedUsers,
      })
      resetForm()
    }
  }

  const resetForm = () => {
    setName('')
    setDescription('')
    setSelectedUsers([])
    setAvatar(null)
    setPreviewUrl('')
    setSearchTerm('')
  }

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Create New Group">
      <div className="py-4">
        <div className="flex items-center gap-4 mb-6">
          <div className="relative">
            <Avatar
              src={previewUrl || group_default_img}
              size="lg"
              className="cursor-pointer hover:opacity-80 transition-opacity"
            />
            <div
              className="absolute bottom-0 right-0 bg-primary text-white rounded-full p-1 cursor-pointer hover:bg-primary-focus"
              onClick={() => document.getElementById('groupAvatar')?.click()}
            >
              <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4" viewBox="0 0 20 20" fill="currentColor">
                <path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
              </svg>
            </div>
            <input type="file" id="groupAvatar" className="hidden" accept="image/*" onChange={handleFileChange} />
          </div>
          <div className="flex-1">
            <Input
              label="Group Name"
              placeholder="Enter group name"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          </div>
        </div>

        <div className="space-y-4">
          <Input
            label="Description"
            placeholder="Enter group description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className="mb-6"
          />

          <Input
            label="Search Members"
            placeholder="Search by name or email"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>

        <div className="max-h-[200px] overflow-y-auto mb-6">
          {filteredUsers.map((user) => (
            <div
              key={user.id}
              className={`flex items-center gap-3 p-2 rounded-lg cursor-pointer hover:bg-base-200 transition-colors ${
                selectedUsers.includes(user.id) ? 'bg-base-200' : ''
              }`}
              onClick={() => handleToggleUser(user.id)}
            >
              <Avatar src={user.avatarURL} size="sm" />
              <div className="flex-1">
                <div className="font-medium">{user.name}</div>
                <div className="text-sm text-base-content/60">{user.email}</div>
              </div>
              <input
                type="checkbox"
                checked={selectedUsers.includes(user.id)}
                onChange={() => handleToggleUser(user.id)}
                className="checkbox"
              />
            </div>
          ))}
        </div>

        <div className="flex justify-between items-center">
          <span className="text-sm text-base-content/60">{selectedUsers.length} members selected</span>
          <div className="flex gap-2">
            <button
              className="btn btn-ghost"
              onClick={() => {
                onClose()
                resetForm()
              }}
            >
              Cancel
            </button>
            <button
              className="btn btn-primary"
              onClick={handleSubmit}
              disabled={!name.trim() || !description.trim() || selectedUsers.length === 0}
            >
              Create Group
            </button>
          </div>
        </div>
      </div>
    </Modal>
  )
}

export default CreateGroupModal
