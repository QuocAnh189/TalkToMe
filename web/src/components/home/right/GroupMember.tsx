//hooks
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

//ui
import Avatar from '@ui/Avatar'
import Modal from '@ui/Modal'
import Button from '@ui/Button'

interface GroupMemberProps {
  member: {
    id: string
    name: string
    email: string
    avatar_url: string
    isOnline: boolean
    isAdmin: boolean
  }
}

const GroupMember = ({ member }: GroupMemberProps) => {
  const [showProfile, setShowProfile] = useState(false)
  const navigate = useNavigate()

  const handleViewFullProfile = () => {
    setShowProfile(false)
    navigate(`/user/${member.id}`)
  }

  return (
    <>
      <div
        className="flex items-center gap-2 p-2 hover:bg-base-200 rounded-lg cursor-pointer transition-colors"
        onClick={() => setShowProfile(true)}
      >
        <Avatar src={member.avatar_url} size="sm" online={member.isOnline} />
        <div className="flex flex-col">
          <span className="font-medium">{member.name}</span>
          <span className="text-xs text-base-content/60">{member.isAdmin ? 'Admin' : 'Member'}</span>
        </div>
      </div>

      <Modal isOpen={showProfile} onClose={() => setShowProfile(false)} title="">
        <div className="flex flex-col items-center gap-4">
          <Avatar src={member.avatar_url} size="lg" online={member.isOnline} />
          <div className="text-center">
            <h3 className="text-xl font-bold">{member.name}</h3>
            {member.email && <p className="text-base-content/60">{member.email}</p>}
            <p className="text-sm mt-2">{member.isOnline ? 'Active now' : 'Offline'}</p>
            <p className="text-sm text-primary">{member.isAdmin ? 'Admin' : 'Members'}</p>
            <Button variant="primary" className="mt-4" onClick={handleViewFullProfile}>
              View Profile
            </Button>
          </div>
        </div>
      </Modal>
    </>
  )
}

export default GroupMember
