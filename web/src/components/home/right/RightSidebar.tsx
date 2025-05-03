import { useState } from 'react'
import Avatar from '@ui/Avatar'
import Button from '@ui/Button'

//assets
import logo_img from '@assets/images/logo.png'

//icons
import { BiBell, BiChevronDown, BiPalette, BiSearch } from 'react-icons/bi'

//data
import { mockGroupMembers } from '../../../data/group_members'
import GroupMember from './GroupMember'

const RightSidebar = () => {
  const [showMembers, setShowMembers] = useState<boolean>(true)

  const groupMembers: any = mockGroupMembers

  return (
    <div className="w-[360px] h-screen border-l border-base-200 p-4 bg-[#e3f2fd]">
      <div className="flex flex-col items-center gap-4">
        <Avatar src={logo_img} size="lg" online={true} />
        <h2 className="text-xl font-bold">Anh Quoc</h2>
        {true && <span className="text-sm text-base-content/60">Active now</span>}
      </div>

      <div className="mt-8">
        <h3 className="font-semibold mb-2">Customize Chat</h3>
        <div className="flex flex-col gap-2">
          <Button variant="ghost" className="justify-start">
            <BiBell className="w-5 h-5 mr-2" />
            Notifications
          </Button>
          <Button variant="ghost" className="justify-start">
            <BiPalette className="w-5 h-5 mr-2" />
            Theme
          </Button>
          <Button variant="ghost" className="justify-start">
            <BiSearch className="w-5 h-5 mr-2" />
            Search in Conversation
          </Button>
        </div>
      </div>

      {true && (
        <div className="mt-8">
          <Button variant="ghost" className="w-full justify-between" onClick={() => setShowMembers(!showMembers)}>
            <span>Members</span>
            <BiChevronDown className={`w-5 h-5 transform ${showMembers ? 'rotate-180' : ''}`} />
          </Button>
          {showMembers && (
            <div className="mt-2 space-y-2">
              {groupMembers?.map((member: any) => (
                <GroupMember key={member.id} member={member} />
              ))}
            </div>
          )}
        </div>
      )}
    </div>
  )
}

export default RightSidebar
