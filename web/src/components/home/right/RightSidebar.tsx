//hooks
import { useState } from 'react'

//ui
import Avatar from '@ui/Avatar'
import Button from '@ui/Button'

//components
import GroupMember from './GroupMember'
import ThemeModal from '../modal/ThemeModal'

//assets
import logo_img from '@assets/images/logo.png'

//icons
import { BiBell, BiChevronDown, BiPalette } from 'react-icons/bi'

//data
import { mockGroupMembers } from '../../../data/group_members'

const RightSidebar = () => {
  const [showMembers, setShowMembers] = useState<boolean>(true)
  const [showThemeModal, setShowThemeModal] = useState<boolean>(false)
  const [currentThemeId, setCurrentThemeId] = useState<string>('default')

  const handleThemeSelect = (themeId: string) => {
    setCurrentThemeId(themeId)
    setShowThemeModal(false)
    // Add your theme change logic here
  }

  const groupMembers: any = mockGroupMembers

  return (
    <>
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
            <Button variant="ghost" className="justify-start" onClick={() => setShowThemeModal(true)}>
              <BiPalette className="w-5 h-5 mr-2" />
              Theme
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
      <ThemeModal
        isOpen={showThemeModal}
        onClose={() => setShowThemeModal(false)}
        onSelectTheme={handleThemeSelect}
        currentThemeId={currentThemeId}
      />
    </>
  )
}

export default RightSidebar
