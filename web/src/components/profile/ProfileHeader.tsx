//hooks
import { useNavigate } from 'react-router-dom'

import { BiHome } from 'react-icons/bi'

const ProfileHeader = () => {
  const navigate = useNavigate()
  return (
    <div className="relative">
      <div className="h-64 bg-gradient-to-r from-blue-500 via-indigo-500 to-purple-500 relative">
        <button
          onClick={() => {
            navigate('/')
          }}
          className="absolute top-4 left-4 bg-white/20 hover:bg-white/30 text-white rounded-lg px-4 py-2 flex items-center gap-2 backdrop-blur-sm transition-all"
        >
          <BiHome className="w-5 h-5" />
          <span>Back Home</span>
        </button>
      </div>
    </div>
  )
}

export default ProfileHeader
