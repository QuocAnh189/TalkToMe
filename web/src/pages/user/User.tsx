import { useState } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
// import { useGetUserByIdQuery } from '@redux/services/user'
import Avatar from '@ui/Avatar'
import Button from '@ui/Button'
import { IoPersonAdd, IoPersonRemove } from 'react-icons/io5'
import { BiHome, BiMessageDetail } from 'react-icons/bi'
import ProtectedLayout from '@components/layout/protected'

// Mock user data
const mockUser = {
  id: '123',
  name: 'John Doe',
  email: 'john.doe@example.com',
  avatar_url: 'https://api.dicebear.com/7.x/avataaars/svg?seed=John',
  role: 'member',
  createdAt: '2023-01-15T08:00:00.000Z',
  isOnline: true,
}

const User = () => {
  const { userId } = useParams()
  const navigate = useNavigate()
  // const { data: user, isLoading } = useGetUserByIdQuery(userId || '')
  const [isFriend, setIsFriend] = useState(false)

  // Simulate loading state for demo
  const isLoading = false
  // Use mock data instead of query
  const user = mockUser

  const handleAddFriend = () => {
    console.log('Adding friend:', userId)
    setIsFriend(true)
  }

  const handleRemoveFriend = () => {
    console.log('Removing friend:', userId)
    setIsFriend(false)
  }

  const handleStartChat = () => {
    console.log('Starting chat with:', userId)
    // Navigate to chat or open chat modal
  }

  if (isLoading) {
    return (
      <ProtectedLayout>
        <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-50 p-8 flex items-center justify-center">
          <div className="loading loading-spinner loading-lg"></div>
        </div>
      </ProtectedLayout>
    )
  }

  if (!user) {
    return (
      <ProtectedLayout>
        <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-50 p-8 flex items-center justify-center">
          <div className="text-xl text-gray-600">User not found</div>
        </div>
      </ProtectedLayout>
    )
  }

  return (
    <ProtectedLayout>
      <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-50 p-8">
        <div className="max-w-4xl mx-auto">
          <div className="bg-white rounded-2xl shadow-xl overflow-hidden">
            {/* Header Banner */}
            <div className="h-48 bg-gradient-to-r from-blue-500 to-indigo-500"></div>

            {/* Profile Info */}
            <div className="relative px-8 pb-8">
              <button
                onClick={() => {
                  navigate('/')
                }}
                className="absolute top-4 left-4 bg-white/20 hover:bg-white/30 text-white rounded-lg px-4 py-2 flex items-center gap-2 backdrop-blur-sm transition-all"
              >
                <BiHome className="w-5 h-5" />
                <span>Back Home</span>
              </button>
              <div className="flex flex-col items-center -mt-20">
                <Avatar src={user.avatar_url} size="xl" className="ring-4 ring-white" />
                <h1 className="text-3xl font-bold mt-4">{user.name}</h1>
                <p className="text-gray-600 mt-1">{user.email}</p>

                {/* Action Buttons */}
                <div className="flex gap-4 mt-6">
                  {!isFriend ? (
                    <Button variant="primary" className="gap-2" onClick={handleAddFriend}>
                      <IoPersonAdd className="w-5 h-5" />
                      Add Friend
                    </Button>
                  ) : (
                    <Button variant="error" className="gap-2" onClick={handleRemoveFriend}>
                      <IoPersonRemove className="w-5 h-5" />
                      Remove Friend
                    </Button>
                  )}
                  <Button variant="secondary" className="gap-2" onClick={handleStartChat}>
                    <BiMessageDetail className="w-5 h-5" />
                    Message
                  </Button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </ProtectedLayout>
  )
}

export default User
