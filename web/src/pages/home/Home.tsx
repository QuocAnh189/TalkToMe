//hooks
// import { useState } from 'react'

//components
import ProtectedLayout from '@components/layout/protected'
import LeftSidebar from '@components/home/left/LeftSidebar'
import Middle from '@components/home/middle/Middle'
import RightSidebar from '@components/home/right/RightSidebar'

// interface ActiveChat {
//   id: string
//   type: 'conversation' | 'group'
// }

const Home = () => {
  // const [activeChat] = useState<ActiveChat | null>(null)
  // const [showRightSidebar] = useState(true)

  return (
    <ProtectedLayout>
      <div className="flex h-screen">
        <LeftSidebar />

        {/* {activeChat ? (
          <Middle />
        ) : (
          <div className="flex-1 flex items-center justify-center text-base-content/60">
            Select a conversation to start chatting
          </div>
        )} */}
        <Middle />

        {/* {activeChat && showRightSidebar && <RightSidebar />} */}
        <RightSidebar />
      </div>
    </ProtectedLayout>
  )
}

export default Home
