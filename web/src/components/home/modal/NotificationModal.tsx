//hooks
import { useState } from 'react'

//ui
import Modal from '@ui/Modal'
import Avatar from '@ui/Avatar'

//data
import { mockNotifications } from '../../../data/notification'

//interfaces
import { INotification } from '@interfaces/notification'

//libs
import { formatDistanceToNow } from 'date-fns'

interface NotificationModalProps {
  isOpen: boolean
  onClose: () => void
}

const NotificationModal = ({ isOpen, onClose }: NotificationModalProps) => {
  const [activeTab, setActiveTab] = useState<'friend' | 'group'>('friend')

  const friendRequests = mockNotifications.filter(
    (notification: INotification) => notification.type === 'friend_request',
  )

  const groupNotifications = mockNotifications.filter((notification: INotification) =>
    notification.type.startsWith('group_'),
  )

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
            Friend Requests {friendRequests.length > 0 && `(${friendRequests.length})`}
          </button>
          <button
            className={`flex-1 py-2 px-4 rounded-lg transition-colors ${
              activeTab === 'group' ? 'bg-primary text-white' : 'hover:bg-base-200'
            }`}
            onClick={() => setActiveTab('group')}
          >
            Group Notifications {groupNotifications.length > 0 && `(${groupNotifications.length})`}
          </button>
        </div>

        <div className="max-h-[400px] overflow-y-auto">
          {activeTab === 'friend' ? (
            <div className="space-y-4">
              {friendRequests.map((notification: INotification) => (
                <div key={notification.id} className="flex items-center gap-4 p-4 bg-base-100 rounded-lg">
                  <Avatar src={notification.data.user.avatar_url} size="md" />
                  <div className="flex-1">
                    <h3 className="font-medium">{notification.data.user.name}</h3>
                    <p className="text-sm text-base-content/60">{notification.content}</p>
                    <p className="text-xs text-base-content/40 mt-1">
                      {formatDistanceToNow(new Date(notification.createdAt), { addSuffix: true })}
                    </p>
                  </div>
                  <div className="flex gap-2">
                    <button
                      className="btn btn-primary btn-sm"
                      onClick={() => handleAcceptFriend(notification.data.user.id)}
                    >
                      Accept
                    </button>
                    <button
                      className="btn btn-ghost btn-sm"
                      onClick={() => handleRejectFriend(notification.data.user.id)}
                    >
                      Reject
                    </button>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <div className="space-y-4">
              {groupNotifications.map((notification: INotification) => (
                <div key={notification.id} className="flex items-center gap-4 p-4 bg-base-100 rounded-lg">
                  <Avatar src={notification.data.user.avatar_url} size="md" />
                  <div className="flex-1">
                    <h3 className="font-medium">{notification.data.user.name}</h3>
                    <p className="text-sm text-base-content/60">{notification.content}</p>
                    <p className="text-xs text-base-content/40 mt-1">
                      {formatDistanceToNow(new Date(notification.createdAt), { addSuffix: true })}
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
