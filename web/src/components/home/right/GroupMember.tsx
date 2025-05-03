//hooks
import { useState } from 'react'

//ui
import Avatar from '@ui/Avatar'
import Modal from '@ui/Modal'

interface GroupMemberProps {
  member: {
    id: string
    name: string
    avatarURL: string
    isOnline: boolean
    role?: string
    email?: string
  }
}

const GroupMember = ({ member }: GroupMemberProps) => {
  const [showProfile, setShowProfile] = useState(false)

  return (
    <>
      <div
        className="flex items-center gap-2 p-2 hover:bg-base-200 rounded-lg cursor-pointer transition-colors"
        onClick={() => setShowProfile(true)}
      >
        <Avatar src={member.avatarURL} size="sm" online={member.isOnline} />
        <div className="flex flex-col">
          <span className="font-medium">{member.name}</span>
          {member.role && <span className="text-xs text-base-content/60">{member.role}</span>}
        </div>
      </div>

      <Modal isOpen={showProfile} onClose={() => setShowProfile(false)} title="User Profile">
        <div className="flex flex-col items-center gap-4">
          <Avatar src={member.avatarURL} size="lg" online={member.isOnline} />
          <div className="text-center">
            <h3 className="text-xl font-bold">{member.name}</h3>
            {member.email && <p className="text-base-content/60">{member.email}</p>}
            <p className="text-sm mt-2">{member.isOnline ? 'Active now' : 'Offline'}</p>
            {member.role && <p className="text-sm text-primary">{member.role}</p>}
          </div>
        </div>
      </Modal>
    </>
  )
}

export default GroupMember
