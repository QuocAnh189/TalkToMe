//hooks
import { useState } from 'react'

//ui
import Avatar from '@ui/Avatar'
import IconButton from '@ui/IconButton'
import Dropdown from '@ui/Dropdown'

//assets
import logo_img from '@assets/images/logo.png'

//icons
import { BiPhone, BiVideo } from 'react-icons/bi'
import { CgMore } from 'react-icons/cg'

const ConversationHeader = () => {
  const [_, setShowProfile] = useState<boolean>(false)

  const handleSetNickname = () => {}
  const handleDeleteChat = () => {}

  return (
    <div className="h-[72px] border-b border-base-200 px-4 flex items-center justify-between mt-6">
      <div className="flex items-center gap-3">
        <Avatar src={logo_img} size="md" online={true} />
        <div>
          <h2 className="font-semibold">Anh Quoc</h2>
          {true && <span className="text-sm text-base-content/60">Active now</span>}
        </div>
      </div>

      <div className="flex items-center gap-2">
        <IconButton tooltip="Voice Call">
          <BiPhone className="w-5 h-5" />
        </IconButton>
        <IconButton tooltip="Video Call">
          <BiVideo className="w-5 h-5" />
        </IconButton>
        <Dropdown
          trigger={
            <IconButton>
              <CgMore className="w-5 h-5" />
            </IconButton>
          }
          items={[
            { label: 'View Profile', onClick: () => setShowProfile(true) },
            { label: 'Set Nickname', onClick: handleSetNickname },
            { label: 'Delete Conversation', onClick: handleDeleteChat },
          ]}
        />
      </div>
    </div>
  )
}

export default ConversationHeader
