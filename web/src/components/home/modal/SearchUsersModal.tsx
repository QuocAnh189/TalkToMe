//hooks
import { useState } from 'react'

//ui
import Modal from '@ui/Modal'
import Input from '@ui/Input'
import Avatar from '@ui/Avatar'
import IconButton from '@ui/IconButton'

//icons
import { BiSearch } from 'react-icons/bi'
import { IoPersonAdd } from 'react-icons/io5'
import { CgProfile } from 'react-icons/cg'

//interfaces
import { ISearchUser } from '@interfaces/user'

interface SearchUsersModalProps {
  isOpen: boolean
  onClose: () => void
  onAddFriend: (userId: string) => void
  onViewProfile: (userId: string) => void
  availableUsers: ISearchUser[]
  currentUserId: string
}

const SearchUsersModal = ({
  isOpen,
  onClose,
  onAddFriend,
  onViewProfile,
  availableUsers,
  currentUserId,
}: SearchUsersModalProps) => {
  const [searchTerm, setSearchTerm] = useState('')
  const [viewMode, setViewMode] = useState<'list' | 'grid'>('grid')

  const filteredUsers = availableUsers.filter(
    (user) =>
      user.id !== currentUserId &&
      (user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        user.email.toLowerCase().includes(searchTerm.toLowerCase())),
  )

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Search Users">
      <div className="py-4">
        <div className="flex items-center gap-4 mb-6">
          <div className="flex-1">
            <Input
              placeholder="Search by name or email"
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              startIcon={<BiSearch className="w-5 h-5" />}
            />
          </div>
          <div className="flex items-center gap-2">
            <button
              className={`btn btn-sm ${viewMode === 'list' ? 'btn-primary' : 'btn-ghost'}`}
              onClick={() => setViewMode('list')}
            >
              List
            </button>
            <button
              className={`btn btn-sm ${viewMode === 'grid' ? 'btn-primary' : 'btn-ghost'}`}
              onClick={() => setViewMode('grid')}
            >
              Grid
            </button>
          </div>
        </div>

        {viewMode === 'grid' ? (
          <div className="grid grid-cols-2 md:grid-cols-2 gap-4 max-h-[600px] overflow-y-auto p-2">
            {filteredUsers.map((user) => (
              <div
                key={user.id}
                className="bg-base-100 rounded-lg p-4 hover:shadow-md transition-shadow cursor-pointer"
              >
                <div className="flex flex-col items-center text-center">
                  <Avatar src={user.avatarURL} size="lg" className="mb-3 hover:opacity-80 transition-opacity" />
                  <h3 className="font-semibold text-lg mb-1">{user.name}</h3>
                  <p className="w-40 text-sm text-base-content/60 mb-3 truncate">{user.email}</p>
                  <div>
                    <IconButton tooltip="View Profile" onClick={() => onViewProfile(user.id)}>
                      <CgProfile className="w-5 h-5" />
                    </IconButton>
                    {!user.isFriend && (
                      <IconButton tooltip="Add Friend" onClick={() => onAddFriend(user.id)}>
                        <IoPersonAdd className="w-5 h-5" />
                      </IconButton>
                    )}
                  </div>
                </div>
              </div>
            ))}
          </div>
        ) : (
          <div className="max-h-[600px] overflow-y-auto">
            {filteredUsers.map((user) => (
              <div
                key={user.id}
                className="flex items-center gap-3 py-3 px-6 rounded-lg hover:bg-base-200 transition-colors"
              >
                <Avatar src={user.avatarURL} size="md" className="cursor-pointer" />
                <div className="flex-1 cursor-pointer">
                  <div className="font-medium">{user.name}</div>
                  <div className="text-sm text-base-content/60 truncate">{user.email}</div>
                </div>
                <IconButton tooltip="View Profile" onClick={() => onViewProfile(user.id)}>
                  <CgProfile className="w-5 h-5" />
                </IconButton>
                {!user.isFriend && (
                  <IconButton tooltip="Add Friend" onClick={() => onAddFriend(user.id)}>
                    <IoPersonAdd className="w-5 h-5" />
                  </IconButton>
                )}
              </div>
            ))}
          </div>
        )}

        {filteredUsers.length === 0 && <div className="text-center text-base-content/60 py-4">No users found</div>}
      </div>
    </Modal>
  )
}

export default SearchUsersModal
