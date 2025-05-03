import { useState } from 'react'
import Modal from '@ui/Modal'
import Avatar from '@ui/Avatar'

interface NotificationModalProps {
  isOpen: boolean
  onClose: () => void
}

const NotificationModal = ({ isOpen, onClose }: NotificationModalProps) => {
  const [activeTab, setActiveTab] = useState<'friend' | 'group'>('friend')

  // Mock data - replace with actual data from your backend
  const friendRequests = [
    {
      id: '1',
      user: {
        id: 'user1',
        name: 'John Doe',
        avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=John',
      },
      timestamp: new Date().toISOString(),
    },
    {
      id: '2',
      user: {
        id: 'user2',
        name: 'Emma Wilson',
        avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Emma',
      },
      timestamp: new Date().toISOString(),
    },
    {
      id: '3',
      user: {
        id: 'user3',
        name: 'Michael Chen',
        avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Michael',
      },
      timestamp: new Date().toISOString(),
    },
    {
      id: '4',
      user: {
        id: 'user4',
        name: 'Sarah Johnson',
        avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah',
      },
      timestamp: new Date().toISOString(),
    },
    {
      id: '5',
      user: {
        id: 'user5',
        name: 'David Kim',
        avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=David',
      },
      timestamp: new Date().toISOString(),
    },
    {
      id: '6',
      user: {
        id: 'user6',
        name: 'Lisa Garcia',
        avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Lisa',
      },
      timestamp: new Date().toISOString(),
    },
    {
      id: '7',
      user: {
        id: 'user7',
        name: 'Alex Taylor',
        avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex',
      },
      timestamp: new Date().toISOString(),
    },
    {
      id: '8',
      user: {
        id: 'user8',
        name: 'Sophia Martinez',
        avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sophia',
      },
      timestamp: new Date().toISOString(),
    },
    {
      id: '9',
      user: {
        id: 'user9',
        name: 'James Wilson',
        avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=James',
      },
      timestamp: new Date().toISOString(),
    },
    {
      id: '10',
      user: {
        id: 'user10',
        name: 'Olivia Brown',
        avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Olivia',
      },
      timestamp: new Date().toISOString(),
    },
  ]

  const groupNotifications = [
    {
      id: '1',
      type: 'add',
      group: {
        id: 'group1',
        name: 'Project Team',
      },
      user: {
        id: 'user2',
        name: 'Jane Smith',
        avatarURL: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Jane',
      },
      timestamp: new Date().toISOString(),
    },
  ]

  const handleAcceptFriend = (userId: string) => {
    console.log('Accepting friend request:', userId)
    // Implement friend request acceptance logic
  }

  const handleRejectFriend = (userId: string) => {
    console.log('Rejecting friend request:', userId)
    // Implement friend request rejection logic
  }

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Notifications">
      <div className="py-4">
        <div className="flex gap-4 mb-6">
          <button
            className={`flex-1 py-2 px-4 rounded-lg transition-colors ${
              activeTab === 'friend' ? 'bg-primary text-white' : 'hover:bg-base-200'
            }`}
            onClick={() => setActiveTab('friend')}
          >
            Friend Requests
          </button>
          <button
            className={`flex-1 py-2 px-4 rounded-lg transition-colors ${
              activeTab === 'group' ? 'bg-primary text-white' : 'hover:bg-base-200'
            }`}
            onClick={() => setActiveTab('group')}
          >
            Group Notifications
          </button>
        </div>

        <div className="max-h-[400px] overflow-y-auto">
          {activeTab === 'friend' ? (
            <div className="space-y-4">
              {friendRequests.map((request) => (
                <div key={request.id} className="flex items-center gap-4 p-4 bg-base-100 rounded-lg">
                  <Avatar src={request.user.avatarURL} size="md" />
                  <div className="flex-1">
                    <h3 className="font-medium">{request.user.name}</h3>
                    <p className="text-sm text-base-content/60">Sent you a friend request</p>
                  </div>
                  <div className="flex gap-2">
                    <button className="btn btn-primary btn-sm" onClick={() => handleAcceptFriend(request.user.id)}>
                      Accept
                    </button>
                    <button className="btn btn-ghost btn-sm" onClick={() => handleRejectFriend(request.user.id)}>
                      Reject
                    </button>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <div className="space-y-4 min-h-[300px]">
              {groupNotifications.map((notification) => (
                <div key={notification.id} className="flex items-center gap-4 p-4 bg-base-100 rounded-lg">
                  <Avatar src={notification.user.avatarURL} size="md" />
                  <div className="flex-1">
                    <h3 className="font-medium">{notification.user.name}</h3>
                    <p className="text-sm text-base-content/60">
                      {notification.type === 'add'
                        ? `Added you to ${notification.group.name}`
                        : `Removed you from ${notification.group.name}`}
                    </p>
                  </div>
                </div>
              ))}
            </div>
          )}

          {((activeTab === 'friend' && friendRequests.length === 0) ||
            (activeTab === 'group' && groupNotifications.length === 0)) && (
            <div className="text-center text-base-content/60 py-8">No notifications</div>
          )}
        </div>
      </div>
    </Modal>
  )
}

export default NotificationModal
